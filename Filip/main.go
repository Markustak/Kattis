package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	/*
		file, err := os.Open("filip.dummy.2.in")
		defer file.Close()
		if err != nil {
			log.Fatal(err)
		}
		scanner := bufio.NewScanner(file)
	*/
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	firstString := scanner.Text()
	firstNumber := reverseString(firstString)
	scanner.Scan()
	secondString := scanner.Text()
	secondNumber := reverseString(secondString)
	if secondNumber > firstNumber {
		fmt.Println(secondNumber)
	} else {
		fmt.Println(firstNumber)
	}
}

func reverseString(input string) (output string) {
	for k, _ := range input {
		output = output + input[len(input)-k-1:len(input)-k]
	}
	return output
}
