package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	fishes     []int
	fishesAlgo []int
)

func getInput(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	fishes = make([]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		tmpList := strings.Split(line, ",")
		for _, val := range tmpList {
			val, _ := strconv.Atoi(val)
			fishes = append(fishes, val)
		}
		//	fmt.Println(fishes)
	}
}

func spawn(days int) {
	// not even used after I found spawnAlgo
	for day := 1; day < days+1; day++ {
		tmpFishes := make([]int, 0)
		for pos, val := range fishes {
			switch val {
			case 0:
				fishes[pos] = 6
				tmpFishes = append(tmpFishes, 8)
			default:
				fishes[pos]--
			}

		}
		fmt.Printf("Day %d: ", day)
		fishes = append(fishes, tmpFishes...)
		fmt.Printf("%d\n", len(fishes))
	}
}

func populate(fishList []int) (initialState []int) {
	initialState = make([]int, 9)
	for _, val := range fishList {
		initialState[val]++
	}
	return initialState
}

func spawnAlgo(fishList []int, days int) {
	for day := 1; day < days+1; day++ {
		//tmpFishes := make([]int, 9)
		fishList = append(fishList, fishList[0])
		fishList = fishList[1:]
		fishList[6] = fishList[6] + fishList[8]
	}
	sum := 0
	for _, val := range fishList {
		sum = sum + val
	}
	fmt.Printf("The Answer for %d days is: %d\n", days, sum)

}

func main() {
	f := flag.String("f", "t", "")
	flag.Parse()
	if string(*f) == "t" {
		getInput("../test_input.txt")
	} else {
		getInput("../input.txt")
	}
	fishesAlgo = populate(fishes)
	spawnAlgo(fishesAlgo, 80)
	spawnAlgo(fishesAlgo, 256)
}
