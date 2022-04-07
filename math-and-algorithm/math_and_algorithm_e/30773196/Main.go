package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const MOD = 100

func solve(N int64, a []int64) {
	sum := int64(0)
	for _, v := range a {
		sum += v
	}
	fmt.Println(sum % MOD)
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
	a := make([]int64, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		a[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	solve(N, a)
}
