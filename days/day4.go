package days

import (
	"fmt"
	"slices"
	"strings"
)

type Day4 struct {
	lines []string
	cards map[string]int
}

func NewDay4(input string) *Day4 {
	return &Day4{
		lines: read_file(input),
		cards: make(map[string]int),
	}
}

func (d *Day4) Challenge1() {
	fmt.Println("Day 4 - Challenge 1")
	total := 0
	points := make(chan int, len(d.lines))

	for _, l := range d.lines {
		go func(l string, p chan int) {
			points := 0.5
			lsplit := strings.Split(l, "|")
			winning_nums := strings.Fields(strings.Split(lsplit[0], ": ")[1])
			elf_nums := strings.Fields(lsplit[1])
			for i := 0; i < len(winning_nums); i++ {
				if slices.Contains(elf_nums, winning_nums[i]) {
					points *= 2
				}
			}

			p <- int(points)
		}(l, points)
	}

	for range d.lines {
		total += <-points
	}

	fmt.Printf("number: %d\n", total)
}

func (d *Day4) Challenge2() {
	fmt.Println("Day 4 - Challenge 2")
	total := 0

	for i := 0; i < len(d.lines); i++ {
		wins := 0
		lsplit := strings.Split(d.lines[i], "|")
		aux := strings.Split(lsplit[0], ": ")
		card := aux[0]
		d.cards[card] = d.cards[card] + 1
		winning_nums := strings.Fields(aux[1])
		elf_nums := strings.Fields(lsplit[1])

		for _, n := range winning_nums {
			if slices.Contains(elf_nums, n) {
				wins++
			}
		}

		for j := i + 1; j < i+1+wins; j++ {
			_card := strings.Split(d.lines[j], ": ")[0]

			d.cards[_card] = d.cards[_card] + d.cards[card]
		}
	}

	for _, n := range d.cards {
		total += n
	}

	fmt.Printf("number: %d\n", total)
}
