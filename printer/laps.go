package printer

import (
	"fmt"
	"interview-test/model"
	"io"
	"text/tabwriter"
)

func PrintLaps(laps []model.Lap, output io.Writer) {
	w := tabwriter.NewWriter(output, 1, 2, 2, ' ', tabwriter.AlignRight|tabwriter.Debug)
	fmt.Fprintln(w, "Código Piloto\tNome Piloto\tNúmero Volta\tTempo da Volta\tVelocidade média")

	for _, l := range laps {
		fmt.Fprintf(w, "%s\t%s\t%v\t%v\t%g\n", l.RunnerID, l.RunnerName, l.LapID, l.LapTime.Format("4:05.000"), l.Speed)
	}
	w.Flush()
}
