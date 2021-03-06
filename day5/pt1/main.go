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
	input map[int][][]int
)

//func getInput(input map[int][][]int) (output map[int][][]int) {
func getInput() {
	f, err := os.Open("../input.txt")
	//f, err := os.Open("../test_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	ansmap := make(map[string]int)
	for scanner.Scan() {
		line := scanner.Text()
		tmp := strings.Fields(line)
		coord := getCoordinates(tmp)
		//fmt.Printf("org coordinates: %v\n", coord)
		coordPath := makePath(coord)
		if len(coordPath) > 0 {
			fmt.Printf("made: %v %T\n", coordPath, coordPath)
			fmt.Println("<---------------------------------------->")
			for _, xy := range coordPath {
				key := strconv.Itoa(xy[0]) + "-" + strconv.Itoa(xy[1])
				_, ok := ansmap[key]
				if ok {
					ansmap[key] = ansmap[key] + 1
				} else {
					ansmap[key] = 1
				}

			}
			//	} else {
			//		fmt.Printf("made: %v\n", coordPath)
			//		fmt.Println("<---------------------------------------->")
		}
	}

	getPt1(ansmap)
	//	return output
}

func getCoordinates(xys []string) (pathCoord [][]int) {
	pathCoord = make([][]int, 0)
	xy1 := strings.Split(xys[0], ",")
	xy2 := strings.Split(xys[2], ",")
	//fmt.Printf("pt1: %s -> pt2: %s\n", xy1, xy2)
	//	fmt.Printf("can i just get rid of -> %s\n", strings.Replace(line, " ->", "", 1))
	x1, _ := strconv.Atoi(xy1[0])
	y1, _ := strconv.Atoi(xy1[1])
	x2, _ := strconv.Atoi(xy2[0])
	y2, _ := strconv.Atoi(xy2[1])
	pathCoord = append(pathCoord, []int{x1, y1}, []int{x2, y2})
	fmt.Printf("pathCoord: %v\n", pathCoord)
	return pathCoord
}

func makePath(xys [][]int) (path [][]int) {
	x1 := xys[0][0]
	y1 := xys[0][1]
	x2 := xys[1][0]
	y2 := xys[1][1]
	path = make([][]int, 0)
	xyTmp := make([]int, 0)

	if x1 == x2 {
		if y1 > y2 {
			// 2,2 -> 2,1 works
			for i := y1; i > y2-1; i-- {
				xyTmp = append(xyTmp, i)
			}
		} else if y1 < y2 {
			// 7,0 -> 7,4 works
			for i := y1; i < y2+1; i++ {
				xyTmp = append(xyTmp, i)
			}
		}
		for _, val := range xyTmp {
			// add all the y values
			path = append(path, []int{x1, val})
		}
	} else if y1 == y2 {
		// 9,4 -> 3,4 works
		if x1 > x2 {
			for i := x1; i > x2-1; i-- {
				xyTmp = append(xyTmp, i)
			}
		} else if x1 < x2 {
			// 0,9 -> 5,9 works
			for i := x1; i < x2+1; i++ {
				xyTmp = append(xyTmp, i)
			}
		}
		for _, val := range xyTmp {
			// add all the x values
			path = append(path, []int{val, y1})
		}
	} else if x1 > x2 {
		if y1 > y2 {
			// 9,9 -> 6,6
			y := y1
			for x := x1; x > x2-1; x-- {
				//for y := y1; y > y2-1; y-- {
				path = append(path, []int{x, y})
				y--
				//	}
			}
		} else if y1 < y2 {
			// 9,7 -> 7,9
			y := y1
			for x := x1; x > x2-1; x-- {
				//	for y := y1; y < y2+1; y++ {
				path = append(path, []int{x, y})
				y++
				//	}
			}
		}
	} else if x1 < x2 {
		if y1 > y2 {
			// 6,9 -> 9,6
			y := y1
			for x := x1; x < x2+1; x++ {
				//	for y := y1; y > y2-1; y-- {
				path = append(path, []int{x, y})
				y--
				//}
			}
		} else if y1 < y2 {
			// 1,1 -> 3,3
			y := y1
			for x := x1; x < x2+1; x++ {
				//	for y := y1; y < y2+1; y++ {
				path = append(path, []int{x, y})
				y++
				//}
			}
		}
	}
	return path
}

func getPt1(pathcnt map[string]int) {
	fmt.Println(pathcnt)
	counter := 0
	for _, val := range pathcnt {
		if val > 1 {
			counter++
		}
	}
	fmt.Printf("The magic answer is: %d\n", counter)
}

func main() {
	getInput()
}
