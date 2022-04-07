package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N uint64, A []uint64) {
	fmt.Println(LCMs(A...))
}

func LCMs(ns ...uint64) uint64 {
	result := uint64(1)
	for _, n := range ns {
		result = LCM(result, n)
	}
	return result
}

func LCM(a, b uint64) uint64 {
	gcd := GCD(a, b)
	return a / gcd * b
}

func GCD(a, b uint64) uint64 {
	for a > 0 && b > 0 {
		if a < b {
			b = b % a
		} else {
			a = a % b
		}
	}

	if a > 0 {
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
	var N uint64
	scanner.Scan()
	N, _ = strconv.ParseUint(scanner.Text(), 10, 64)
	A := make([]uint64, N)
	for i := uint64(0); i < N; i++ {
		scanner.Scan()
		A[i], _ = strconv.ParseUint(scanner.Text(), 10, 64)
	}
	solve(N, A)
}
