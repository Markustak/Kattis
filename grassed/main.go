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
	scanner.Scan()

	cost, _ := strconv.ParseFloat(scanner.Text(), 64)
	scanner.Scan()
	numOfLawns, _ := strconv.Atoi(scanner.Text())

	if local {
		fmt.Printf("cost %v, numofLawns %v, \n", cost, numOfLawns)
	}
	totalCost := 0.0
	for i := 0; i < numOfLawns; i++ {
		scanner.Scan()
		width, _ := strconv.ParseFloat(scanner.Text(), 64)
		scanner.Scan()
		length, _ := strconv.ParseFloat(scanner.Text(), 64)
		if local {
			fmt.Printf("width %v, length %v \n", width, length)
		}
		totalCost += width * length * cost
	}

	fmt.Println(totalCost)
}
