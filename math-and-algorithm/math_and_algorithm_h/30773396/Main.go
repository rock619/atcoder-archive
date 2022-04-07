package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N int64, S int64) {
	count := 0
	for i := int64(1); i <= N; i++ {
		for j := int64(1); j <= N; j++ {
			if i+j <= S {
				count++
			}
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
	var N int64
	scanner.Scan()
	N, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	var S int64
	scanner.Scan()
	S, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	solve(N, S)
}
