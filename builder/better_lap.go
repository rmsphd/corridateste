package builder

import (
	"interview-test/model"
	"sort"
)

// BetterLapForEachRunner melhor volta de cada piloto
func BetterLapForEachRunner(laps []model.Lap) []model.Lap {
	results := make([]model.Lap, 0)

	for _, lap := range laps {
		i, r, exist := getLap(results, lap.RunnerID)

		if exist {
			results[i] = updateLapForEachRunner(lap, r)
		} else {
			results = append(results, lap)
		}
	}

	return results
}

// BetterLap melhor volta da corrida
func BetterLap(laps []model.Lap) model.Lap {
	if len(laps) > 0 {
		sort.SliceStable(laps, func(a int, b int) bool {
			return laps[a].LapTime.Before(laps[b].LapTime)
		})

		return laps[0]
	}
	return model.Lap{}
}

// getLap recupera um volta para um corredor
func getLap(results []model.Lap, id string) (int, model.Lap, bool) {
	for i, r := range results {
		if r.RunnerID == id {
			return i, r, true
		}
	}
	return -1, model.Lap{}, false
}

// updateLapForEachRunner atualiza os dados de um corredor
func updateLapForEachRunner(lap model.Lap, r model.Lap) model.Lap {
	if lap.LapTime.Before(r.LapTime) {
		r.LapID = lap.LapID
		r.LapTime = lap.LapTime
	}
	return r
}
