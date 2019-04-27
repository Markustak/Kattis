package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cup struct {
	color  string
	radius float64
}

func main() {
	/*
		file, err := os.Open("hissingmicrophone-03.in")
		defer file.Close()
		if err != nil {
			log.Fatal(err)
		}
		scanner := bufio.NewScanner(file)
	*/
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	in := scanner.Text()

	if strings.Contains(in, "ss") {
		fmt.Println("hiss")
	} else {
		fmt.Println("no hiss")
	}
}
