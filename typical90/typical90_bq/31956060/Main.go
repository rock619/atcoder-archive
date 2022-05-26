package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const mod = 1000000007

func PowMod(x, n, m int64) int64 {
	if m == 1 {
		return 0
	}

	r := int64(1)
	for y := SafeMod(x, m); n > 0; n >>= 1 {
		if n&1 != 0 {
			r = (r * y) % m
		}
		y = (y * y) % m
	}
	return r
}

func SafeMod(x, m int64) int64 {
	x %= m
	if x < 0 {
		return x + m
	}
	return x
}

func solve(N int64, K int64) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	switch {
	case K < 3 && K < N:
		fmt.Fprintln(w, 0)
		return
	case N == 1:
		fmt.Fprintln(w, K)
		return
	case N == 2:
		fmt.Fprintln(w, K*(K-1)%mod)
		return
	}

	result := K * (K - 1) % mod
	result *= PowMod(K-2, N-2, mod)
	result %= mod

	fmt.Fprintln(w, result)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1048576
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)

	var err error
	scanner.Scan()
	N, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	scanner.Scan()
	K, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	solve(N, K)
}
