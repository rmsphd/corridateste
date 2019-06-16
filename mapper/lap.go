package mapper

import (
	"fmt"
	"interview-test/model"
	"strconv"
	"strings"
	"time"
)

// newLap Cria um novo objeto Lap a partir de um slice da linha
func newLap(line []string) (*model.Lap, error) {
	hour, err := time.Parse("15:04:05.000", line[0])
	if err != nil {
		fmt.Printf("[newLap] Hour Error: %v\n", err)
		return nil, err
	}

	lapID, err := strconv.ParseUint(line[3], 10, 32)
	if err != nil {
		fmt.Printf("[newLap] Lap Error: %v\n", err)
		return nil, err
	}

	t, err := time.Parse("4:05.000", line[4])
	if err != nil {
		fmt.Printf("[newLap] Time Error: %v\n", err)
		return nil, err
	}

	speed, err := strconv.ParseFloat(strings.Replace(line[5], ",", ".", 1), 32)
	if err != nil {
		fmt.Printf("[newLap] Speed Error: %v\n", err)
		return nil, err
	}

	lap := model.Lap{
		Hour:       hour,
		RunnerID:   line[1],
		RunnerName: line[2],
		LapID:      uint(lapID),
		LapTime:    t,
		Speed:      float32(speed),
	}

	return &lap, nil
}

// NewLaps map de cada linha
func NewLaps(lines [][]string) ([]model.Lap, error) {
	laps := make([]model.Lap, 0)

	for _, line := range lines {
		lap, err := newLap(line)
		if err != nil {
			fmt.Printf("[NewLaps] Error: %v\n", err)
			return nil, err
		}

		laps = append(laps, *lap)
	}

	return laps, nil
}
