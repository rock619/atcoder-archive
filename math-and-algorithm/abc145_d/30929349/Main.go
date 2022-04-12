package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const MOD = 1000000007

func ModCombination(a, b, mod int64) int64 {
	numerator, denominator := int64(1), int64(1)
	for i := int64(1); i <= a+b; i++ {
		numerator *= i
		numerator %= mod
	}
	for i := int64(1); i <= a; i++ {
		denominator *= i
		denominator %= mod
	}
	for i := int64(1); i <= b; i++ {
		denominator *= i
		denominator %= mod
	}

	return ModDiv(numerator, denominator, mod)
}

func ModDiv(a, b, mod int64) int64 {
	return a * ModPow(b, mod-2, mod) % mod
}

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

func solve(X int64, Y int64) {
	if 2*Y < X || (2*Y-X)%3 != 0 || 2*X < Y || (2*X-Y)%3 != 0 {
		fmt.Println(0)
		return
	}

	move1Count, move2Count := (2*Y-X)/3, (2*X-Y)/3
	fmt.Println(ModCombination(move1Count, move2Count, MOD))
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
