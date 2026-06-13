package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func main() {
	fmt.Printf("m: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	
	energy, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("Error")
	}

	energy = energy * 300000000 * 300000000
	fmt.Printf("E: %v", energy)
}
