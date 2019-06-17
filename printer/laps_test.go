package printer

import (
	"bytes"
	"interview-test/model"
	"os"
	"testing"
	"time"
)

func ExamplePrintLaps() {
	th, _ := time.Parse("4:05.000", "1:23.123")
	lt, _ := time.Parse("4:05.000", "1:23.123")
	laps := []model.Lap{
		{
			Hour:       th,
			RunnerID:   "001",
			RunnerName: "test name",
			LapID:      2,
			LapTime:    lt,
			Speed:      40.123,
		},
	}

	PrintLaps(laps, os.Stdout)

	// Output: Código Piloto|  Nome Piloto|  Número Volta|  Tempo da Volta|Velocidade média
	//             001|    test name|             2|        1:23.123|40.123
}

func TestPrintLaps(t *testing.T) {
	th, _ := time.Parse("4:05.000", "1:23.123")
	lt, _ := time.Parse("4:05.000", "1:23.123")
	laps := []model.Lap{
		{
			Hour:       th,
			RunnerID:   "001",
			RunnerName: "test name",
			LapID:      2,
			LapTime:    lt,
			Speed:      40.123,
		},
	}
	var w bytes.Buffer
	PrintLaps(laps, &w)
	got := w.String()
	want := ``

	if got == want {
		t.Errorf("PrintLaps() = \n%s\n, want \n%s", got, want)
	}

}
