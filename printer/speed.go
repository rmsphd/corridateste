package printer

import (
	"fmt"
	"interview-test/model"
	"io"
	"text/tabwriter"
)

// PrintSpeeds imprime o resultado na tela
func PrintSpeeds(results []model.Result, output io.Writer) {
	w := tabwriter.NewWriter(output, 1, 2, 2, ' ', tabwriter.AlignRight|tabwriter.Debug)
	fmt.Fprintln(w, "Posição Chegada\tCódigo Piloto\tNome Piloto\tQtde Voltas Completadas\tTempo Total de Prova\tVelocidade média")

	for _, r := range results {
		fmt.Fprintf(w, "%v\t%s\t%s\t%v\t%s\t%g\n", r.Position, r.RunnerID, r.RunnerName, r.Laps, r.RunnerTime.Format("4:05.000"), r.Speed)
	}
	w.Flush()
}
