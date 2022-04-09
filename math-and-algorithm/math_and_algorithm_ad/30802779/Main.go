package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

const (
	YES = "Yes"
	NO  = "No"
)

func solve(N int64, X int64, A []int64) {
	sort.Slice(A, func(i, j int) bool {
		return A[i] < A[j]
	})

	left, right := int64(0), N-1
	for left <= right {
		switch mid := (left + right) / 2; {
		case A[mid] == X:
			fmt.Println(YES)
			return
		case A[mid] > X:
			right = mid - 1
		default:
			left = mid + 1
		}
	}

	fmt.Println(NO)
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
	var X int64
	scanner.Scan()
	X, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	A := make([]int64, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		A[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	solve(N, X, A)
}
