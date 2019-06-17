package builder

import (
	"fmt"
	"interview-test/model"
	"reflect"
	"testing"
	"time"
)

func ExampleBetterLapForEachRunner() {
	h, _ := time.Parse("15:04:05.000", "08:12:01.123")
	lt, _ := time.Parse("4:05.000", "1:23.123")

	laps := []model.Lap{
		{
			Hour:       h,
			RunnerID:   "001",
			RunnerName: "test name",
			LapID:      2,
			LapTime:    lt,
			Speed:      40.123,
		},
	}

	fmt.Printf("%v", BetterLapForEachRunner(laps))

	// Output: [{0000-01-01 08:12:01.123 +0000 UTC 001 test name 2 0000-01-01 00:01:23.123 +0000 UTC 40.123}]
}

func TestBetterLapForEachRunner(t *testing.T) {
	h, _ := time.Parse("15:04:05.000", "08:12:01.123")
	lt, _ := time.Parse("4:05.000", "1:23.123")
	ltr, _ := time.Parse("4:05.000", "2:46.246")
	ltr2, _ := time.Parse("4:05.000", "4:09.369")
	laps := []model.Lap{
		{
			Hour:       h,
			RunnerID:   "001",
			RunnerName: "test name",
			LapID:      2,
			LapTime:    ltr,
			Speed:      40.123,
		},
		{
			Hour:       h,
			RunnerID:   "001",
			RunnerName: "test name",
			LapID:      1,
			LapTime:    lt,
			Speed:      40.123,
		},
		{
			Hour:       h,
			RunnerID:   "002",
			RunnerName: "test name 2",
			LapID:      1,
			LapTime:    lt,
			Speed:      40.123,
		},
		{
			Hour:       h,
			RunnerID:   "002",
			RunnerName: "test name 2",
			LapID:      2,
			LapTime:    ltr2,
			Speed:      40.123,
		},
		{
			Hour:       h,
			RunnerID:   "003",
			RunnerName: "test name 3",
			LapID:      1,
			LapTime:    lt,
			Speed:      40.123,
		},
	}
	want := []model.Lap{
		{
			Hour:       h,
			RunnerID:   "001",
			RunnerName: "test name",
			LapID:      1,
			LapTime:    lt,
			Speed:      40.123,
		},
		{
			Hour:       h,
			RunnerID:   "002",
			RunnerName: "test name 2",
			LapID:      1,
			LapTime:    lt,
			Speed:      40.123,
		},
		{
			Hour:       h,
			RunnerID:   "003",
			RunnerName: "test name 3",
			LapID:      1,
			LapTime:    lt,
			Speed:      40.123,
		},
	}

	if got := BetterLapForEachRunner(laps); !reflect.DeepEqual(got, want) {
		t.Errorf("BetterLapForEachRunner() = %v, want %v", got, want)
	}

}

func ExampleBetterLap() {
	h, _ := time.Parse("15:04:05.000", "08:12:01.123")
	lt, _ := time.Parse("4:05.000", "1:23.123")

	laps := []model.Lap{
		{
			Hour:       h,
			RunnerID:   "001",
			RunnerName: "test name",
			LapID:      2,
			LapTime:    lt,
			Speed:      40.123,
		},
	}

	fmt.Printf("%v", BetterLap(laps))

	// Output: {0000-01-01 08:12:01.123 +0000 UTC 001 test name 2 0000-01-01 00:01:23.123 +0000 UTC 40.123}
}

func TestBetterLap(t *testing.T) {
	h, _ := time.Parse("15:04:05.000", "08:12:01.123")
	lt, _ := time.Parse("4:05.000", "1:23.123")
	ltr, _ := time.Parse("4:05.000", "2:46.246")
	ltr2, _ := time.Parse("4:05.000", "4:09.369")
	laps := []model.Lap{
		{
			Hour:       h,
			RunnerID:   "001",
			RunnerName: "test name",
			LapID:      2,
			LapTime:    ltr,
			Speed:      40.123,
		},
		{
			Hour:       h,
			RunnerID:   "001",
			RunnerName: "test name",
			LapID:      1,
			LapTime:    lt,
			Speed:      40.123,
		},
		{
			Hour:       h,
			RunnerID:   "002",
			RunnerName: "test name 2",
			LapID:      1,
			LapTime:    lt,
			Speed:      40.123,
		},
		{
			Hour:       h,
			RunnerID:   "002",
			RunnerName: "test name 2",
			LapID:      2,
			LapTime:    ltr2,
			Speed:      40.123,
		},
		{
			Hour:       h,
			RunnerID:   "003",
			RunnerName: "test name 3",
			LapID:      1,
			LapTime:    lt,
			Speed:      40.123,
		},
	}
	want := model.Lap{
		Hour:       h,
		RunnerID:   "001",
		RunnerName: "test name",
		LapID:      1,
		LapTime:    lt,
		Speed:      40.123,
	}

	if got := BetterLap(laps); !reflect.DeepEqual(got, want) {
		t.Errorf("BetterLap() = %v, want %v", got, want)
	}

}

func TestBetterLap_Empty(t *testing.T) {
	laps := []model.Lap{}
	want := model.Lap{}

	if got := BetterLap(laps); !reflect.DeepEqual(got, want) {
		t.Errorf("BetterLap() = %v, want %v", got, want)
	}

}
