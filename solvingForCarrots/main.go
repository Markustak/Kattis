package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	/*file, err := os.Open("carrots.02.in")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)*/

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	count := 1
	secondValue := ""
	for scanner.Scan() {
		readline := scanner.Text()

		if count == 2 {
			secondValue = readline
		}
		count++
	}
	fmt.Println(secondValue)
}
