package main

import (
	"fmt"
	"bufio"
	"strings"
	"os"
)

func convert(text string) {
	replacer := strings.NewReplacer(":)", "🙂", ":(", "🙁")
	fmt.Println(replacer.Replace(text))
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	convert(input)

}
