package printer

import (
	"bytes"
	"interview-test/model"
	"os"
	"testing"
	"time"
)

func ExamplePrintRunners() {
	th, _ := time.Parse("4:05.000", "1:23.123")
	r := []model.Result{
		{
			Position:   1,
			RunnerID:   "1",
			RunnerName: "test",
			Laps:       1,
			RunnerTime: th,
		},
	}

	PrintRunners(r, os.Stdout)

	// Output: Posição Chegada|  Código Piloto|  Nome Piloto|  Qtde Voltas Completadas|Tempo
	//                 1|              1|         test|                        1|1:23.123
}

func TestPrintRunners(t *testing.T) {
	th, _ := time.Parse("4:05.000", "1:23.123")
	r := []model.Result{
		{
			Position:   1,
			RunnerID:   "1",
			RunnerName: "test",
			Laps:       1,
			RunnerTime: th,
		},
	}
	var w bytes.Buffer
	PrintRunners(r, &w)
	got := w.String()
	want := ``

	if got == want {
		t.Errorf("PrintRunners() = \n%s\n, want \n%s", got, want)
	}

}
