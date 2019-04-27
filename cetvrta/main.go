package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type cup struct {
	color  string
	radius float64
}

func main() {

	scanner := &bufio.Scanner{}

	local := len(os.Args) > 1
	if local {
		file, err := os.Open("2.in")
		defer file.Close()
		if err != nil {
			log.Fatal(err)
		}
		scanner = bufio.NewScanner(file)
	} else {
		scanner = bufio.NewScanner(os.Stdin)
	}
	scanner.Split(bufio.ScanWords)
	xCoords := make(map[int]int)
	yCoords := make(map[int]int)

	for i := 0; i < 3; i++ {
		scanner.Scan()
		x, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		y, _ := strconv.Atoi(scanner.Text())
		xCoords = checkAndAd(xCoords, x)
		yCoords = checkAndAd(yCoords, y)
	}
	missingX := getMissing(xCoords)
	missingY := getMissing(yCoords)
	fmt.Printf("%v %v \n", missingX, missingY)
}

func checkAndAd(coords map[int]int, coor int) map[int]int {
	if val, ok := coords[coor]; ok {
		coords[coor] = val + 1
	} else {
		coords[coor] = 1
	}
	return coords
}

func getMissing(coords map[int]int) int {
	for i, v := range coords {
		if v == 1 {
			return i
		}
	}
	return 0
}
