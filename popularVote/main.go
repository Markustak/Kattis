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
		file, err := os.Open("vote-01.in")
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
		numOfCandidates, _ := strconv.Atoi(scanner.Text())
		winner := 0
		mostVotes := 0
		totalVotes := 0
		draw := false
		for j := 0; j < numOfCandidates; j++ {
			scanner.Scan()
			votes, _ := strconv.Atoi(scanner.Text())
			if votes > mostVotes {
				winner = j + 1
				mostVotes = votes
				draw = false
			} else if votes == mostVotes {
				draw = true
			}
			totalVotes += votes
		}
		if draw {
			fmt.Println("no winner")
		} else if totalVotes/2 < mostVotes {
			fmt.Println("majority winner", winner)
		} else {
			fmt.Println("minority winner", winner)
		}
	}
}
