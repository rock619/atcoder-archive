package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const MOD = 1000000007

func ModDiv(a, b, mod int64) int64 {
	return a * ModPow(b, mod-2, mod) % mod
}

func ModPow(a, b, mod int64) int64 {
	p := a
	result := int64(1)
	for i := 0; i < 62; i++ {
		if b&(1<<i) != 0 {
			result *= p
			result %= mod
		}
		p *= p
		p %= mod
	}

	return result
}

func solve(N int64) {
	fmt.Println(ModDiv(ModPow(4, N+1, MOD)-1, 3, MOD))
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
