package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N int64) {
	results := make([]int64, N+1)
	for i := int64(0); i <= N; i++ {
		switch i {
		case 0, 1:
			results[i] = 1
		default:
			results[i] = results[i-1] + results[i-2]
		}
	}
	fmt.Println(results[N])
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
