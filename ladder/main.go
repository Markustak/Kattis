package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {

	file, err := os.Open("ladder.00.in")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	//	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	height, _ := strconv.ParseFloat(scanner.Text(), 64)
	scanner.Scan()
	angle, _ := strconv.ParseFloat(scanner.Text(), 64)

	b := height * math.Cos(angle)
	fmt.Println("h", height, " a", angle, " b", b)
	fmt.Println("R", math.Sqrt(b*b+height*height))
	fmt.Println("Test", math.Sqrt(5*5+7*7))
}
