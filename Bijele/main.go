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

	answer := ""
	correctSet := []int{1, 1, 2, 2, 2, 8}
	i := 0
	for scanner.Scan() {
		current, _ := strconv.Atoi(scanner.Text())
		answer = answer + " " + strconv.Itoa(correctSet[i]-current)
		i++
	}
	fmt.Println(answer)
}
