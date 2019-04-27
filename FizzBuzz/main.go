package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	/*
		file, err := os.Open("fizzbuzz-02.in")
		defer file.Close()
		if err != nil {
			log.Fatal(err)
		}
		scanner := bufio.NewScanner(file)
	*/
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	x, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	y, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	for i := 1; i <= n; i++ {
		if i%x == 0 {
			fmt.Print("Fizz")
		}
		if i%y == 0 {
			fmt.Print("Buzz")
		}
		if i%x != 0 && i%y != 0 {
			fmt.Print(i)
		}
		fmt.Print("\n")
	}
}
