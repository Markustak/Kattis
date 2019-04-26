package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	/*
		file, err := os.Open("sample02.in")
		defer file.Close()
		if err != nil {
			log.Fatal(err)
		}
		scanner := bufio.NewScanner(file)
	*/
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	gold, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	silver, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	copper, _ := strconv.Atoi(scanner.Text())
	buyingPower := gold*3 + silver*2 + copper
	buy := ""
	if buyingPower >= 8 {
		buy = "Province or "
	} else if buyingPower >= 5 {
		buy = "Duchy or "
	} else if buyingPower >= 2 {
		buy = "Estate or "
	}

	if buyingPower >= 6 {
		buy += "Gold"
	} else if buyingPower >= 3 {
		buy += "Silver"
	} else {
		buy += "Copper"
	}
	fmt.Println(buy)
}
