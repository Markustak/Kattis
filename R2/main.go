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
	data := make([]int, 2)
	i := 0
	for scanner.Scan() {
		readline := scanner.Text()
		data[i], _ = strconv.Atoi(readline)
		i++
	}
	fmt.Println(data[1]*2 - data[0])
}
