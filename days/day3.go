package days

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Day3 struct {
	lines   []string
	numbers []string
}

func NewDay3(input string) *Day3 {
	return &Day3{
		lines: read_file(input),
		numbers: []string{
			"1", "2", "3", "4", "5", "6", "7", "8", "9", "0",
		},
	}
}

func (d *Day3) Challenge1() {
	fmt.Println("Day 3 - Challenge 1")
	total := 0
	nbuilder := ""
	var isnum, is_part bool
	surround := make(map[string]string)

	for i := 0; i < len(d.lines); i++ { // loop rows
		for j := 0; j < len(d.lines[i]); j++ { // loop cols
			cur := string(d.lines[i][j])
			isnum = slices.Contains(d.numbers, cur)
			if isnum { // Not my best work.
				nbuilder += cur
				if j == 0 {
					if i == 0 {
						surround["r"] = string(d.lines[i][j+1])
						surround["br"] = string(d.lines[i+1][j+1])
						surround["b"] = string(d.lines[i+1][j])
					} else if i == len(d.lines)-1 {
						surround["t"] = string(d.lines[i-1][j])
						surround["tr"] = string(d.lines[i-1][j+1])
						surround["r"] = string(d.lines[i][j+1])
					} else {
						surround["t"] = string(d.lines[i-1][j])
						surround["tr"] = string(d.lines[i-1][j+1])
						surround["r"] = string(d.lines[i][j+1])
						surround["br"] = string(d.lines[i+1][j+1])
						surround["b"] = string(d.lines[i+1][j])
					}
				} else if j == len(d.lines[i])-1 {
					if i == 0 {
						surround["l"] = string(d.lines[i][j-1])
						surround["bl"] = string(d.lines[i+1][j-1])
						surround["b"] = string(d.lines[i+1][j])
					} else if i == len(d.lines)-1 {
						surround["t"] = string(d.lines[i-1][j])
						surround["tl"] = string(d.lines[i-1][j-1])
						surround["l"] = string(d.lines[i][j-1])
					} else {
						surround["t"] = string(d.lines[i-1][j])
						surround["tl"] = string(d.lines[i-1][j-1])
						surround["l"] = string(d.lines[i][j-1])
						surround["bl"] = string(d.lines[i+1][j-1])
						surround["b"] = string(d.lines[i+1][j])
					}
				} else {
					if i == 0 {
						surround["l"] = string(d.lines[i][j-1])
						surround["bl"] = string(d.lines[i+1][j-1])
						surround["b"] = string(d.lines[i+1][j])
						surround["br"] = string(d.lines[i+1][j+1])
						surround["r"] = string(d.lines[i][j+1])
					} else if i == len(d.lines)-1 {
						surround["l"] = string(d.lines[i][j-1])
						surround["tl"] = string(d.lines[i-1][j-1])
						surround["t"] = string(d.lines[i-1][j])
						surround["tr"] = string(d.lines[i-1][j+1])
						surround["r"] = string(d.lines[i][j+1])
					} else {
						surround["l"] = string(d.lines[i][j-1])
						surround["tl"] = string(d.lines[i-1][j-1])
						surround["t"] = string(d.lines[i-1][j])
						surround["tr"] = string(d.lines[i-1][j+1])
						surround["r"] = string(d.lines[i][j+1])
						surround["br"] = string(d.lines[i+1][j+1])
						surround["b"] = string(d.lines[i+1][j])
						surround["bl"] = string(d.lines[i+1][j-1])
					}
				}

				for _, v := range surround {
					if _, err := strconv.Atoi(v); err != nil && v != "." && v != "" {
						is_part = true
						break
					}
				}

			} else {
				num, _ := strconv.Atoi(nbuilder)
				if is_part {
					total += num
				}
				is_part = false
				nbuilder = ""
			}

			for k, _ := range surround {
				surround[k] = ""
			}
		}
	}

	fmt.Printf("number: %d\n", total)
}

