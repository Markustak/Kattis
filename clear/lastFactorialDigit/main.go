package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	scanner.Split(bufio.ScanWords)
	scanner.Scan()

	_ = scanner.Text()
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		count := 0
		for i := 1; i <= n; i++ {

			if count == 0 {
				count += i
			} else {
				count = count * i
			}
		}
		fmt.Println(count % 10)
	}
}
