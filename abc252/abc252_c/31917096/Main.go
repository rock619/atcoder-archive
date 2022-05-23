package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solve(N int64, S []string) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	mins := make([]int64, 10)
	for n := 0; n <= 9; n++ {
		m := make(map[int]int, 10)
		for i := int64(0); i < N; i++ {
			m[strings.Index(S[i], strconv.Itoa(n))]++
		}

		min := 0
		for i := 0; i < 10; i++ {
			c, ok := m[i]
			if !ok {
				continue
			}
			v := (c-1)*10 + i
			if v > min {
				min = v
			}
		}
		mins[n] = int64(min)
	}
	fmt.Fprintln(w, Min(mins...))
}

func Min(v ...int64) int64 {
	switch len(v) {
	case 0:
		panic("Min: len(v) == 0")
	case 1:
		return v[0]
	case 2:
		if v[0] < v[1] {
			return v[0]
		}
		return v[1]
	default:
		m := v[0]
		for i := 1; i < len(v); i++ {
			if v[i] < m {
				m = v[i]
			}
		}
		return m
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
	S := make([]string, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		S[i] = scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	solve(N, S)
}
