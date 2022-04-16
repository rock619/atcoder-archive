package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const MOD = 1000000007

func solve(N int64, A []int64) {
	c := int64(1)
	sum := int64(0)
	for _, a := range A {
		sum = (sum + a*c) % MOD
		c = 2 * c % MOD
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
	A := make([]int64, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		A[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	solve(N, A)
}
