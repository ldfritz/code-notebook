package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	uniqRows := make(map[string]bool)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		row := scanner.Text()
		if uniqRows[row] {
			continue
		}
		uniqRows[row] = true
		fmt.Println(row)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "bufio.Scanner.Err: %v\n", err)
		os.Exit(1)
	}
}
