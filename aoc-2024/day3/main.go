package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func isSym(val rune) bool {
	return !unicode.IsDigit(val) && val != '.'
}

func isPart(i, j, rl, cl int, mx [][]rune) bool {
	if i > 0 && isSym(mx[i-1][j]) {
		return true
	}
	if i < rl-1 && !unicode.IsDigit(mx[i+1][j]) && mx[i+1][j] != '.' {
		return true
	}
	if j < cl-1 && !unicode.IsDigit(mx[i][j+1]) && mx[i][j+1] != '.' {
		return true
	}
	if j > 0 && !unicode.IsDigit(mx[i][j-1]) && mx[i][j-1] != '.' {
		return true
	}
	if i < rl-1 && j < cl-1 && !unicode.IsDigit(mx[i+1][j+1]) && mx[i+1][j+1] != '.' {
		return true
	}
	if i > 0 && j > 0 && !unicode.IsDigit(mx[i-1][j-1]) && mx[i-1][j-1] != '.' {
		return true
	}
	if i > 0 && j < cl-1 && !unicode.IsDigit(mx[i-1][j+1]) && mx[i-1][j+1] != '.' {
		return true
	}
	if i < rl-1 && j > 0 && !unicode.IsDigit(mx[i+1][j-1]) && mx[i+1][j-1] != '.' {
		return true
	}
	return false
}

func findNum(i, j int, mx [][]rune) (int, string) {
	cl := len(mx[0])
	for j > 0 && unicode.IsNumber(mx[i][j-1]) {
		j -= 1
	}

	ds := []rune{}
	initj := j
	for j < cl && unicode.IsNumber(mx[i][j]) {
		ds = append(ds, mx[i][j])
		j += 1
	}
	n, _ := strconv.Atoi(string(ds))
	return n, fmt.Sprintf("%d-%d", i, initj)
}

func part1(rl, cl int, mx [][]rune) {
	nums := []int{}
	tn := []rune{}
	isP := false
	for i := 0; i < rl; i++ {
		for j := 0; j < cl; j++ {
			if unicode.IsNumber(mx[i][j]) {
				tn = append(tn, mx[i][j])
				if isPart(i, j, rl, cl, mx) {
					isP = true
				}
			} else {
				if len(tn) > 0 {
					n, _ := strconv.Atoi(string(tn))
					if isP {
						nums = append(nums, n)
					}
				}
				tn = []rune{}
				isP = false
			}
		}
		if len(tn) > 0 {
			n, _ := strconv.Atoi(string(tn))
			if isP {
				nums = append(nums, n)
			}
		}
		tn = []rune{}
		isP = false
	}

	res := 0
	for _, n := range nums {
		res += n
	}
	fmt.Printf("Ans is %d\n", res)
}

func part2(rl, cl int, mx [][]rune) {
	gears := []int{}
	gear_idx := [][]int{}
	for i := 0; i < rl; i++ {
		for j := 0; j < cl; j++ {
			if isSym(mx[i][j]) {
				if i > 0 && unicode.IsNumber(mx[i-1][j]) {
					gear_idx = append(gear_idx, []int{i - 1, j})
				}
				if j > 0 && unicode.IsNumber(mx[i][j-1]) {
					gear_idx = append(gear_idx, []int{i, j - 1})
				}
				if i < rl && unicode.IsNumber(mx[i+1][j]) {
					gear_idx = append(gear_idx, []int{i + 1, j})
				}
				if j < cl && unicode.IsNumber(mx[i][j+1]) {
					gear_idx = append(gear_idx, []int{i, j + 1})
				}
				if i > 0 && j > 0 && unicode.IsNumber(mx[i-1][j-1]) {
					gear_idx = append(gear_idx, []int{i - 1, j - 1})
				}
				if i < rl && j < cl && unicode.IsNumber(mx[i+1][j+1]) {
					gear_idx = append(gear_idx, []int{i + 1, j + 1})
				}
				if i > 0 && j < cl && unicode.IsNumber(mx[i-1][j+1]) {
					gear_idx = append(gear_idx, []int{i - 1, j + 1})
				}
				if i < rl && j > 0 && unicode.IsNumber(mx[i+1][j-1]) {
					gear_idx = append(gear_idx, []int{i + 1, j - 1})
				}

				if len(gear_idx) > 0 {
					ns := []int{}
					visited := map[string]bool{}

					for i, _ := range gear_idx {
						n, b := findNum(gear_idx[i][0], gear_idx[i][1], mx)
						_, ok := visited[b]
						if !ok {
							ns = append(ns, n)
						}
						visited[b] = true
					}
					if len(ns) == 2 {
						gears = append(gears, ns[0]*ns[1])
					}

					gear_idx = [][]int{}
				}
			}
		}
	}
	s := 0
	for _, g := range gears {
		s += g
	}
	fmt.Printf("Ans is %d\n", s)
}

func main() {
	fc, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Error while reading the file %v\n", err)
	}
	cn := string(fc)

	rows := strings.Split(cn, "\n")
	rows = rows[:len(rows)-1]
	rl := len(rows)
	cl := len(rows[0])

	mx := make([][]rune, rl)
	for i := 0; i < rl; i++ {
		mx[i] = make([]rune, cl)
	}

	for i, row := range rows {
		for j, val := range row {
			mx[i][j] = val
		}
	}

	part2(rl, cl, mx)
}
