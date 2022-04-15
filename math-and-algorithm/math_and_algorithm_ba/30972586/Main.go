package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N int64) {
	for n := int64(1); n <= N; n = 2*n + 1 {
		if n == N {
			fmt.Println("Second")
			return
		}
	}
	fmt.Println("First")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var N int64
	scanner.Scan()
	N, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	solve(N)
}
