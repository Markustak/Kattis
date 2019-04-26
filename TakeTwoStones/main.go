package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	/*	file, err := os.Open("3.in")
		defer file.Close()
		if err != nil {
			log.Fatal(err)
		}*/
	scanner := bufio.NewScanner(os.Stdin)
	numOfStones := 0
	for scanner.Scan() {
		numOfStones, _ = strconv.Atoi(scanner.Text())

	}
	currentStone := 1
	for currentStone <= numOfStones-1 {
		currentStone = currentStone + 2
	}

	stonesleft := numOfStones - currentStone - 1
	if stonesleft%2 == 0 {
		fmt.Println("Bob")
	} else {
		fmt.Println("Alice")
	}
}
