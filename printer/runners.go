package printer

import (
	"fmt"
	"interview-test/model"
	"io"
	"text/tabwriter"
)

// PrintRunners imprime o resultado
func PrintRunners(results []model.Result, output io.Writer) {
	w := tabwriter.NewWriter(output, 1, 2, 2, ' ', tabwriter.AlignRight|tabwriter.Debug)
	fmt.Fprintln(w, "Posição Chegada\tCódigo Piloto\tNome Piloto\tQtde Voltas Completadas\tTempo")

	for _, r := range results {
		fmt.Fprintf(w, "%v\t%s\t%s\t%v\t%s\n", r.Position, r.RunnerID, r.RunnerName, r.Laps, r.RunnerTime.Format("4:05.000"))
	}
	w.Flush()
}
