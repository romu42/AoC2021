package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

var (
	pM   int   // previous measurement
	m    int   // messurement
	cnt  int   // counter
	w    []int // moving window
	wcnt int   // window counter
)

func main() {

	f, err := os.Open("input.txt")
	//f, err := os.Open("test_input.txt")
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	pM = 0
	cnt = 0
	m = 0
	w = make([]int, 0, 6)

	for scanner.Scan() {
		m, err = strconv.Atoi(scanner.Text())
		check(err)
		if pM == 0 {
			pM = m
		}
		if pM < m {
			cnt++
		}
		pM = m

		// initiate our sliding windows
		if len(w) < 4 {
			w = append(w, m)

		} else {
			// count the windows that the second window is greater than the first
			w = append(w, m)
			if (w[0] + w[1] + w[2]) < (w[1] + w[2] + w[3]) {
				wcnt++
			}
			w = w[1:]
		}
		// get the last windows the windows that the second window is greater than the first
		// probably a smarter way to do this, but I am not that smart yet
		for len(w) > 3 {
			if (w[0] + w[1] + w[2]) < (w[1] + w[2] + w[3]) {
				wcnt++
			}
			w = w[1:]
		}

	}

	fmt.Println(cnt)
	fmt.Println(wcnt)

	check(err)
}
