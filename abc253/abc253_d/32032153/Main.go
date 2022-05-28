package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N int64, A int64, B int64) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	sum := (1 + N) * N / 2
	for a := A; a <= N; a += A {
		sum -= a
	}
	for b := B; b <= N; b += B {
		sum -= b
	}
	l := lcm(A, B)
	for c := l; c <= N; c += l {
		sum += c
	}
	fmt.Fprintln(w, sum)
}

func lcm(a, b int64) int64 {
	gcd := gcd(a, b)
	return a / gcd * b
}

func gcd(a, b int64) int64 {
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
	A, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	scanner.Scan()
	B, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	solve(N, A, B)
}
