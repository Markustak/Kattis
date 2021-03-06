package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	cmd := os.Args[0]

	testEnv := len(cmd) > 1
	scanner := &bufio.Scanner{}
	if testEnv {
		file, err := os.Open("1.in")
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
	numOfCases, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < numOfCases; i++ {
		scanner.Scan()
		noAdvertising, _ := strconv.Atoi(scanner.Text())

		scanner.Scan()
		advertising, _ := strconv.Atoi(scanner.Text())

		scanner.Scan()
		cost, _ := strconv.Atoi(scanner.Text())

		if noAdvertising > advertising-cost {
			fmt.Println("do not advertise")
		} else if noAdvertising < advertising-cost {
			fmt.Println("advertise")
		} else {
			fmt.Println("does not matter")
		}
	}
}
