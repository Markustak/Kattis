package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	scanner := &bufio.Scanner{}

	local := len(os.Args) > 1
	if local {
		file, err := os.Open("3.in")
		defer file.Close()
		if err != nil {
			log.Fatal(err)
		}
		scanner = bufio.NewScanner(file)
	} else {
		scanner = bufio.NewScanner(os.Stdin)
	}
	scanner.Split(bufio.ScanWords)
	unique := make(map[int]bool)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		if !unique[num%42] {
			unique[num%42] = true
		}
	}
	fmt.Println(len(unique))
}
