package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N int64, K int64) {
	count := int64(0)
	for a := int64(1); a <= N; a++ {
		for b := Max(1, a-K+1); b <= Min(N, a+K-1); b++ {
			for c := Max(1, a-K+1); c <= Min(N, a+K-1); c++ {
				if Abs(b-c) < K {
					count++
				}
			}
		}
	}
	fmt.Println(N*N*N - count)
}

func Max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int64) int64 {
	if a > b {
		return b
	}
	return a
}

func Abs(n int64) int64 {
	if n < 0 {
		return -n
	}
	return n
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
	var K int64
	scanner.Scan()
	K, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	solve(N, K)
}
