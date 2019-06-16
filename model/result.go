package model

import "time"

// Result struct
type Result struct {
	Position   uint
	RunnerID   string
	RunnerName string
	Laps       uint
	RunnerTime time.Time
}
