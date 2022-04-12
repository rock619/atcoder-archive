package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const MOD = 1000000007

func ModPow(a, b, mod int64) int64 {
	p := a
	result := int64(1)
	for i := 0; i < 30; i++ {
		if b&(1<<i) != 0 {
			result *= p
			result %= mod
		}
		p *= p
		p %= mod
	}

	return result
}

func ModDiv(a, b, mod int64) int64 {
	return a * ModPow(b, mod-2, mod) % mod
}

func solve(X int64, Y int64) {
	numerator, denominator := int64(1), int64(1)
	for i := int64(1); i <= X+Y; i++ {
		numerator *= i
		numerator %= MOD
	}
	for i := int64(1); i <= X; i++ {
		denominator *= i
		denominator %= MOD
	}
	for i := int64(1); i <= Y; i++ {
		denominator *= i
		denominator %= MOD
	}

	fmt.Println(ModDiv(numerator, denominator, MOD))
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var X int64
	scanner.Scan()
	X, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	var Y int64
	scanner.Scan()
	Y, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	solve(X, Y)
}
