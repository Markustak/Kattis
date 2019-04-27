package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type cup struct {
	color  string
	radius float64
}

func main() {
	/*
		file, err := os.Open("cups-01.in")
		defer file.Close()
		if err != nil {
			log.Fatal(err)
		}
		scanner := bufio.NewScanner(file)
	*/
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	numOfCups, _ := strconv.Atoi(scanner.Text())
	cups := make([]cup, numOfCups)
	i := 0
	for scanner.Scan() {
		v := scanner.Text()
		iInt, err := strconv.ParseFloat(v, 64)

		if err != nil {
			cups[i/2].color = v
		} else {
			if cups[i/2].color != "" {
				cups[i/2].radius = iInt
			} else {
				cups[i/2].radius = iInt / 2
			}
		}
		i++
	}

	sort.Slice(cups, func(i, j int) bool {
		return cups[i].radius < cups[j].radius
	})
	for _, v := range cups {
		fmt.Println(v.color)
	}
}