func (d *Day3) Challenge2() {
	fmt.Println("Day 3 - Challenge 2")
	total := 0
	gearBuilder := [2]string{"", ""}
	var isnum, is_part bool
	surround := make(map[string][2]string)
	gears := make(map[int]map[int][]int)

	for i := 0; i < len(d.lines); i++ { // loop rows
		for j := 0; j < len(d.lines[i]); j++ { // loop cols
			cur := string(d.lines[i][j])
			isnum = slices.Contains(d.numbers, cur)
			if isnum { // Not my best work.
				gearBuilder[0] += cur
				if j == 0 {
					if i == 0 {
						surround["r"] = [2]string{string(d.lines[i][j+1]), fmt.Sprintf("%d-%d", i, j+1)}
						surround["br"] = [2]string{string(d.lines[i+1][j+1]), fmt.Sprintf("%d-%d", i+1, j+1)}
						surround["b"] = [2]string{string(d.lines[i+1][j]), fmt.Sprintf("%d-%d", i+1, j)}
					} else if i == len(d.lines)-1 {
						surround["t"] = [2]string{string(d.lines[i-1][j]), fmt.Sprintf("%d-%d", i-1, j)}
						surround["tr"] = [2]string{string(d.lines[i-1][j+1]), fmt.Sprintf("%d-%d", i-1, j+1)}
						surround["r"] = [2]string{string(d.lines[i][j+1]), fmt.Sprintf("%d-%d", i, j+1)}
					} else {
						surround["t"] = [2]string{string(d.lines[i-1][j]), fmt.Sprintf("%d-%d", i-1, j)}
						surround["tr"] = [2]string{string(d.lines[i-1][j+1]), fmt.Sprintf("%d-%d", i-1, j+1)}
						surround["r"] = [2]string{string(d.lines[i][j+1]), fmt.Sprintf("%d-%d", i, j+1)}
						surround["br"] = [2]string{string(d.lines[i+1][j+1]), fmt.Sprintf("%d-%d", i+1, j+1)}
						surround["b"] = [2]string{string(d.lines[i+1][j]), fmt.Sprintf("%d-%d", i+1, j)}
					}
				} else if j == len(d.lines[i])-1 {
					if i == 0 {
						surround["l"] = [2]string{string(d.lines[i][j-1]), fmt.Sprintf("%d-%d", i, j-1)}
						surround["bl"] = [2]string{string(d.lines[i+1][j-1]), fmt.Sprintf("%d-%d", i+1, j-1)}
						surround["b"] = [2]string{string(d.lines[i+1][j]), fmt.Sprintf("%d-%d", i+1, j)}
					} else if i == len(d.lines)-1 {
						surround["t"] = [2]string{string(d.lines[i-1][j]), fmt.Sprintf("%d-%d", i-1, j)}
						surround["tl"] = [2]string{string(d.lines[i-1][j-1]), fmt.Sprintf("%d-%d", i-1, j-11)}
						surround["l"] = [2]string{string(d.lines[i][j-1]), fmt.Sprintf("%d-%d", i, j-1)}
					} else {
						surround["t"] = [2]string{string(d.lines[i-1][j]), fmt.Sprintf("%d-%d", i-1, j)}
						surround["tl"] = [2]string{string(d.lines[i-1][j-1]), fmt.Sprintf("%d-%d", i-1, j-1)}
						surround["l"] = [2]string{string(d.lines[i][j-1]), fmt.Sprintf("%d-%d", i, j-1)}
						surround["bl"] = [2]string{string(d.lines[i+1][j-1]), fmt.Sprintf("%d-%d", i+1, j-1)}
						surround["b"] = [2]string{string(d.lines[i+1][j]), fmt.Sprintf("%d-%d", i+1, j)}
					}
				} else {
					if i == 0 {
						surround["l"] = [2]string{string(d.lines[i][j-1]), fmt.Sprintf("%d-%d", i, j-1)}
						surround["bl"] = [2]string{string(d.lines[i+1][j-1]), fmt.Sprintf("%d-%d", i+1, j-1)}
						surround["b"] = [2]string{string(d.lines[i+1][j]), fmt.Sprintf("%d-%d", i+1, j)}
						surround["br"] = [2]string{string(d.lines[i+1][j+1]), fmt.Sprintf("%d-%d", i+1, j+1)}
						surround["r"] = [2]string{string(d.lines[i][j+1]), fmt.Sprintf("%d-%d", i, j+1)}
					} else if i == len(d.lines)-1 {
						surround["l"] = [2]string{string(d.lines[i][j-1]), fmt.Sprintf("%d-%d", i, j-1)}
						surround["tl"] = [2]string{string(d.lines[i-1][j-1]), fmt.Sprintf("%d-%d", i-1, j-1)}
						surround["t"] = [2]string{string(d.lines[i-1][j]), fmt.Sprintf("%d-%d", i-1, j)}
						surround["tr"] = [2]string{string(d.lines[i-1][j+1]), fmt.Sprintf("%d-%d", i-1, j+1)}
						surround["r"] = [2]string{string(d.lines[i][j+1]), fmt.Sprintf("%d-%d", i, j+1)}
					} else {
						surround["l"] = [2]string{string(d.lines[i][j-1]), fmt.Sprintf("%d-%d", i, j-1)}
						surround["tl"] = [2]string{string(d.lines[i-1][j-1]), fmt.Sprintf("%d-%d", i-1, j-1)}
						surround["t"] = [2]string{string(d.lines[i-1][j]), fmt.Sprintf("%d-%d", i-1, j)}
						surround["tr"] = [2]string{string(d.lines[i-1][j+1]), fmt.Sprintf("%d-%d", i-1, j+1)}
						surround["r"] = [2]string{string(d.lines[i][j+1]), fmt.Sprintf("%d-%d", i, j+1)}
						surround["br"] = [2]string{string(d.lines[i+1][j+1]), fmt.Sprintf("%d-%d", i+1, j+1)}
						surround["b"] = [2]string{string(d.lines[i+1][j]), fmt.Sprintf("%d-%d", i+1, j)}
						surround["bl"] = [2]string{string(d.lines[i+1][j-1]), fmt.Sprintf("%d-%d", i+1, j-1)}
					}
				}

				for _, v := range surround {
					if v[0] == "*" {
						is_part = true
						if gearBuilder[1] == "" {
							gearBuilder[1] = v[1]
						}

						break
					}
				}

			} else {
				num, _ := strconv.Atoi(gearBuilder[0])
				if is_part {
					gindexes := strings.Split(gearBuilder[1], "-")
					grow, _ := strconv.Atoi(gindexes[0])
					gcol, _ := strconv.Atoi(gindexes[1])
					if _, ok := gears[grow]; !ok {
						gears[grow] = make(map[int][]int)
					}
					gears[grow][gcol] = append(gears[grow][gcol], num)
				}
				is_part = false
				gearBuilder[0] = ""
				gearBuilder[1] = ""
			}

			for k, _ := range surround {
				surround[k] = [2]string{}
			}
		}
	}
	for _, v := range gears {
		for _, nums := range v {
			if len(nums) == 2 {
				total += nums[0] * nums[1]
			}
		}
	}

	fmt.Printf("number: %d\n", total)
}
