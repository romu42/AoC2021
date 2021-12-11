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
	boardNum    int
	bingoCalls  []int
	boardList   []string
	bingoBoards map[string][][]int
	markerBoard map[string][][]string
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

func createBingoBoard(line string) (br []int) {
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

//func printBingoBoards(bB map[string][][]int, mB map[string][][]string) {
//	for i := 0; i < len(bB); i++ {
//		bd := fmt.Sprintf("Board%d", i)
//		mbd := fmt.Sprintf("MBoard%d", i)
//		fmt.Printf("Board %d\n", i)
//		for row := 0; row < 5; row++ {
//			for col := 0; col < 5; col++ {
//				fmt.Printf("%2d%s ", bB[bd][row][col], mB[mbd][row][col])
//			}
//			fmt.Println()
//		}
//		fmt.Println()
//		fmt.Println()
//	}
//}

func createMarkerBoard(numBoards int) (mB map[string][][]string) {
	mB = make(map[string][][]string)
	for i := 0; i < numBoards; i++ {
		for row := 0; row < 5; row++ {
			mb := fmt.Sprintf("MBoard%d", i)
			values := []string{"-", "-", "-", "-", "-"}
			mB[mb] = append(mB[mb], values)
		}
		//		fmt.Println(mB)
	}

	return mB
}

func main() {
	//f, err := os.Open("../test_input.txt")
	f, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	firstTime := false
	boardNum = 0
	bd := ""
	bingoBoards = make(map[string][][]int)
	bingoCalls = make([]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if !firstTime {
			firstTime = true
			bingoCalls = createBingoCalls(line)
			continue
		}
		if line == "" {
			// create empty board
			bd = fmt.Sprintf("Board%d", boardNum)
			// add board to list
			boardList = append(boardList, bd)
			bingoBoards[bd] = make([][]int, 0)
			boardNum++
			continue
		}
		//		fmt.Println(line)

		row := createBingoBoard(line)
		bingoBoards[bd] = append(bingoBoards[bd], row)

	}

	boardNum = len(bingoBoards)
	fmt.Printf("number of bingo boards: %d\n", boardNum)
	fmt.Println(boardList)
	markerBoard = createMarkerBoard(boardNum)
	fmt.Println(markerBoard)

	for _, call := range bingoCalls {
		if true {
			markBoard(boardList, call)
			//		printBingoBoards(boardList)
			winners := checkForBingo(boardList)
			if len(winners) > 0 {
				fmt.Println("Found  Winner!")
				for _, winBoard := range winners {
					fmt.Printf("board %s\n", winBoard)
					score := calcScore(winBoard)
					fmt.Printf("Winning score is: %d\nOn board: %s\nWith call %d\n", score*call, winBoard, call)
					boardList = cleanBoardList(winBoard)
					if len(boardList) == 0 {
						log.Fatal("no more boards!")
					}
				}
			}
		} else {
			break
		}
	}
}

func cleanBoardList(winner string) (tmpList []string) {
	tmpList = make([]string, 0)
	for _, key := range boardList {
		if key != winner {
			tmpList = append(tmpList, key)
		}
	}
	return tmpList
}

func calcScore(winner string) (score int) {
	score = 0
	fmt.Printf("Calculating score %s\n", winner)
	bd := winner
	mbd := "M" + winner
	for row := 0; row < 5; row++ {
		for col := 0; col < 5; col++ {
			if markerBoard[mbd][row][col] == "-" {
				score = score + bingoBoards[bd][row][col]
			}

		}
	}

	return score
}

func checkForBingo(bdList []string) (winners []string) {
	winners = make([]string, 0)
	for _, key := range bdList {
		bd := key
		mbd := "M" + key
		// Check Rows
		for row := 0; row < 5; row++ {
			rowCnt := 0
			for col := 0; col < 5; col++ {
				if markerBoard[mbd][row][col] == "x" {
					rowCnt++
				}
			}
			if rowCnt == 5 {
				winners = append(winners, bd)
			}
		}
		// Check Columns
		for col := 0; col < 5; col++ {
			colCnt := 0
			for row := 0; row < 5; row++ {
				if markerBoard[mbd][row][col] == "x" {
					colCnt++
				}
			}
			if colCnt == 5 {
				winners = append(winners, bd)
			}
		}
	}
	return winners
}

func printBingoBoards(bdList []string) {
	for _, key := range bdList {
		bd := key
		mbd := "M" + key
		fmt.Printf("Board %s\n", key)
		for row := 0; row < 5; row++ {
			for col := 0; col < 5; col++ {
				fmt.Printf("%2d%s ", bingoBoards[bd][row][col], markerBoard[mbd][row][col])
			}
			fmt.Println()
		}
		fmt.Println()
		fmt.Println()
	}
}

func markBoard(bdList []string, bCall int) {
	for _, key := range bdList {
		bd := key
		mbd := "M" + key
		for row := 0; row < 5; row++ {
			for col := 0; col < 5; col++ {
				if bingoBoards[bd][row][col] == bCall {
					markerBoard[mbd][row][col] = "x"
				}
			}
		}
	}
}
