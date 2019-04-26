package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	/*file, err := os.Open("a.in")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)*/

	scanner := bufio.NewScanner(os.Stdin)
	number := 0
	for scanner.Scan() {
		number, _ = strconv.Atoi(scanner.Text())

	}

	for i := 0; i < number; i++ {
		fmt.Println(i+1, "Abracadabra")
	}
}
