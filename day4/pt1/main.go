package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	boardNum     int                // keeps count of the number of boards
	bingoCalls   []int              // the bingo calls
	bingoBoards2 map[string][][]int // another approach to the bingoBoards
)

func createBingoCalls(calls string) (bC []int) {
	tmp := strings.Split(calls, ",")
	values := make([]int, 0, len(tmp))
	for _, raw := range tmp {
		v, err := strconv.Atoi(raw)
		if err != nil {
			log.Print(err)
			continue
		}
		values = append(values, v)
	}
	//fmt.Println(values)
	return values
}

func createRowsColumns(line string) (br []int) {
	tmp := strings.Fields(line)
	values := make([]int, 0, len(tmp))
	for _, raw := range tmp {
		v, err := strconv.Atoi(raw)
		if err != nil {
			log.Print(err)
			continue
		}
		values = append(values, v)
	}
	//	fmt.Println(values)
	return values
}

func createColumns(bB2 map[string][][]int) (wc map[string][][]int) {
	for key, rows := range bB2 {
		tmpslice := make([]int, 0)
		//		fmt.Println(key)
		//		fmt.Printf("This is a row: %v\n", rows)
		for i := 0; i < len(rows); i++ {
			for k, _ := range rows {
				tmpslice = append(tmpslice, rows[k][i])
			}
			bB2[key] = append(bB2[key], tmpslice)
			//			fmt.Println(tmpslice)
			tmpslice = nil
		}
	}
	return bB2
}

func sumSquares(winningBoard [][]int, winningCall int) {
	// clean out doubles
	m := make(map[int]bool)

	for rCCnt := 0; rCCnt < (len(winningBoard)); rCCnt++ {
		for _, r := range winningBoard[rCCnt] {
			m[r] = true
		}
	}
	fmt.Println(m)
	total := 0

	for k, _ := range m {
		total = total + k
	}

	total = total - winningCall
	fmt.Println(total)
	fmt.Printf("the magic answer is: %d\n", total*winningCall)
}

func playBingo(bB map[string][][]int, bCalls []int) {

	for call := 0; call < len(bCalls); call++ {
		for i := 0; i < len(bB); i++ {
			bd := fmt.Sprintf("Board%d", i)
			//fmt.Println(bB[bd])
			for rCCnt := 0; rCCnt < (len(bB[bd])); rCCnt++ {
				//fmt.Println(bB[bd][rCCnt])
				// tmp := bB[bd][rCCnt]
				for plats := 0; plats < (len(bB[bd][rCCnt])); plats++ {
					//fmt.Printf("compare call: %d to square: %d\n", bCalls[call], bB[bd][rCCnt][plats])
					if bCalls[call] == bB[bd][rCCnt][plats] {
						bB[bd][rCCnt] = append(bB[bd][rCCnt][:plats], bB[bd][rCCnt][plats+1:]...)
						//fmt.Println(bB[bd][rCCnt])
						if len(bB[bd][rCCnt]) == 0 {
							fmt.Printf("Bingo mf with call: %d  and board: %v\n", bCalls[call], bB[bd])
							sumSquares(bB[bd], bCalls[call])
							log.Fatal()
						}
					}
				}
			}
		}
	}
}

func main() {
	f, err := os.Open("../test_input.txt")
	//f, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	firstTime := false
	boardNum = 0
	bd := ""
	bingoBoards2 = make(map[string][][]int)
	bingoCalls = make([]int, 0)

	for scanner.Scan() {
		line := scanner.Text()
		if !firstTime {
			// Get the bingo calls
			//fmt.Printf("First run: %s\n", line)
			firstTime = true
			bingoCalls = createBingoCalls(line)
			continue
		}
		if line == "" {
			// create empty board
			bd = fmt.Sprintf("Board%d", boardNum)
			bingoBoards2[bd] = make([][]int, 0)
			boardNum++
			continue
		}
		//		fmt.Println(line)

		row := createRowsColumns(line)
		//fmt.Println(row)
		bingoBoards2[bd] = append(bingoBoards2[bd], row)

	}
	boardNum = len(bingoBoards2)
	fmt.Printf("number of bingo boards: %d\n", boardNum)

	bingoBoards2 = createColumns(bingoBoards2)
	//fmt.Println(bingoBoards2)
	playBingo(bingoBoards2, bingoCalls)
}
