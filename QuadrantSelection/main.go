package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	x, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	y, _ := strconv.Atoi(scanner.Text())
	if x >= 1 && y >= 1 {
		fmt.Println("1")
	} else if x <= -1 && y >= 1 {
		fmt.Println("2")
	} else if x <= -1 && y <= -1 {
		fmt.Println("3")
	} else if x >= 1 && y <= -1 {
		fmt.Println("4")
	}
}
