package builder

import (
	"fmt"
	"interview-test/model"
	"reflect"
	"testing"
	"time"
)

func ExampleCalculateChampions() {
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

	fmt.Printf("%v", CalculateChampions(laps))

	// Output: [{1 001 test name 2 0000-01-01 00:01:23.123 +0000 UTC}]
}

func TestCalculateChampions(t *testing.T) {
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
			LapTime:    lt,
			Speed:      40.123,
		},
		{
			Hour:       h,
			RunnerID:   "001",
			RunnerName: "test name",
			LapID:      1,
			LapTime:    ltr,
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
	want := []model.Result{
		{
			Position:   1,
			RunnerID:   "002",
			RunnerName: "test name 2",
			Laps:       2,
			RunnerTime: ltr,
		},
		{
			Position:   2,
			RunnerID:   "001",
			RunnerName: "test name",
			Laps:       2,
			RunnerTime: ltr2,
		},
		{
			Position:   3,
			RunnerID:   "003",
			RunnerName: "test name 3",
			Laps:       1,
			RunnerTime: lt,
		},
	}

	if got := CalculateChampions(laps); !reflect.DeepEqual(got, want) {
		t.Errorf("CalculateChampions() = %v, want %v", got, want)
	}

}
