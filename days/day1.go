package days

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

type Day1 struct {
	numbers map[string]string
	lines   []string
}

func NewDay1(input string) *Day1 {
	return &Day1{
		numbers: map[string]string{
			"one":   "1",
			"two":   "2",
			"three": "3",
			"four":  "4",
			"five":  "5",
			"six":   "6",
			"seven": "7",
			"eight": "8",
			"nine":  "9",
		},
		lines: read_file(input),
	}
}

func (d *Day1) Challenge1() {
	fmt.Println("Day 1 - Challenge 1")
	numCh := make(chan int, len(d.lines))
	errCh := make(chan error)
	total := 0

	for _, l := range d.lines {
		go d.process_line(l, numCh, errCh)
	}

	for range d.lines {
		select {
		case num := <-numCh:
			total += num
		case err := <-errCh:
			panic(err)
		}
	}

	fmt.Printf("number: %d\n", total)
}

func (d *Day1) Challenge2() {
	fmt.Println("Day 1 - Challenge 2")
	numCh := make(chan int, len(d.lines))
	errCh := make(chan error)
	total := 0

	// TODO PRE-PROCESS LINE
	var wg sync.WaitGroup
	for i := 0; i < len(d.lines); i++ {
		wg.Add(1)
		go func(l *string, wg *sync.WaitGroup) {
			defer wg.Done()
			for k, v := range d.numbers {
				*l = strings.ReplaceAll(*l, k, fmt.Sprintf("%s%s%s", k[:len(k)-1], v, string(k[len(k)-1])))
			}
		}(&d.lines[i], &wg)
	}

	wg.Wait()

	for _, l := range d.lines {
		go d.process_line(l, numCh, errCh)
	}

	for range d.lines {
		select {
		case num := <-numCh:
			total += num
		case err := <-errCh:
			panic(err)
		}
	}

	fmt.Printf("number: %d\n", total)
}

func (d *Day1) process_line(l string, numCh chan int, errCh chan error) {
	numA, numB := "", ""

	for i := 0; i < len(l); i++ {
		if numA == "" {
			if _, err := strconv.Atoi(string(l[i])); err == nil {
				numA = string(l[i])
			}
		}
		if numB == "" {
			if _, err := strconv.Atoi(string(l[len(l)-1-i])); err == nil {
				numB = string(l[len(l)-1-i])
			}
		}
	}

	fNum, err := strconv.Atoi(numA + numB)
	if err != nil {
		errCh <- err
		return
	}

	numCh <- fNum
}
