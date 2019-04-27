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
		file, err := os.Open("sample.in")
		defer file.Close()
		if err != nil {
			log.Fatal(err)
		}
		scanner = bufio.NewScanner(file)
	} else {
		scanner = bufio.NewScanner(os.Stdin)
	}
	scanner.Split(bufio.ScanWords)
	setNum := 0
	for scanner.Scan() {
		setNum++
		numOfNames, _ := strconv.Atoi(scanner.Text())
		names := make([]string, numOfNames)

		flip := false
		j := 0
		k := 1
		for i := 0; i < numOfNames; i++ {
			scanner.Scan()
			name := scanner.Text()
			if !flip {
				names[j] = name
				j++
			} else {
				names[numOfNames-k] = name
				k++
			}
			flip = !flip
		}
		if len(names) > 0 {
			fmt.Println("SET", setNum)
			for _, v := range names {
				fmt.Println(v)
			}
		}
	}
}
