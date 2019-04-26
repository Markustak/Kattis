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
	for scanner.Scan() {
		distanceTraveled := 0
		numOfEntries, _ := strconv.Atoi(scanner.Text())
		if numOfEntries < 0 {
			break
		}
		times := make([]int, numOfEntries)
		for i := 0; i < numOfEntries; i++ {
			scanner.Scan()
			speed, _ := strconv.Atoi(scanner.Text())
			scanner.Scan()
			times[i], _ = strconv.Atoi(scanner.Text())
			if i > 0 {
				distanceTraveled = distanceTraveled + (speed * (times[i] - times[i-1]))
			} else {
				distanceTraveled = distanceTraveled + speed*times[i]
			}
		}
		fmt.Println(distanceTraveled, "miles")
	}
}
