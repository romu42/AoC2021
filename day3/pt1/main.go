package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var (
	// put your variables in here
	lineL   int     // line length
	slices  [][]int // slice of slices to figure out MCB and LCB
	oneTime bool    // one time run flag
	MCB     string  // most common bytes
	LCB     string  // least common bytes
)

func createSliceOfSlices(length int) {
	slices = make([][]int, length)
	for i := 0; i < length; i++ {
		slices[i] = make([]int, 2)
	}
}

func main() {

	//f, err := os.Open("../test_input.txt")
	f, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	oneTime = false
	for scanner.Scan() {
		line := scanner.Text()
		if !oneTime {
			oneTime = true
			lineL = (len(line))
			createSliceOfSlices(lineL)
		}

		for pos, char := range line {
			switch char {
			// 0 in Runish
			case 48:
				slices[pos][0]++
			// 1 in Runish
			case 49:
				slices[pos][1]++
			}
		}
	}

	lcb_str := ""
	mcb_str := ""

	for i := 0; i < lineL; i++ {
		cnt0 := slices[i][0]
		cnt1 := slices[i][1]
		if cnt0 > cnt1 {
			lcb_str = lcb_str + "1"
			mcb_str = mcb_str + "0"
		} else {
			lcb_str = lcb_str + "0"
			mcb_str = mcb_str + "1"
		}

	}

	MCB, err := strconv.ParseInt(mcb_str, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	LCB, err := strconv.ParseInt(lcb_str, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(MCB)
	fmt.Println(LCB)
	fmt.Println(MCB * LCB)
}
