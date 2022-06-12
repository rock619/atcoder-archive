package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func solve(N, M int, a, b []int, Q int, x, y []int) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	graph := make([][]int, N)
	for i := range a {
		ai, bi := a[i]-1, b[i]-1
		graph[ai] = append(graph[ai], bi)
		graph[bi] = append(graph[bi], ai)
	}

	lastUpdated := make([]int, N)
	for i := range lastUpdated {
		lastUpdated[i] = -1
	}

	degreeThreshold := 1
	for degreeThreshold*degreeThreshold <= 2*M {
		degreeThreshold++
	}

	large := make([]int, 0, N)
	for i, v := range graph {
		if len(v) >= degreeThreshold {
			large = append(large, i)
		}
	}

	linked := make([][]bool, N)
	for i := range linked {
		linked[i] = make([]bool, len(large))
	}
	for i, v := range large {
		for _, next := range graph[v] {
			linked[next][i] = true
		}
		linked[v][i] = true
	}

	lastUpdated2 := make([]int, len(large))
	for i := range lastUpdated2 {
		lastUpdated2[i] = -1
	}
	for i := range x {
		v := x[i] - 1
		last := lastUpdated[v]
		for j := range large {
			if linked[v][j] {
				UpdateMax(&last, lastUpdated2[j])
			}
		}
		if last == -1 {
			fmt.Fprintln(w, 1)
		} else {
			fmt.Fprintln(w, y[last])
		}

		if len(graph[v]) < degreeThreshold {
			lastUpdated[v] = i
			for _, next := range graph[v] {
				lastUpdated[next] = i
			}
		} else {
			index := sort.Search(len(large), func(i int) bool {
				return large[i] >= v
			})
			lastUpdated2[index] = i
		}
	}
}

func main() {
	s := NewScanner()
	N := s.Int()
	M := s.Int()
	a, b := s.Ints2(M)
	Q := s.Int()
	x, y := s.Ints2(Q)

	solve(N, M, a, b, Q, x, y)
}

type Scanner struct {
	*bufio.Scanner
}

func NewScanner() *Scanner {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 4096), 1048576)
	scanner.Split(bufio.ScanWords)
	return &Scanner{
		Scanner: scanner,
	}
}

func (s *Scanner) Int() int {
	s.Scan()
	v, err := strconv.Atoi(s.Text())
	if err != nil {
		panic(err)
	}
	return v
}

func (s *Scanner) Ints(size int) []int {
	v := make([]int, size)
	for i := 0; i < size; i++ {
		v[i] = s.Int()
	}
	return v
}

func (s *Scanner) Ints2(size int) ([]int, []int) {
	v1 := make([]int, size)
	v2 := make([]int, size)
	for i := 0; i < size; i++ {
		v1[i] = s.Int()
		v2[i] = s.Int()
	}
	return v1, v2
}

func (s *Scanner) Ints3(size int) ([]int, []int, []int) {
	v1 := make([]int, size)
	v2 := make([]int, size)
	v3 := make([]int, size)
	for i := 0; i < size; i++ {
		v1[i] = s.Int()
		v2[i] = s.Int()
		v3[i] = s.Int()
	}
	return v1, v2, v3
}

func (s *Scanner) Ints4(size int) ([]int, []int, []int, []int) {
	v1 := make([]int, size)
	v2 := make([]int, size)
	v3 := make([]int, size)
	v4 := make([]int, size)
	for i := 0; i < size; i++ {
		v1[i] = s.Int()
		v2[i] = s.Int()
		v3[i] = s.Int()
		v4[i] = s.Int()
	}
	return v1, v2, v3, v4
}

func (s *Scanner) IntsN(h, w int) [][]int {
	v := make([][]int, h)
	for i := 0; i < h; i++ {
		v[i] = make([]int, w)
		for j := 0; j < w; j++ {
			v[i][j] = s.Int()
		}
	}
	return v
}

func (s *Scanner) Bytes() []byte {
	s.Scan()
	return s.Scanner.Bytes()
}

func (s *Scanner) BytesN(h int) [][]byte {
	v := make([][]byte, h)
	for i := 0; i < h; i++ {
		v[i] = s.Bytes()
	}
	return v
}

func (s *Scanner) Byte() byte {
	return s.Bytes()[0]
}

func (s *Scanner) ByteN(n int) []byte {
	v := make([]byte, n)
	for i := 0; i < n; i++ {
		v[i] = s.Byte()
	}
	return v
}

func UpdateMax(max *int, v int) {
	if v > *max {
		*max = v
	}
}
