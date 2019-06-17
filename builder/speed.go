package builder

import (
	"interview-test/model"
	"time"
)

// AverageSpeed Velocidade média de cada piloto durante toda corrida
func AverageSpeed(laps []model.Lap) []model.Result {
	results := make([]model.Result, 0)

	for _, lap := range laps {
		i, r, exist := getResult(results, lap.RunnerID)

		if exist {
			results[i] = updateSpeedResult(lap, r)
		} else {
			results = append(results, newResult(lap))
		}

	}

	rankResults(results)
	calculateAverage(results)

	return results
}

// calculateAverage calcula a média de velocidade
func calculateAverage(results []model.Result) {
	for i, result := range results {
		result.Speed = result.Speed / float32(result.Laps)
		results[i] = result
	}
}

// updateSpeedResult Atualiza os dados de um corredor com velocidade
func updateSpeedResult(lap model.Lap, r model.Result) model.Result {
	moreTime := (time.Duration(lap.LapTime.Minute()) * time.Minute) +
		(time.Duration(lap.LapTime.Second()) * time.Second) +
		(time.Duration(lap.LapTime.Nanosecond()) * time.Nanosecond)

	if lap.LapID > r.Laps {
		r.Laps = lap.LapID
	}
	r.RunnerTime = r.RunnerTime.Add(moreTime)
	r.Speed = r.Speed + lap.Speed
	return r
}
