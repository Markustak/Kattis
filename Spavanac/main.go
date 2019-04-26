package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	/*
		file, err := os.Open("spavanac.1.in")
		defer file.Close()
		if err != nil {
			log.Fatal(err)
		}
		scanner := bufio.NewScanner(file)
	*/
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	hours, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	minutes, _ := strconv.Atoi(scanner.Text())
	combined := hours*60 + minutes - 45
	if combined < 0 {
		combined = 24*60 + combined
	}
	h := combined / 60
	m := combined % 60

	fmt.Println(h, m)
}
