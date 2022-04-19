package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(s string) {
	count := 0
	for _, r := range s {
		if r == '1' {
			count++
		}
	}
	fmt.Println(count)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	s := scanner.Text()
	solve(s)
}
