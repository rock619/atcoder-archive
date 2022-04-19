package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(a int64, b int64) {
	if a*b%2 == 1 {
		fmt.Println("Odd")
	} else {
		fmt.Println("Even")
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	a, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	scanner.Scan()
	b, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	solve(a, b)
}
