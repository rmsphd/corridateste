package builder

import (
	"interview-test/model"
	"sort"
	"time"
)

// CalculateChampions calcula os campeões da corrida
func CalculateChampions(laps []model.Lap) []model.Result {
	results := make([]model.Result, 0)

	for _, lap := range laps {
		i, r, exist := getResult(results, lap.RunnerID)

		if exist {
			results[i] = updateResult(lap, r)
		} else {
			results = append(results, newResult(lap))
		}
	}

	rankResults(results)

	return results
}

// newResult Cria um novo resultado para um corredor
func newResult(lap model.Lap) model.Result {
	return model.Result{
		RunnerID:   lap.RunnerID,
		RunnerName: lap.RunnerName,
		Laps:       lap.LapID,
		RunnerTime: lap.LapTime,
	}
}

// updateResult Atualiza os dados de um corredor
func updateResult(lap model.Lap, r model.Result) model.Result {
	moreTime := (time.Duration(lap.LapTime.Minute()) * time.Minute) +
		(time.Duration(lap.LapTime.Second()) * time.Second) +
		(time.Duration(lap.LapTime.Nanosecond()) * time.Nanosecond)

	if lap.LapID > r.Laps {
		r.Laps = lap.LapID
	}
	r.RunnerTime = r.RunnerTime.Add(moreTime)
	return r
}

// rankResults ordena os resultado e define a posição de cada um
func rankResults(results []model.Result) {
	sort.SliceStable(results, func(a int, b int) bool {
		if results[a].Laps == results[b].Laps {
			return results[a].RunnerTime.Before(results[b].RunnerTime)
		}
		return results[a].Laps > results[b].Laps
	})

	for i, r := range results {
		r.Position = uint(i + 1)
		results[i] = r
	}
}

func getResult(results []model.Result, id string) (int, model.Result, bool) {
	for i, r := range results {
		if r.RunnerID == id {
			return i, r, true
		}
	}
	return -1, model.Result{}, false
}
