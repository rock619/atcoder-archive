package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solve(N int64) {
	divisors := make([]int64, 0)
	for i := int64(1); i*i <= N; i++ {
		if N%i == 0 {
			divisors = append(divisors, i, N/i)
		}
	}
	print(divisors)
}

func print(ns []int64) {
	parts := make([]string, len(ns))
	for i, v := range ns {
		parts[i] = strconv.FormatInt(v, 10)
	}
	fmt.Println(strings.Join(parts, "\n"))
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
