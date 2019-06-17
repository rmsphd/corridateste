package builder

import (
	"fmt"
	"interview-test/model"
	"reflect"
	"testing"
	"time"
)

func ExampleAverageSpeed() {
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

	fmt.Printf("%v", AverageSpeed(laps))

	// Output: [{1 001 test name 2 0000-01-01 00:01:23.123 +0000 UTC 20.0615}]
}

func TestAverageSpeed(t *testing.T) {
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
			Speed:      40.123,
		},
		{
			Position:   2,
			RunnerID:   "001",
			RunnerName: "test name",
			Laps:       2,
			RunnerTime: ltr2,
			Speed:      40.123,
		},
		{
			Position:   3,
			RunnerID:   "003",
			RunnerName: "test name 3",
			Laps:       1,
			RunnerTime: lt,
			Speed:      40.123,
		},
	}

	if got := AverageSpeed(laps); !reflect.DeepEqual(got, want) {
		t.Errorf("AverageSpeed() = %v, want %v", got, want)
	}

}
