package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	/*
		file, err := os.Open("sample.in")
		defer file.Close()
		if err != nil {
			log.Fatal(err)
		}
		scanner := bufio.NewScanner(file)
	*/
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	//We don't need the first number
	scanner.Scan()
	for scanner.Scan() {
		number, _ := strconv.Atoi(scanner.Text())
		if number%2 == 0 {
			fmt.Println(number, "is even")
		} else {
			fmt.Println(number, "is odd")
		}
	}

}
