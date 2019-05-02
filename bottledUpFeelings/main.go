package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type item struct {
	value  int
	weight int
}

func main() {

	scanner := &bufio.Scanner{}

	local := len(os.Args) > 1
	if local {
		file, err := os.Open("001.in")
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

	oil, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	bottleOne, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	bottleTwo, _ := strconv.Atoi(scanner.Text())

	doWork(oil, bottleOne, bottleTwo)
}

func doWork(oil, bottleOne, bottleTwo int) {
	if oil < bottleOne && oil < bottleTwo {
		fmt.Println("Impossible")
		return
	}
	remain := oil % bottleOne
	if remain == 0 {
		fmt.Println(oil/bottleOne, " ", 0)
		return
	}

	bottleOneCount := oil / bottleOne
	for bottleOneCount != 0 {
		if (oil-(bottleOne*bottleOneCount))%bottleTwo == 0 {
			bottleTwoCount := (oil - (bottleOne * bottleOneCount)) / bottleTwo
			fmt.Printf("%v %v\n", bottleOneCount, bottleTwoCount)
			return
		}

		bottleOneCount--
	}
	if oil%bottleTwo == 0 {
		fmt.Printf("%v %v", 0, oil/bottleTwo)
		return
	}
	fmt.Println("Impossible")
}
