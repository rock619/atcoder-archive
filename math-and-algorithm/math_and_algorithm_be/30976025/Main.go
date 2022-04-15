package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"strconv"
)

func solve(N int64, K int64, V []int64) {
	result := int64(0)
	for i := 1; i < (1 << len(V)); i++ {
		lcm := int64(1)
		for j := 0; j < len(V); j++ {
			if (i & (1 << j)) != 0 {
				lcm = LCM(lcm, V[j])
			}
		}
		if c := bits.OnesCount(uint(i)); c%2 == 0 {
			result -= N / lcm
		} else {
			result += N / lcm
		}
	}
	fmt.Println(result)
}

func LCM(a, b int64) int64 {
	gcd := GCD(a, b)
	return a / gcd * b
}

func GCD(a, b int64) int64 {
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
	var N int64
	scanner.Scan()
	N, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	var K int64
	scanner.Scan()
	K, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	V := make([]int64, K)
	for i := int64(0); i < K; i++ {
		scanner.Scan()
		V[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	solve(N, K, V)
}
