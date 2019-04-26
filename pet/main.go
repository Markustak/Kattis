package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	/*
		file, err := os.Open("pet.2.in")
		defer file.Close()
		if err != nil {
			log.Fatal(err)
		}
		scanner := bufio.NewScanner(file)
	*/
	scanner := bufio.NewScanner(os.Stdin)
	highest := 0
	winner := -1
	player := 1
	for scanner.Scan() {
		sum := 0
		s := scanner.Text()
		scores := strings.Split(s, " ")
		for _, v := range scores {
			vI, _ := strconv.Atoi(v)
			sum += vI
		}
		if highest < sum {
			highest = sum
			winner = player
		}
		player++
	}
	fmt.Println(winner, highest)
}
