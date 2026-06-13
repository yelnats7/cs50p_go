package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func main() {
	fmt.Printf("How much was the meal? ")
	dollarScanner := bufio.NewScanner(os.Stdin)
	dollarScanner.Scan()
	dollars := dolarsToFloat(dollarScanner.Text())

	fmt.Printf("What percentage would you like to tip? ")
	percentScanner := bufio.NewScanner(os.Stdin)
	percentScanner.Scan()
	percent := percentToFloat(percentScanner.Text())

	tip := dollars * percent
	fmt.Printf("Leave $%.2f", tip)
}

func dolarsToFloat(d string) float64 {
	value, err := strconv.ParseFloat(strings.ReplaceAll(d, "$", ""), 64)
	if err != nil {
		fmt.Println("Error")
	}
	return value
}

func percentToFloat(p string) float64 {
	value, err := strconv.ParseFloat(strings.ReplaceAll(p, "%", ""), 64)
	if err != nil {
		fmt.Println("Error")
	}
	return value / float64(100)
}
