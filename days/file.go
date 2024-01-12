package days

import (
	"bufio"
	"fmt"
	"os"
)

func read_file(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		panic(fmt.Sprintf("could not open file: %s", err.Error()))
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
