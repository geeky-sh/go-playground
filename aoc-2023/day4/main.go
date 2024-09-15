package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("./small-input.txt")
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
		return
	}
	defer f.Close()

	rd := bufio.NewReader(f)
	res := 0
	r := 0
	for {
		l, err := rd.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}

			log.Fatalf("Error reading line: %v", err)
			return
		}

		ss := strings.Split(l, ":")
		vs := ss[1]

		numL := strings.Split(vs, "|")
		winC := strings.Split(numL[0], " ")
		myC := strings.Split(numL[1], " ")

		winMap := map[string]int{}
		fmt.Println(winC)
		for _, n := range winC {
			if n == "" {
				continue
			}
			_, ok := winMap[n]
			if !ok {
				winMap[n] = 0
			}
			winMap[n] += 1
		}

		r = 0
		fmt.Println(winMap)
		for _, n := range myC {
			if n == "" {
				continue
			}
			_, ok := winMap[n]
			if ok {
				r += 1
			}
		}
		fmt.Printf("Card has %d win combinations\n", r)
		if r > 0 {
			fmt.Println(int(math.Pow(2, float64(r-1))))
			res += int(math.Pow(2, float64(r-1)))
		}
		fmt.Println(res)
	}

	fmt.Printf("Answer is %v\n", res)
}
