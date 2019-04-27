package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type scenario struct {
	count      int
	currentPos byte
	leavePos   byte
}

func main() {

	scanner := &bufio.Scanner{}

	local := len(os.Args) > 1
	if local {
		file, err := os.Open("toilet-1.in")
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
	input := scanner.Text()

	alwaysUp := scenario{
		count:      0,
		currentPos: input[0],
		leavePos:   85, //U
	}

	alwaysDown := scenario{
		count:      0,
		currentPos: input[0],
		leavePos:   68, //D
	}
	ownPreference := scenario{
		count:      0,
		currentPos: input[0],
		leavePos:   input[0],
	}
	fmt.Println(checkScenario(alwaysUp, input, false))
	fmt.Println(checkScenario(alwaysDown, input, false))
	fmt.Println(checkScenario(ownPreference, input, true))
}

func checkScenario(scenario scenario, input string, changingLeavepos bool) int {

	for i := 1; i < len(input); i++ {
		if changingLeavepos {
			scenario.leavePos = input[i]
		}
		if scenario.currentPos != input[i] {
			scenario.count++
			scenario.currentPos = input[i]
		}
		if scenario.currentPos != scenario.leavePos {
			scenario.count++
			scenario.currentPos = scenario.leavePos
		}
	}
	return scenario.count
}
