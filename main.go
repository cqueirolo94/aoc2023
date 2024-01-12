package main

import (
	"aoc2023/days"
	"fmt"
)

const (
	DAY_1_INPUT = "./inputs/day1.txt"
	DAY_2_INPUT = "./inputs/day2.txt"
)

type DayChallenge interface {
	Challenge1()
	Challenge2()
}

func main() {
	fmt.Println("aoc 2023")
	// DAY 1
	d1 := days.NewDay1(DAY_1_INPUT)
	d1.Challenge1()
	d1.Challenge2()

	// DAY 2
	d2 := days.NewDay2(DAY_2_INPUT)
	d2.Challenge1()
	d2.Challenge2()

}
