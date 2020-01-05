package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	counts := make(map[string]int)
	line := bufio.NewScanner(os.Stdin)

	for line.Scan() {
		counts[line.Text()]++
	}

	for data, counter := range counts {
		if counter > 1 {
			fmt.Println(data, counter)
		}
	}
}
