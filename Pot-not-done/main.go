package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	file, err := os.Open("pot.3.in")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	//sscanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	var sum (int64) = 0
	for scanner.Scan() {
		numString := scanner.Text()
		powerof, _ := strconv.ParseInt(numString[len(numString)-1:len(numString)], 10, 64)
		num, _ := strconv.ParseInt(numString[:len(numString)-1], 10, 64)

		partSum := num
		var i (int64) = 1
		for i < powerof {
			if i == 1 {
				partSum = num * num
			} else {
				partSum = partSum * num
			}
			i++
		}
		sum += partSum
	}
	fmt.Println(sum)
}
