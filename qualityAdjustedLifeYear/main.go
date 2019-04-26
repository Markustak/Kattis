package main

import (
	"bufio"
	"fmt"
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
	totalQuality := 0.0
	for scanner.Scan() {
		readline := scanner.Text()
		numOfperiods, _ := strconv.Atoi(readline)

		for i := 0; i < numOfperiods; i++ {
			scanner.Scan()
			quality, _ := strconv.ParseFloat(scanner.Text(), 64)
			scanner.Scan()
			years, _ := strconv.ParseFloat(scanner.Text(), 64)
			totalQuality = totalQuality + quality*years

		}
	}
	fmt.Println(totalQuality)
}
