package main

import (
	"database/sql"
	"errors"
	"log"
)

type GearData struct {
	GearStats Stats
	Head      string
	Arms      string
	Body      string
	Legs      string
}

type StrikerStats struct {
	Name       string
	BaseStats  Stats
	GearCombos []GearData
}

type StrikerRow struct {
	Name      string
	Strength  int
	Speed     int
	Shooting  int
	Passing   int
	Technique int
	Head      string
	Arms      string
	Body      string
	Legs      string
}

func scanStrikerRow(rows *sql.Rows) (StrikerRow, error) {
	var strikerRow StrikerRow
	err := rows.Scan(
		&strikerRow.Name,
		&strikerRow.Strength,
		&strikerRow.Speed,
		&strikerRow.Shooting,
		&strikerRow.Passing,
		&strikerRow.Technique,
		&strikerRow.Head,
		&strikerRow.Arms,
		&strikerRow.Body,
		&strikerRow.Legs,
	)
	return strikerRow, err
}

func collectStrikerRows(rows *sql.Rows) ([]StrikerRow, error) {
	defer rows.Close()
	var results []StrikerRow
	for rows.Next() {
		strikerRow, err := scanStrikerRow(rows)
		if err != nil {
			log.Printf("error scanning data >> %s", err)
			return []StrikerRow{}, err
		}

		results = append(results, strikerRow)
	}

	if err := rows.Err(); err != nil {
		log.Printf("error while iterating through rows >> %s", err)
		return []StrikerRow{}, err
	}

	return results, nil
}

func addComboToDB(db *sql.DB, striker Striker) {
	query := `
		INSERT INTO striker (striker_name, strength, speed, shooting, passing, technique, head, arms, body, legs)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		RETURNING striker_name`

	var name string
	err := db.QueryRow(query,
		striker.Name,
		striker.GearStats.Strength,
		striker.GearStats.Speed,
		striker.GearStats.Shooting,
		striker.GearStats.Passing,
		striker.GearStats.Technique,
		striker.Head.Name,
		striker.Arms.Name,
		striker.Body.Name,
		striker.Legs.Name).Scan(&name)

	if err != nil {
		panic(err)
	}
}

func GetStrikerFromDB(db *sql.DB, strikerName string) (StrikerStats, error) {
	if !IsValidStriker(strikerName) {
		log.Printf("invalid striker name input: %s", strikerName)
		return StrikerStats{}, errors.New("InvalidStrikerName")
	}

	query := `
		SELECT strength, speed, shooting, passing, technique, head, arms, body, legs FROM striker
		WHERE striker_name=$1`

	rows, err := db.Query(query, strikerName)
	if err != nil {
		log.Printf("error getting data for strikerName %s >> %s", strikerName, err)
		return StrikerStats{}, err
	}
	defer rows.Close()

	strikerStats := StrikerStats{
		Name:      strikerName,
		BaseStats: ALL_STRIKERS[strikerName].BaseStats,
	}

	for rows.Next() {
		var data GearData
		err = rows.Scan(
			&data.GearStats.Strength,
			&data.GearStats.Speed,
			&data.GearStats.Shooting,
			&data.GearStats.Passing,
			&data.GearStats.Technique,
			&data.Head,
			&data.Arms,
			&data.Body,
			&data.Legs,
		)
		if err != nil {
			log.Printf("error scanning data for strikerName %s >> %s", strikerName, err)
			return StrikerStats{}, err
		}

		strikerStats.GearCombos = append(strikerStats.GearCombos, data)
	}

	if err := rows.Err(); err != nil {
		log.Printf("error while iterating through rows for strikerName %s >> %s", strikerName, err)
		return StrikerStats{}, err
	}

	return strikerStats, nil
}

func GetBalancedStrikers(db *sql.DB) ([]StrikerRow, error) {
	rows, err := db.Query(GetBalancedQuery())
	if err != nil {
		log.Printf("error getting data >> %s", err)
		return []StrikerRow{}, err
	}
	return collectStrikerRows(rows)
}

func GetStrikersForStatPriority(db *sql.DB, statsList []string) ([]StrikerRow, error) {
	for _, stat := range statsList {
		if !IsValidStat(stat) {
			log.Printf("invalid stat input: %s", stat)
			return []StrikerRow{}, errors.New("InvalidStrikerStat")
		}
	}

	rows, err := db.Query(GetPrioritizedStatsQuery(statsList))
	if err != nil {
		log.Printf("error getting data for stat list >> %s", err)
		return []StrikerRow{}, err
	}

	return collectStrikerRows(rows)
}
