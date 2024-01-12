package days

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Day2 struct {
	lines []string
	cubes map[string]int
	games map[int]map[string][]int
}

func NewDay2(input string) *Day2 {
	lines := read_file(input)
	games := make(map[int]map[string][]int)
	for _, l := range lines {
		l = l[5:]
		l = strings.ReplaceAll(l, ":", ",")
		l = strings.ReplaceAll(l, ";", ",")
		l = strings.ReplaceAll(l, ", ", ",")
		lSplitted := strings.Split(l, ",")

		// Parse game number
		g, _ := strconv.Atoi(lSplitted[0])
		games[g] = map[string][]int{
			"red":   make([]int, 0),
			"green": make([]int, 0),
			"blue":  make([]int, 0),
		}

		// Parse draws
		lSplitted = lSplitted[1:]
		for _, d := range lSplitted {
			dSplitted := strings.Split(d, " ")
			n, _ := strconv.Atoi(dSplitted[0])
			games[g][dSplitted[1]] = append(games[g][dSplitted[1]], n)
		}
	}
	return &Day2{
		lines: lines,
		cubes: map[string]int{
			"red":   12,
			"green": 13,
			"blue":  14,
		},
		games: games,
	}
}

func (d *Day2) Challenge1() {
	fmt.Println("Day 2 - Challenge 1")
	total := 0
	gc := make(chan int, len(d.games))

	for game, draws := range d.games {
		go func(game int, draws map[string][]int, gc chan int) {
			for color, draw := range draws {
				for _, num := range draw {
					if num > d.cubes[color] {
						gc <- 0
						return
					}
				}
			}
			gc <- game
		}(game, draws, gc)
	}

	for range d.games {
		total += <-gc
	}

	fmt.Printf("number: %d\n", total)
}

func (d *Day2) Challenge2() {
	fmt.Println("Day 2 - Challenge 2")
	total := 0
	pwc := make(chan int, len(d.games))

	for _, draws := range d.games {
		go func(draws map[string][]int, pwc chan int) {
			pw := 1
			for _, draw := range draws {
				pw *= slices.Max(draw)
			}
			pwc <- pw
		}(draws, pwc)
	}

	for range d.games {
		total += <-pwc
	}

	fmt.Printf("number: %d\n", total)
}
