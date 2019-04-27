package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	sky        [][]string
	numOfLines int
	numOfRows  int
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	caseCount := 1
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {

		numOfLines, _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		numOfRows, _ = strconv.Atoi(scanner.Text())

		sky = make([][]string, numOfLines)

		for i := range sky {
			scanner.Scan()
			row := scanner.Text()
			sky[i] = make([]string, numOfRows)
			for j, v := range row {
				sky[i][j] = string(v)
			}
		}
		starCount := 0
		for i := range sky {
			for j := 0; j < numOfRows; j++ {
				if string(sky[i][j]) == "-" {
					starCount++
					findEndStar(i, j)
				}

			}
		}
		fmt.Printf("case %v: %v\n", caseCount, starCount)
		caseCount++
	}
}

func findEndStar(i, j int) {
	sky[i][j] = "#"

	if i < numOfLines-1 && string(sky[i+1][j]) == "-" {
		findEndStar(i+1, j)
	}
	if i != 0 && string(sky[i-1][j]) == "-" {
		findEndStar(i-1, j)
	}
	if j < numOfRows-1 && string(sky[i][j+1]) == "-" {
		findEndStar(i, j+1)
	}
	if j != 0 && string(sky[i][j-1]) == "-" {
		findEndStar(i, j-1)
	}
}
