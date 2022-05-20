package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func solve(N int64, A []int64) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	lisLens := LISLens(A)
	Reverse(A)
	lisLensReversed := LISLens(A)

	max := int64(0)
	for i := int64(0); i < N; i++ {
		if l := lisLens[i] + lisLensReversed[N-1-i] - 1; l > max {
			max = l
		}
	}
	fmt.Fprintln(w, max)
}

func LISLens(a []int64) []int64 {
	lisLens := make([]int64, len(a))
	lis := make([]int64, 1, len(a))
	lis[0] = a[0]
	lisLens[0] = 1
	for i := 1; i < len(a); i++ {
		index := sort.Search(len(lis), func(j int) bool {
			return lis[j] >= a[i]
		})
		if index >= len(lis) {
			lis = append(lis, a[i])
		} else {
			lis[index] = a[i]
		}
		lisLens[i] = int64(len(lis))
	}
	return lisLens
}

func Reverse(s []int64) {
	for left, right := 0, len(s)-1; left < right; left, right = left+1, right-1 {
		s[left], s[right] = s[right], s[left]
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
	N, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	A := make([]int64, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		A[i], err = strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	solve(N, A)
}
