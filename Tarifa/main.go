package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	/*
		file, err := os.Open("tarifa.1.in")
		defer file.Close()
		if err != nil {
			log.Fatal(err)
		}
		scanner := bufio.NewScanner(file)
	*/
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	perMonth, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	numOfMonths, _ := strconv.Atoi(scanner.Text())
	remaining := perMonth
	for i := 0; i < numOfMonths; i++ {
		scanner.Scan()
		month, _ := strconv.Atoi(scanner.Text())
		remaining = remaining + perMonth - month
	}
	fmt.Println(remaining)
}
