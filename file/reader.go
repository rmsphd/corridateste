package file

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
)

// Reader Le o arquivo separado por espaços
func Reader(file io.Reader) ([][]string, error) {
	re := regexp.MustCompile("[\\s–]+")

	scanner := bufio.NewScanner(file)

	scanner.Scan() // Skip the first line

	lines := make([][]string, 0)

	for scanner.Scan() {
		line := scanner.Text()

		lines = append(lines, re.Split(line, -1))
	}

	if scanner.Err() != nil {
		fmt.Printf("[Reader] Error: %v\n", scanner.Err())
		return nil, scanner.Err()
	}

	return lines, nil
}
