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

func solve(N int64, K int64, A []int64) {
	if rest := K - AbsSum(A...); rest >= 0 && rest%2 == 0 {
		fmt.Println(YES)
		return
	}
	fmt.Println(NO)
}

func AbsSum(ints ...int64) int64 {
	sum := int64(0)
	for _, n := range ints {
		if n < 0 {
			sum -= n
		} else {
			sum += n
		}
	}
	return sum
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
	A := make([]int64, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		A[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	solve(N, K, A)
}
