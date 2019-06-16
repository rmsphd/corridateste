package model

import "time"

// Lap struct
type Lap struct {
	Hour       time.Time
	RunnerID   string
	RunnerName string
	LapID      uint
	LapTime    time.Time
	Speed      float32
}
