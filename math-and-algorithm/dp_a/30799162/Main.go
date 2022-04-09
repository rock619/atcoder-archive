package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N int64, h []int64) {
	result := make([]int64, N)
	for i := int64(0); i < N; i++ {
		switch i {
		case 0:
			result[i] = 0
		case 1:
			result[i] = Abs(h[i] - h[i-1])
		default:
			result[i] = Min(result[i-2]+Abs(h[i]-h[i-2]), result[i-1]+Abs(h[i]-h[i-1]))
		}
	}

	fmt.Println(result[len(result)-1])
}

func Abs(n int64) int64 {
	if n < 0 {
		return -n
	}
	return n
}

func Min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
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
	h := make([]int64, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		h[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	solve(N, h)
}
