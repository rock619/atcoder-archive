package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N int, W int, w []int, v []int) {
	ww := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := ww.Flush(); err != nil {
			panic(err)
		}
	}()

	inf := 1 << 62
	maxVSum := 1000 * N
	dp := make([][]int, N+1)
	for i := range dp {
		dp[i] = make([]int, maxVSum+1)
		for j := range dp[i] {
			dp[i][j] = inf
		}
	}

	dp[0][0] = 0

	for i := 0; i < N; i++ {
		for j := 0; j <= maxVSum; j++ {
			if dp[i][j] == inf {
				continue
			}

			UpdateMin(&dp[i+1][j], dp[i][j])

			UpdateMin(&dp[i+1][j+v[i]], dp[i][j]+w[i])
		}
	}

	max := 0
	for i := 0; i <= maxVSum; i++ {
		if dp[N][i] == inf || dp[N][i] > W {
			continue
		}

		UpdateMax(&max, i)
	}
	fmt.Fprintln(ww, max)
}

func UpdateMin(min *int, v int) {
	if v < *min {
		*min = v
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
	W := sc.Int()
	w := make([]int, N)
	v := make([]int, N)
	for i := 0; i < N; i++ {
		w[i] = sc.Int()
		v[i] = sc.Int()
	}
	solve(N, W, w, v)
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
