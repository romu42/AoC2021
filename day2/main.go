package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

var (
	pt1Vp int      // pt1 vertical position
	pt1Hp int      // pt1 horizontal position
	pt2Vp int      // pt2 vertical position
	pt2Hp int      // pt2 horizontal position
	mov   []string // movement direction, distance
	aim   int      // depth adjuster
)

func pt1(dir string, dis int) {
	switch dir {
	case "forward":
		pt1Hp = pt1Hp + dis
	case "down":
		pt1Vp = pt1Vp + dis
	case "up":
		pt1Vp = pt1Vp - dis
	}
}

func pt2(dir string, dis int) {
	switch dir {
	case "forward":
		pt2Hp = pt2Hp + dis
		pt2Vp = pt2Vp + (dis * aim)
	case "down":
		aim = aim + dis
	case "up":
		aim = aim - dis
	}
}

func main() {
	//f, err := os.Open("test_input.txt")
	f, err := os.Open("input.txt")
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		mov = strings.Split(line, " ")
		dir := mov[0]
		dis, err := strconv.Atoi(mov[1])
		check(err)
		pt1(dir, dis)
		pt2(dir, dis)
	}

	fmt.Printf("part1: %d\n", (pt1Hp * pt1Vp))
	fmt.Printf("part2: %d\n", (pt2Hp * pt2Vp))
}
