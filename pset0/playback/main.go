package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := strings.ReplaceAll(scanner.Text(), " ", "...")
	fmt.Println(input)
}
