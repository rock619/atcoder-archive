package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(Q int64, Queries []Query) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	m := make(map[int64]int64)
	max, min := int64(0), int64(1)<<60
	for i := int64(0); i < Q; i++ {
		q := Queries[i]
		switch q.Type {
		case 1:
			m[q.X]++
			max = Max(max, q.X)
			min = Min(min, q.X)
		case 2:
			count := Min(q.C, m[q.X])
			m[q.X] -= count
			if m[q.X] <= 0 {
				delete(m, q.X)
				if q.X == max || q.X == min {
					max, min = int64(0), int64(1)<<60
					for k := range m {
						if k > max {
							max = k
						}
						if k < min {
							min = k
						}
					}
				}
			}
		default:
			fmt.Fprintln(w, max-min)
		}
	}

}

func Max(v ...int64) int64 {
	switch len(v) {
	case 0:
		panic("Max: len(v) == 0")
	case 1:
		return v[0]
	case 2:
		if v[0] > v[1] {
			return v[0]
		}
		return v[1]
	default:
		m := v[0]
		for i := 1; i < len(v); i++ {
			if v[i] > m {
				m = v[i]
			}
		}
		return m
	}
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

type Query struct {
	Type, X, C int64
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1048576
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)

	var err error
	scanner.Scan()
	Q, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	Queries := make([]Query, Q)
	for i := int64(0); i < Q; i++ {
		scanner.Scan()
		Queries[i].Type, err = strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
		if Queries[i].Type == 3 {
			continue
		}
		scanner.Scan()
		Queries[i].X, err = strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
		if Queries[i].Type == 1 {
			continue
		}
		scanner.Scan()
		Queries[i].C, err = strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	solve(Q, Queries)
}
