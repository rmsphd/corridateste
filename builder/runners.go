package builder

import (
	"interview-test/model"
	"time"
)

// TimeRunners tempo de cada piloto chegou após o vencedor
func TimeRunners(laps []model.Lap) []model.Result {
	results := CalculateChampions(laps)
	var firstTime time.Time
	if len(results) > 0 {
		firstTime = results[0].RunnerTime
	}

	for i, r := range results {
		diff := r.RunnerTime.Sub(firstTime)
		r.RunnerTime = time.Time{}.Add(diff)

		results[i] = r
	}

	return results
}
