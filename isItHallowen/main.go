package main

import (
	"bufio"
	"fmt"
	"os"
)

type cup struct {
	color  string
	radius float64
}

func main() {
	/*
		file, err := os.Open("1.in")
		defer file.Close()
		if err != nil {
			log.Fatal(err)
		}
		scanner := bufio.NewScanner(file)
	*/
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	date := scanner.Text()

	if date == "OCT 31" || date == "DEC 25" {
		fmt.Println("yup")
	} else {
		fmt.Println("nope")
	}
}
