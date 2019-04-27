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
		file, err := os.Open("tri-02.in")
		defer file.Close()
		if err != nil {
			log.Fatal(err)
		}
		scanner = bufio.NewScanner(file)
	} else {
		scanner = bufio.NewScanner(os.Stdin)
	}
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	first, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	second, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	third, _ := strconv.Atoi(scanner.Text())

	if first+second == third {
		fmt.Printf("%v+%v=%v\n", first, second, third)
	} else if first-second == third {
		fmt.Printf("%v-%v=%v\n", first, second, third)
	} else if first*second == third {
		fmt.Printf("%v*%v=%v\n", first, second, third)
	} else if first/second == third {
		fmt.Printf("%v/%v=%v\n", first, second, third)
	} else if first == second+third {
		fmt.Printf("%v=%v+%v\n", first, second, third)
	} else if first == second-third {
		fmt.Printf("%v=%v-%v\n", first, second, third)
	} else if first == second*third {
		fmt.Printf("%v=%v*%v\n", first, second, third)
	} else if first == second/third {
		fmt.Printf("%v=%v/%v\n", first, second, third)
	}
}
