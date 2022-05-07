package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N int64, K int64, a []int64) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	max := 1
	m := make(map[int64]int, K)
	for left, right := 0, 0; right < len(a); right++ {
		m[a[right]] = right
		switch valueLen := int64(len(m)); {
		case valueLen <= K:
			if l := right - left + 1; l > max {
				max = l
			}
		case valueLen > K:
			for int64(len(m)) > K {
				if left == m[a[left]] {
					delete(m, a[left])
				}
				left++
			}
		}
	}
	fmt.Fprintln(w, max)
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
	a := make([]int64, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		a[i], err = strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	solve(N, K, a)
}
