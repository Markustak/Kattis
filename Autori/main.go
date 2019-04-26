package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	/*
		file, err := os.Open("autori.1.in")
		defer file.Close()
		if err != nil {
			log.Fatal(err)
		}
		scanner := bufio.NewScanner(file)
	*/
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	splitLine := strings.Split(line, "-")
	answer := ""
	for _, name := range splitLine {
		answer = answer + name[:1]
	}
	fmt.Println(answer)
}
