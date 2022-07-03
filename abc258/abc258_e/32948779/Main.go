package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N int, Q int, X int, W []int, K []int) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	w2 := make([]int, 2*N)
	copy(w2, W)
	copy(w2[N:], W)

	sum := Sum(W...)
	floor := (X / sum) * N
	x := X % sum
	counts := make([]int, N)
	for l, r, s := 0, 0, 0; l < N; l++ {
		counts[l] = floor
		if r < l {
			r, s = l, 0
		}
		for s < x {
			s += w2[r]
			r++
		}
		counts[l] += r - l
		s -= w2[l]
	}

	visited := make(map[int]int)
	offset, cycleLen := 0, 0
	result := make([]int, 0, N)
	for i, current := 1, 0; ; i++ {
		result = append(result, current)
		if prev, ok := visited[current]; ok {
			offset = prev
			cycleLen = i - offset
			break
		}
		visited[current] = i
		current = (current + counts[current]) % N
	}

	for _, k := range K {

		r := 0
		if k--; k <= offset {
			r = result[k]
		} else {
			r = result[(k-offset)%cycleLen+offset]
		}
		fmt.Fprintln(w, counts[r])
	}
}

func Sum(v ...int) int {
	switch len(v) {
	case 0:
		return 0
	case 1:
		return v[0]
	case 2:
		return v[0] + v[1]
	default:
		s := v[0]
		for i := 1; i < len(v); i++ {
			s += v[i]
		}
		return s
	}
}

func main() {
	s := NewScanner()

	N := s.Int()
	Q := s.Int()
	X := s.Int()
	W := make([]int, N)
	for i := 0; i < N; i++ {
		W[i] = s.Int()
	}
	K := make([]int, Q)
	for i := 0; i < Q; i++ {
		K[i] = s.Int()
	}

	solve(N, Q, X, W, K)
}

type Scanner struct {
	*bufio.Scanner
}

func NewScanner() *Scanner {
	s := bufio.NewScanner(os.Stdin)
	s.Buffer(make([]byte, 4096), 1048576)
	s.Split(bufio.ScanWords)
	return &Scanner{
		Scanner: s,
	}
}

func (s *Scanner) Int() int {
	if ok := s.Scan(); !ok {
		panic(s.Err())
	}
	v, err := strconv.Atoi(s.Text())
	if err != nil {
		panic(err)
	}
	return v
}

func (s *Scanner) IntN(n int) []int {
	v := make([]int, n)
	for i := 0; i < n; i++ {
		v[i] = s.Int()
	}
	return v
}

func (s *Scanner) IntN2(n int) ([]int, []int) {
	v1 := make([]int, n)
	v2 := make([]int, n)
	for i := 0; i < n; i++ {
		v1[i] = s.Int()
		v2[i] = s.Int()
	}
	return v1, v2
}

func (s *Scanner) IntN3(n int) ([]int, []int, []int) {
	v1 := make([]int, n)
	v2 := make([]int, n)
	v3 := make([]int, n)
	for i := 0; i < n; i++ {
		v1[i] = s.Int()
		v2[i] = s.Int()
		v3[i] = s.Int()
	}
	return v1, v2, v3
}

func (s *Scanner) IntN4(n int) ([]int, []int, []int, []int) {
	v1 := make([]int, n)
	v2 := make([]int, n)
	v3 := make([]int, n)
	v4 := make([]int, n)
	for i := 0; i < n; i++ {
		v1[i] = s.Int()
		v2[i] = s.Int()
		v3[i] = s.Int()
		v4[i] = s.Int()
	}
	return v1, v2, v3, v4
}

func (s *Scanner) IntNN(h, w int) [][]int {
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
	if ok := s.Scan(); !ok {
		panic(s.Err())
	}

	b := s.Scanner.Bytes()
	v := make([]byte, len(b))
	copy(v, b)
	return v
}

func (s *Scanner) BytesN(n int) [][]byte {
	v := make([][]byte, n)
	for i := 0; i < n; i++ {
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

func (s Scanner) Float() float64 {
	s.Scan()
	v, err := strconv.ParseFloat(s.Text(), 64)
	if err != nil {
		panic(err)
	}
	return v
}
