package printer

import (
	"fmt"
	"interview-test/model"
	"os"
	"text/tabwriter"
)

// PrintChampions imprime o resultado na tela
func PrintChampions(results []model.Result, output *os.File) {
	w := tabwriter.NewWriter(output, 1, 2, 2, ' ', tabwriter.AlignRight|tabwriter.Debug)
	fmt.Fprintln(w, "Posição Chegada\tCódigo Piloto\tNome Piloto\tQtde Voltas Completadas\tTempo Total de Prova")

	for _, r := range results {
		fmt.Fprintf(w, "%v\t%s\t%s\t%v\t%v\n", r.Position, r.RunnerID, r.RunnerName, r.Laps, r.RunnerTime)
	}
	w.Flush()
}
