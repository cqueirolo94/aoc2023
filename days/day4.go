package days

import "fmt"

type Day4 struct {
	lines []string
}

func NewDay4(input string) *Day4 {
	return &Day4{
		lines: read_file(input),
	}
}

func (d *Day4) Challenge1() {
	fmt.Println("Day 4 - Challenge 1")
	total := 0

	fmt.Printf("number: %d\n", total)
}

func (d *Day4) Challenge2() {
	fmt.Println("Day 4 - Challenge 2")
	total := 0

	fmt.Printf("number: %d\n", total)
}
