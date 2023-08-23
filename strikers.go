package main

import "database/sql"

type Stats struct {
	Strength  int
	Speed     int
	Shooting  int
	Passing   int
	Technique int
}

type Gear struct {
	Name        string
	BodyPart    string
	StatsEffect Stats
}

type Striker struct {
	Name      string
	BaseStats Stats
	Head      Gear
	Arms      Gear
	Body      Gear
	Legs      Gear
	GearStats Stats
}

func (stats *Stats) updateStats(statsToAdd *Stats) {
	stats.Strength += statsToAdd.Strength
	stats.Speed += statsToAdd.Speed
	stats.Shooting += statsToAdd.Shooting
	stats.Passing += statsToAdd.Passing
	stats.Technique += statsToAdd.Technique
}

func (stats *Stats) updateStatsFromGear(gear *Gear) {
	stats.updateStats(&gear.StatsEffect)
}

func (striker *Striker) addGear(gear Gear) {
	switch bodyPart := gear.BodyPart; bodyPart {
	case HEAD:
		striker.Head = gear
	case ARMS:
		striker.Arms = gear
	case BODY:
		striker.Body = gear
	case LEGS:
		striker.Legs = gear
	}

	striker.GearStats.updateStatsFromGear(&gear)
}

func CalcAllStrikerGearStatCombos(db *sql.DB) {
	for _, striker := range ALL_STRIKERS {
		for _, headGear := range HEAD_GEAR {
			for _, armGear := range ARM_GEAR {
				for _, bodyGear := range BODY_GEAR {
					for _, legGear := range LEG_GEAR {
						strikerCombo := striker
						strikerCombo.GearStats.updateStats(&strikerCombo.BaseStats)

						strikerCombo.addGear(headGear)
						strikerCombo.addGear(armGear)
						strikerCombo.addGear(bodyGear)
						strikerCombo.addGear(legGear)

						addComboToDB(db, strikerCombo)
					}
				}
			}
		}
	}
}

func sumBaseStats(stats Stats) int {
	return stats.Strength + stats.Speed + stats.Shooting + stats.Passing + stats.Technique
}

func CalcStatTotal() map[string]int {
	results := make(map[string]int)
	for _, striker := range ALL_STRIKERS {
		sum := sumBaseStats(striker.BaseStats)
		results[striker.Name] = sum
	}

	return results
}

func IsValidStat(stat string) bool {
	return Contains(stat, ALL_STATS)
}

func IsValidStriker(name string) bool {
	_, ok := ALL_STRIKERS[name]
	return ok
}
