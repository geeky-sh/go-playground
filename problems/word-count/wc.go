package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	var lineOp = flag.Bool("l", false, "This option returns no. of lines")
	var wordOp = flag.Bool("w", false, "This option returns no. of words")
	var charOp = flag.Bool("c", false, "This option returns no. of characters")
	// var multiChars = flag.Bool("m", false, "This option represents no. of multi-byte characters")
	flag.Parse()

	var allOp = false
	if !*lineOp && !*wordOp && !*charOp {
		allOp = true
	}

	var f *os.File
	var err error
	fp := flag.Arg(0)
	if fp == "" {
		f = os.Stdin
	} else {
		f, err = os.Open(fp)
		if err != nil {
			log.Fatalf("Error while opening the file %v", err)
		}
	}
	defer f.Close()

	// if *lineOp || allOp {
	// 	sc := bufio.NewScanner(f)
	// 	sc.Split(bufio.ScanLines)
	// 	lineAns := 0
	// 	for sc.Scan() {
	// 		lineAns += 1
	// 	}

	// 	fmt.Printf("\t%d", lineAns)
	// 	f.Seek(0, 0)
	// }

	// if *wordOp || allOp {
	// 	sc := bufio.NewScanner(f)
	// 	sc.Split(bufio.ScanWords)
	// 	wordAns := 0
	// 	for sc.Scan() {
	// 		wordAns += 1
	// 	}

	// 	fmt.Printf("\t%d", wordAns)
	// 	f.Seek(0, 0)
	// }

	// if *charOp || allOp {
	// 	sc := bufio.NewScanner(f)
	// 	sc.Split(bufio.ScanBytes)
	// 	charAns := 0
	// 	for sc.Scan() {
	// 		charAns += 1
	// 	}

	// 	fmt.Printf("\t%d", charAns)
	// }

	// fmt.Printf(" %s\n", fp)

	lc := 0
	wc := 0
	cc := 0
	for {
		bs := make([]byte, 64)
		c, err := f.Read(bs)
		if err != nil {
			log.Fatalf("Error reading file %v\n", err)
		}
		cc += c

		for _, b := range bs {
			if b == byte(' ') {
				wc += 1
			}

			// it is observed that last line of the file is always newline
			if b == byte('\n') {
				lc += 1
			}
		}
		if c < 64 {
			wc += 1
			break
		}
	}

	if *lineOp || allOp {
		fmt.Printf("\t%d", lc)
	}
	if *wordOp || allOp {
		fmt.Printf("\t%d", wc)
	}
	if *charOp || allOp {
		fmt.Printf("\t%d", cc)
	}

	fmt.Printf(" %s\n", fp)
}
