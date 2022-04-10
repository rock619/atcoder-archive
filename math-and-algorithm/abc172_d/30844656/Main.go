package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N int64) {
	fs := make([]int64, N)
	for n := int64(1); n <= N; n++ {
		for m := n; m <= N; m += n {
			fs[m-1]++
		}
	}

	sum := int64(0)
	for n := int64(1); n <= N; n++ {
		sum += n * fs[n-1]
	}
	fmt.Println(sum)
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
