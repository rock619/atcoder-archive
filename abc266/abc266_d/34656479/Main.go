package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type P struct {
	X, A int
}

func solve(N int, T []int, X []int, A []int) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()
	m := make(map[int]P)
	for i := range T {
		m[T[i]] = P{X: X[i], A: A[i]}
	}
	dp := make([][]int, 100000+1)
	for i := range dp {
		dp[i] = make([]int, 5)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	dp[0][0] = 0
	for i := 1; i <= 100000; i++ {
		for j := 0; j < 5; j++ {
			for k := j - 1; k <= j+1; k++ {
				if k < 0 || k >= 5 {
					continue
				}
				UpdateMax(&dp[i][j], dp[i-1][k])
			}
			if dp[i][j] == -1 {
				continue
			}
			p, ok := m[i]
			if !ok || p.X != j {
				continue
			}
			dp[i][j] += p.A
		}
	}

	fmt.Fprintln(w, Max(dp[100000]...))
}

func Max(v ...int) int {
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

func UpdateMax(max *int, v int) {
	if v > *max {
		*max = v
	}
}

func main() {
	sc := NewScanner()
	N := sc.Int()
	T := make([]int, N)
	X := make([]int, N)
	A := make([]int, N)
	for i := 0; i < N; i++ {
		T[i] = sc.Int()
		X[i] = sc.Int()
		A[i] = sc.Int()
	}
	solve(N, T, X, A)
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

func (s *Scanner) Scan() {
	if ok := s.Scanner.Scan(); !ok {
		panic(s.Err())
	}
}

func (s *Scanner) Int() int {
	s.Scan()
	v, err := strconv.Atoi(s.Scanner.Text())
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
	s.Scan()
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

func (s *Scanner) Float() float64 {
	s.Scan()
	v, err := strconv.ParseFloat(s.Text(), 64)
	if err != nil {
		panic(err)
	}
	return v
}

func (s *Scanner) Text() string {
	s.Scan()
	return s.Scanner.Text()
}
