package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type item struct {
	value  int
	weight int
}

func main() {

	scanner := &bufio.Scanner{}

	local := len(os.Args) > 1
	if local {
		file, err := os.Open("l-sample.in")
		defer file.Close()
		if err != nil {
			log.Fatal(err)
		}
		scanner = bufio.NewScanner(file)
	} else {
		scanner = bufio.NewScanner(os.Stdin)
	}
	scanner.Split(bufio.ScanWords)

	scanner.Scan() //Throwing num of testcases since we dont need it

	for scanner.Scan() {
		stops, _ := strconv.Atoi(scanner.Text())
		checkPassengers(stops)
	}
}

func checkPassengers(stops int) {
	if stops == 1 {
		fmt.Println(1)
		return
	}
	pass := 0.0
	for i := 0; i < stops; i++ {
		if i == 0 {
			pass = 1.0
		} else {
			pass = (pass + 0.5) * 2
		}
	}

	fmt.Println(int(pass))
}
