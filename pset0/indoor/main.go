package main

import (
	"fmt"
	"bufio"
	"strings"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := strings.ToLower(scanner.Text())
	fmt.Println(input)
}
