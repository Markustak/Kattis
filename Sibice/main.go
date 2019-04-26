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
		file, err := os.Open("1.in")
		defer file.Close()
		if err != nil {
			log.Fatal(err)
		}
		scanner := bufio.NewScanner(file)
	*/
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	scanner.Scan()
	boxWidth, _ := strconv.ParseFloat(scanner.Text(), 64)
	scanner.Scan()
	boxHeight, _ := strconv.ParseFloat(scanner.Text(), 64)
	diagnoal := math.Sqrt((boxHeight * boxHeight) + (boxWidth * boxWidth))
	for scanner.Scan() {
		matchSize, _ := strconv.ParseFloat(scanner.Text(), 64)
		if matchSize <= diagnoal {
			fmt.Println("DA")
		} else {
			fmt.Println("NE")
		}
	}

}
