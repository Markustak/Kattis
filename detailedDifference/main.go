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
		file, err := os.Open("sample.in")
		defer file.Close()
		if err != nil {
			log.Fatal(err)
		}
		scanner = bufio.NewScanner(file)
	} else {
		scanner = bufio.NewScanner(os.Stdin)
	}
	scanner.Scan()

	for scanner.Scan() {
		stringOne := scanner.Text()
		scanner.Scan()
		stringTwo := scanner.Text()

		answerString := ""
		for i, v := range stringOne {

			if rune(stringTwo[i]) == v {
				answerString += "."
			} else {
				answerString += "*"
			}
		}
		fmt.Printf("%v\n%v\n%v\n", stringOne, stringTwo, answerString)
	}
}
