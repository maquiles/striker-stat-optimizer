package main

import "fmt"

func GetAllQuery() string {
	return `SELECT * FROM striker`
}

func GetBalancedQuery() string {
	return `
	SELECT * FROM striker 
	WHERE strength > 11 AND strength < 14
		AND speed > 11 AND speed < 14
		AND shooting > 11 AND shooting < 14
		AND passing > 11 AND passing < 14
		AND technique > 11 AND technique < 14;`
}

func GetPrioritizedStatsQuery(stats []string) string {
	numStats := len(stats)
	if numStats == 0 {
		return GetAllQuery()
	}

	query := "SELECT * FROM striker WHERE "
	priority := make(map[string]struct{})
	first := true

	for i := 0; i < numStats; i++ {
		currPriorityStat := stats[i]
		priority[currPriorityStat] = struct{}{}

		for j := 0; j < len(ALL_STATS); j++ {
			currStat := ALL_STATS[j]
			if _, ok := priority[currStat]; ok {
				continue
			}

			if !first {
				query += "AND "
			}

			query += fmt.Sprintf("%s > %s ", currPriorityStat, currStat)
			first = false
		}
	}

	query += "ORDER BY"

	for k := 0; k < numStats; k++ {
		query += fmt.Sprintf(" %s DESC", stats[k])
		if !(numStats-1 == k) {
			query += ","
		}
	}

	return query
}
