package main

import (
	"fmt"
	"interview-test/builder"
	"interview-test/file"
	"interview-test/mapper"
	"interview-test/printer"
	"os"
)

func main() {
	args := os.Args
	if len(args) <= 1 {
		fmt.Printf(`Error: Need to use with argument file name
Example: %v corrida.log
`, args[0])
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

	results := builder.CalculateChampions(laps)

	printer.PrintChampions(results)

}
