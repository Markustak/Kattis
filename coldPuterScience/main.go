package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	/*
		file, err := os.Open("cold-002.in")
		defer file.Close()
		if err != nil {
			log.Fatal(err)
		}
		scanner := bufio.NewScanner(file)
	*/
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	numOfDays, _ := strconv.Atoi(scanner.Text())
	daysBelow := 0

	for i := 0; i < numOfDays; i++ {
		scanner.Scan()
		day, _ := strconv.Atoi(scanner.Text())
		if day < 0 {
			daysBelow++
		}
	}
	fmt.Println(daysBelow)
}
