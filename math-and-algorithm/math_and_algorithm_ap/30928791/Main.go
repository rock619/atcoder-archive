package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const MOD = 1000000007

func solve(N int64) {
	ints := make([]int64, N)
	ints[0] = 1
	ints[1] = 1
	for i := int64(2); i < N; i++ {
		ints[i] = (ints[i-1] + ints[i-2]) % MOD
	}

	fmt.Println(ints[N-1])
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
