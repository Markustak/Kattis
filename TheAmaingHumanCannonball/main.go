package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	/*
		file, err := os.Open("test0.in")
		defer file.Close()
		if err != nil {
			log.Fatal(err)
		}
		scanner := bufio.NewScanner(file)
	*/
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	numOfCases, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < numOfCases; i++ {
		scanner.Scan()
		velocity, _ := strconv.ParseFloat(scanner.Text(), 64)

		scanner.Scan()
		angle, _ := strconv.ParseFloat(scanner.Text(), 64)

		scanner.Scan()
		distance, _ := strconv.ParseFloat(scanner.Text(), 64)

		scanner.Scan()
		lowerEdge, _ := strconv.ParseFloat(scanner.Text(), 64)

		scanner.Scan()
		upperEdge, _ := strconv.ParseFloat(scanner.Text(), 64)
		fmt.Println(safeOrNot(velocity, angle, distance, lowerEdge, upperEdge))
	}
}

func safeOrNot(v, a, d, l, u float64) string {
	t := d / (v * math.Cos(a*math.Pi/180.0))
	pos := v*t*math.Sin(a*math.Pi/180.0) - 0.5*9.81*t*t
	if l+1 < pos && u-1 > pos {
		return "Safe"
	}
	return "Not Safe"
}
