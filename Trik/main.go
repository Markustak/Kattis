package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	/*
		file, err := os.Open("2.in")
		defer file.Close()
		if err != nil {
			log.Fatal(err)
		}
		scanner := bufio.NewScanner(file)
	*/
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	position := 1
	moves := scanner.Text()
	for _, move := range moves {
		switch move {
		case 65:
			if position == 1 {
				position = 2
			} else if position == 2 {
				position = 1
			}
		case 66:
			if position == 2 {
				position = 3
			} else if position == 3 {
				position = 2
			}
		case 67:
			if position == 1 {
				position = 3
			} else if position == 3 {
				position = 1
			}
		}
	}
	fmt.Println(position)
}
