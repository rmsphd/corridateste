package main

import (
	"fmt"
	"interview-test/builder"
	"interview-test/file"
	"interview-test/mapper"
	"interview-test/model"
	"interview-test/printer"
	"os"
)

func main() {
	args := os.Args
	if len(args) <= 1 {
		fmt.Printf("Error: Need to use with argument file name\nExample: %v corrida.log [saida.txt]\n", args[0])
		return
	}

	name := args[1]

	f, err := os.Open(name)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer f.Close()

	lines, err := file.Reader(f)
	if err != nil {
		return
	}

	laps, err := mapper.NewLaps(lines)
	if err != nil {
		return
	}

	w := os.Stdout
	if len(args) > 2 {
		w, err = os.OpenFile(args[2], os.O_RDWR|os.O_CREATE, 0666)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
		defer w.Close()
	}

	fmt.Fprintf(w, "\nResultado da Corrida\n")
	results := builder.CalculateChampions(laps)
	printer.PrintChampions(results, w)

	fmt.Fprintf(w, "\nMelhor volta de cada piloto\n")
	betterLaps := builder.BetterLapForEachRunner(laps)
	printer.PrintLaps(betterLaps, w)

	fmt.Fprintf(w, "\nMelhor volta da corrida\n")
	betterLap := builder.BetterLap(laps)
	printer.PrintLaps([]model.Lap{betterLap}, w)

	fmt.Fprintf(w, "\nVelocidade média de cada piloto durante toda corrida\n")
	speeds := builder.AverageSpeed(laps)
	printer.PrintSpeeds(speeds, w)

	fmt.Fprintf(w, "\nQuanto tempo cada piloto chegou após o vencedor\n")
	runners := builder.TimeRunners(laps)
	printer.PrintRunners(runners, w)

}
