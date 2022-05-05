package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(A int64, B int64, C int64) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	gcd := GCD(A, B, C)
	fmt.Fprintln(w, A/gcd+B/gcd+C/gcd-3)
}

func GCD(ints ...int64) int64 {
	result := int64(0)
	for _, i := range ints {
		result = gcd(result, i)
	}
	return result
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
	A, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	scanner.Scan()
	B, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	scanner.Scan()
	C, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	solve(A, B, C)
}
