package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	/*
		file, err := os.Open("faktor.dummy.3.in")
		defer file.Close()
		if err != nil {
			log.Fatal(err)
		}
		scanner := bufio.NewScanner(file)
	*/
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	numOfArticles, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	expImpactFactor, _ := strconv.Atoi(scanner.Text())
	numToBribe := numOfArticles*expImpactFactor - (numOfArticles - 1)
	fmt.Println(numToBribe)

}
