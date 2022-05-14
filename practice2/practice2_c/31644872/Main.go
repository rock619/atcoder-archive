package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func FloorSum(n, m, a, b int64) int64 {
	ans := int64(0)
	if a < 0 {
		a2 := SafeMod(a, m)
		ans -= n * (n - 1) / 2 * ((a2 - a) / m)
		a = a2
	}
	if b < 0 {
		b2 := SafeMod(b, m)
		ans -= n * ((b2 - b) / m)
		b = b2
	}
	return ans + int64(floorSumU(uint64(n), uint64(m), uint64(a), uint64(b)))
}

func SafeMod(x, m int64) int64 {
	x %= m
	if x < 0 {
		return x + m
	}
	return x
}

// @param n `n < 2^32`
// @param m `1 <= m < 2^32`
// floorSumU sum_{i=0}^{n-1} floor((ai + b) / m) (mod 2^64)
func floorSumU(n, m, a, b uint64) uint64 {
	for ans := uint64(0); ; {
		if a >= m {
			ans += n * (n - 1) / 2 * (a / m)
			a %= m
		}
		if b >= m {
			ans += n * (b / m)
			b %= m
		}

		yMax := a*n + b
		if yMax < m {
			return ans
		}
		// yMax < m * (n + 1)
		// floor(yMax / m) <= n
		n = yMax / m
		b = yMax % m
		m, a = a, m
	}
}

func solve(T int64, N []int64, M []int64, A []int64, B []int64) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	for i := int64(0); i < T; i++ {
		fmt.Fprintln(w, FloorSum(N[i], M[i], A[i], B[i]))
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1048576
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)

	var err error
	scanner.Scan()
	T, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	N := make([]int64, T)
	M := make([]int64, T)
	A := make([]int64, T)
	B := make([]int64, T)
	for i := int64(0); i < T; i++ {
		scanner.Scan()
		N[i], err = strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
		scanner.Scan()
		M[i], err = strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
		scanner.Scan()
		A[i], err = strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
		scanner.Scan()
		B[i], err = strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	solve(T, N, M, A, B)
}
