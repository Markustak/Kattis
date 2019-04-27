package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type cup struct {
	color  string
	radius float64
}

func main() {

	scanner := &bufio.Scanner{}

	local := len(os.Args) > 1
	if local {
		file, err := os.Open("1.in")
		defer file.Close()
		if err != nil {
			log.Fatal(err)
		}
		scanner = bufio.NewScanner(file)
	} else {
		scanner = bufio.NewScanner(os.Stdin)
	}
	scanner.Scan()

	text := scanner.Text()
	per := "PER"
	i := 0
	dayCount := 0
	for _, v := range text {

		if v != rune(per[i]) {
			dayCount++
		}
		if i != 2 {
			i++
		} else {
			i = 0
		}
	}

	fmt.Println(dayCount)
}
