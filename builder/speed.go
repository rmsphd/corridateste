package builder

import (
	"interview-test/model"
)

// AverageSpeed Velocidade média de cada piloto durante toda corrida
func AverageSpeed(laps []model.Lap) []model.Result {
	results := CalculateChampions(laps)

	sumSpeedResult(laps, results)
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
func sumSpeedResult(laps []model.Lap, results []model.Result) {
	for i, result := range results {
		for _, lap := range laps {
			if lap.RunnerID == result.RunnerID {
				result.Speed += lap.Speed
			}
		}
		results[i] = result
	}
}
