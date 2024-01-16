package days

import "fmt"

type Day5 struct {
	lines []string
}

func NewDay5(input string) *Day5 {
	return &Day5{
		lines: read_file(input),
	}
}

/*
src to dest map
dest | source | offset
*/
func (d *Day5) Challenge1() {
	fmt.Println("Day 5 - Challenge 1")
	total := 0

	fmt.Printf("number: %d\n", total)
}

func (d *Day5) Challenge2() {
	fmt.Println("Day 5 - Challenge 2")
	total := 0

	fmt.Printf("number: %d\n", total)
}
