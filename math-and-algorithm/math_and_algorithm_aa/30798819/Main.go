package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solve(N int64, A []int64) {
	C := make([]int64, N)
	MergeSort(A, C, 0, N)
	fmt.Println(strings.Trim(fmt.Sprint(A), "[]"))
}

func MergeSort(A, C []int64, l, r int64) {
	if r-l == 1 {
		return
	}

	m := (l + r) / 2
	MergeSort(A, C, l, m)
	MergeSort(A, C, m, r)

	c1, c2, cnt := l, m, int64(0)
	for c1 != m || c2 != r {
		switch {
		case c1 == m:
			C[cnt] = A[c2]
			c2++
		case c2 == r, A[c1] < A[c2]:
			C[cnt] = A[c1]
			c1++
		default:
			C[cnt] = A[c2]
			c2++
		}
		cnt++
	}

	for i := int64(0); i < cnt; i++ {
		A[l+i] = C[i]
	}
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
	A := make([]int64, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		A[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	solve(N, A)
}
