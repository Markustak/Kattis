package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type cup struct {
	color  string
	radius float64
}

func main() {
	/*
		file, err := os.Open("everywhere-01.in")
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
		numOfTrips, _ := strconv.Atoi(scanner.Text())
		cities := make(map[string]bool)
		unique := 0
		for j := 0; j < numOfTrips; j++ {
			scanner.Scan()
			city := scanner.Text()
			if !cities[city] {
				cities[city] = true
				unique++
			}
		}
		fmt.Println(unique)
	}
}
