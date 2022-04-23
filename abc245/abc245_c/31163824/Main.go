package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	YES = "Yes"
	NO  = "No"
)

func solve(N int64, K int64, A []int64, B []int64) {
	canChooseA, canChooseB := true, true
	for i := int64(1); i < N; i++ {
		a, b := true, true
		if !((canChooseA && Abs(A[i]-A[i-1]) <= K) || (canChooseB && Abs(A[i]-B[i-1]) <= K)) {
			a = false
		}

		if !((canChooseA && Abs(B[i]-A[i-1]) <= K) || (canChooseB && Abs(B[i]-B[i-1]) <= K)) {
			b = false
		}

		canChooseA, canChooseB = a, b
	}

	if canChooseA || canChooseB {
		fmt.Println(YES)
	} else {
		fmt.Println(NO)
	}
}

func Abs(n int64) int64 {
	if n < 0 {
		return -n
	}
	return n
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	N, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	scanner.Scan()
	K, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	A := make([]int64, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		A[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	B := make([]int64, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		B[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	solve(N, K, A, B)
}
