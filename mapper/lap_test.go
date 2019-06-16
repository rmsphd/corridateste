package mapper

import (
	"interview-test/model"
	"reflect"
	"testing"
	"time"
)

func TestNewLaps(t *testing.T) {
	args := [][]string{
		[]string{"08:12:01.123", "001", "test name", "3", "1:23.123", "40,123"},
	}

	h, _ := time.Parse("15:04:05.000", "08:12:01.123")
	th, _ := time.Parse("4:05.000", "1:23.123")

	want := []model.Lap{
		model.Lap{
			Hour:       h,
			RunnerID:   "001",
			RunnerName: "test name",
			LapID:      uint(3),
			LapTime:    th,
			Speed:      float32(40.123),
		},
	}

	got, err := NewLaps(args)
	if err != nil {
		t.Errorf("NewLaps() error = %v", err)
		return
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("NewLaps() = %v, want %v", got, want)
	}
}

func TestNewLaps_InvalidHour(t *testing.T) {
	args := [][]string{
		[]string{"08", "001", "test name", "3", "1:23.123", "40,123"},
	}

	got, err := NewLaps(args)
	if err == nil {
		t.Errorf("NewLaps() error = %v", err)
		return
	}
	if got != nil {
		t.Errorf("NewLaps() = %v, want %v", got, nil)
	}
}

func TestNewLaps_InvalidTime(t *testing.T) {
	args := [][]string{
		[]string{"08:12:01.123", "001", "test name", "3", "1", "40,123"},
	}

	got, err := NewLaps(args)
	if err == nil {
		t.Errorf("NewLaps() error = %v", err)
		return
	}
	if got != nil {
		t.Errorf("NewLaps() = %v, want %v", got, nil)
	}
}

func TestNewLaps_InvalidSpeed(t *testing.T) {
	args := [][]string{
		[]string{"08:12:01.123", "001", "test name", "3", "1:23.123", "40A123"},
	}

	got, err := NewLaps(args)
	if err == nil {
		t.Errorf("NewLaps() error = %v", err)
		return
	}
	if got != nil {
		t.Errorf("NewLaps() = %v, want %v", got, nil)
	}
}

func TestNewLaps_InvalidLap(t *testing.T) {
	args := [][]string{
		[]string{"08:12:01.123", "001", "test name", "a", "1:23.123", "40,123"},
	}

	got, err := NewLaps(args)
	if err == nil {
		t.Errorf("NewLaps() error = %v", err)
		return
	}
	if got != nil {
		t.Errorf("NewLaps() = %v, want %v", got, nil)
	}
}
