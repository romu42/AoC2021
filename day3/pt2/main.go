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
	lineL     int     // line length
	slicesOgr [][]int // slice of slices for Oxygen generator rating
	slicesCsr [][]int // slice of slices for CO2 scrubber rating
	oneTime   bool    // one time run flag
	counter   int     // just a counter
)

func updateSlicesAll(input string) {
	// Build Slice of slices to find the OGR and CSR
	s0 := make([]int, 0)
	for _, val := range input {
		s0 = append(s0, int(val-'0'))
	}
	slicesOgr = append(slicesOgr, s0)
	slicesCsr = append(slicesCsr, s0)
}

func getOgr(i int) {
	zero := 0
	one := 0
	zeroSlice := make([][]int, 0)
	oneSlice := make([][]int, 0)
	for pos, val := range slicesOgr {
		//fmt.Println(val[i])
		switch val[i] {
		case 0:
			zero++
			zeroSlice = append(zeroSlice, slicesOgr[pos])
		case 1:
			one++
			oneSlice = append(oneSlice, slicesOgr[pos])
		}
	}
	//fmt.Println(zero)
	//fmt.Println(one)
	if one >= zero {
		//fmt.Println("one is mcb :")
		//fmt.Println(oneSlice)
		slicesOgr = oneSlice

	} else {
		//	fmt.Println("zero is mcb :")
		//fmt.Println(zeroSlice)
		slicesOgr = zeroSlice
	}
}

func getCsr(i int) {
	zero := 0
	one := 0
	zeroSlice := make([][]int, 0)
	oneSlice := make([][]int, 0)
	for pos, val := range slicesCsr {
		switch val[i] {
		case 0:
			zero++
			zeroSlice = append(zeroSlice, slicesCsr[pos])
		case 1:
			one++
			oneSlice = append(oneSlice, slicesCsr[pos])
		}
	}
	//fmt.Println(zero)
	//fmt.Println(one)
	if zero <= one {
		//fmt.Println("zero is lcb :")
		//fmt.Println(zeroSlice)
		slicesCsr = zeroSlice

	} else {
		//fmt.Println("one is lcb :")
		//fmt.Println(oneSlice)
		slicesCsr = oneSlice
	}
}

func converter(sliced [][]int) (cb int64) {
	cbstr := ""
	for _, val := range sliced[0] {
		cbstr = cbstr + fmt.Sprint(val)
	}
	cb, err := strconv.ParseInt(cbstr, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	return cb

}

func main() {

	//f, err := os.Open("../test_input.txt")
	f, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	oneTime := false
	slicesOgr = make([][]int, 0)
	slicesCsr = make([][]int, 0)

	for scanner.Scan() {
		line := scanner.Text()
		if !oneTime {
			lineL = len(line)
		}
		updateSlicesAll(line)
	}

	for i := 0; i < lineL; i++ {
		if len(slicesOgr) > 1 {
			getOgr(i)
		}

		if len(slicesCsr) > 1 {
			getCsr(i)
		}
	}
	Ogr := converter(slicesOgr)
	Csr := converter(slicesCsr)
	fmt.Printf("CO2 scrubber rating: %d\n", Csr)
	fmt.Printf("Oxegen generator ragting: %d\n", Ogr)
	fmt.Printf("Answer is: %d\n", Ogr*Csr)

}
