package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solve(N int64) {
	factors := PF(N)
	print(factors)
}

func print(ns []int64) {
	parts := make([]string, len(ns))
	for i, v := range ns {
		parts[i] = strconv.FormatInt(v, 10)
	}
	fmt.Println(strings.Join(parts, " "))
}

func PF(n int64) []int64 {
	factors := make([]int64, 0)
	for i := int64(2); i*i <= n; i++ {
		if n%i == 0 {
			factors = append(factors, i)
			n /= i
			i = 1
		}
	}

	return append(factors, n)
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
