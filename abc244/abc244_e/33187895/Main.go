package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const mod = 998244353

func solve(N int, M int, K int, S int, T int, X int, U []int, V []int) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	graph := make([][]int, N)
	for i := range U {
		u, v := U[i]-1, V[i]-1
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	dp := make([][][]int, K+1)
	for i := range dp {
		dp[i] = make([][]int, N)
		for j := range dp[i] {
			dp[i][j] = []int{0, 0}
		}
	}
	s, t := S-1, T-1
	dp[0][s][0] = 1
	x := X - 1
	for i := 0; i < K; i++ {
		for j := 0; j < N; j++ {
			for _, k := range graph[j] {
				if k == x {
					dp[i+1][k][0] += dp[i][j][1]
					dp[i+1][k][1] += dp[i][j][0]
				} else {
					dp[i+1][k][0] += dp[i][j][0]
					dp[i+1][k][1] += dp[i][j][1]
				}
				dp[i+1][k][0] %= mod
				dp[i+1][k][1] %= mod
			}
		}
	}

	fmt.Fprintln(w, dp[K][t][0])
}

func main() {
	s := NewScanner()
	N := s.Int()
	M := s.Int()
	K := s.Int()
	S := s.Int()
	T := s.Int()
	X := s.Int()
	U := make([]int, M)
	V := make([]int, M)
	for i := 0; i < M; i++ {
		U[i] = s.Int()
		V[i] = s.Int()
	}
	solve(N, M, K, S, T, X, U, V)
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
