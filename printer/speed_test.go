package printer

import (
	"bytes"
	"interview-test/model"
	"os"
	"testing"
	"time"
)

func ExamplePrintSpeeds() {
	th, _ := time.Parse("4:05.000", "1:23.123")
	r := []model.Result{
		{
			Position:   1,
			RunnerID:   "1",
			RunnerName: "test",
			Laps:       1,
			RunnerTime: th,
			Speed:      40.123,
		},
	}

	PrintSpeeds(r, os.Stdout)

	// Output: Posição Chegada|  Código Piloto|  Nome Piloto|  Qtde Voltas Completadas|  Tempo Total de Prova|Velocidade média
	//                 1|              1|         test|                        1|              1:23.123|40.123
}

func TestPrintSpeeds(t *testing.T) {
	th, _ := time.Parse("4:05.000", "1:23.123")
	r := []model.Result{
		{
			Position:   1,
			RunnerID:   "1",
			RunnerName: "test",
			Laps:       1,
			RunnerTime: th,
			Speed:      40.123,
		},
	}
	var w bytes.Buffer
	PrintSpeeds(r, &w)
	got := w.String()
	want := ``

	if got == want {
		t.Errorf("PrintSpeeds() = \n%s\n, want \n%s", got, want)
	}

}
