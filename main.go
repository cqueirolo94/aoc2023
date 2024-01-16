package main

import (
	"aoc2023/days"
	"fmt"
)

const (
	TEST_INPUT  = "./inputs/test.txt"
	DAY_1_INPUT = "./inputs/day1.txt"
	DAY_2_INPUT = "./inputs/day2.txt"
	DAY_3_INPUT = "./inputs/day3.txt"
	DAY_4_INPUT = "./inputs/day4.txt"
	DAY_5_INPUT = "./inputs/day5.txt"
)

type DayChallenge interface {
	Challenge1()
	Challenge2()
}

/*
EXAMPLE

type DayN struct {
	lines []string
}

func NewDayN(input string) *DayN {
	return &DayN{
		lines: read_file(input),
	}
}

func (d *DayN) Challenge1() {
	fmt.Println("Day N - Challenge 1")
	total := 0

	fmt.Printf("number: %d\n", total)
}

func (d *DayN) Challenge2() {
	fmt.Println("Day N - Challenge 2")
	total := 0

	fmt.Printf("number: %d\n", total)
}
*/

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

	// DAY 3
	d3 := days.NewDay3(DAY_3_INPUT)
	d3.Challenge1()
	d3.Challenge2()

	// DAY 4
	d4 := days.NewDay4(DAY_4_INPUT)
	d4.Challenge1()
	d4.Challenge2()

	// DAY 5
	d5 := days.NewDay5(TEST_INPUT)
	d5.Challenge1()
	d5.Challenge2()
}
