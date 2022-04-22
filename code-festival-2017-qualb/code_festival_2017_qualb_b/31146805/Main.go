package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	YES = "YES"
	NO  = "NO"
)

func solve(N int64, D []int64, M int64, T []int64) {
	m := make(map[int64]int64, N)
	for _, d := range D {
		m[d]++
	}
	for _, t := range T {
		if m[t] == 0 {
			fmt.Println(NO)
			return
		}
		m[t]--
	}
	fmt.Println(YES)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	N, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	D := make([]int64, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		D[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	scanner.Scan()
	M, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	T := make([]int64, M)
	for i := int64(0); i < M; i++ {
		scanner.Scan()
		T[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	solve(N, D, M, T)
}
