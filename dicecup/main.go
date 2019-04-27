package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	scanner := &bufio.Scanner{}

	local := len(os.Args) > 1
	if local {
		file, err := os.Open("sample03.in")
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
	diceOne, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	diceTwo, _ := strconv.Atoi(scanner.Text())

	if diceOne == diceTwo {
		fmt.Println(diceOne + 1)
	} else {
		start := diceTwo + 1
		end := diceTwo + (diceOne - diceTwo) + 1
		if start > end {
			start, end = end, start
		}
		for start <= end {
			fmt.Println(start)
			start++
		}
	}
}
