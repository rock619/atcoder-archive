package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

const mod = 1000000007

func solve(N, K int, A []int) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	Compress(A)
	m := make(map[int]struct{})
	for _, v := range A {
		m[v] = struct{}{}
	}

	ft := NewFenwickTree(len(m) + 2)
	ft.Add(A[N-1], 1)
	cl := make([]int, N+1)
	for l, r, inversions := N, N, 0; r >= 1; r-- {
		for l >= 1 && inversions <= K {
			l--
			if l == 0 {
				continue
			}
			inversions += ft.sum(A[l-1])
			ft.Add(A[l-1], 1)
		}

		ft.Add(A[r-1], -1)
		inversions += ft.Sum(len(m)+2, A[r-1]+1)
		cl[r] = l
	}

	dp := make([]int, N+1)
	dp[0] = 1
	ru := make([]int, N+1)
	ru[0] = 1
	for i := 1; i <= N; i++ {
		if cl[i] == 0 {
			dp[i] = ru[i-1]
		} else {
			dp[i] = (ru[i-1] - ru[cl[i]-1] + mod) % mod
		}
		ru[i] = ru[i-1] + dp[i]
		ru[i] %= mod
	}
	fmt.Fprintln(w, dp[N])
}

func Compress(s []int) {
	s2 := make([]int, len(s))
	copy(s2, s)
	sort.Ints(s2)
	j := 0
	for i := 1; i < len(s2); i++ {
		if s2[j] == s2[i] {
			continue
		}
		j++
		s2[j] = s2[i]
	}
	s2 = s2[:j+1]
	for i, v := range s {
		index := sort.SearchInts(s2, v)
		s[i] = index
	}
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

func main() {
	s := NewScanner()
	N := s.Int()
	K := s.Int()
	A := s.IntN(N)

	solve(N, K, A)
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
	s.Scan()
	v, err := strconv.Atoi(s.Text())
	if err != nil {
		panic(err)
	}
	return v
}

func (s *Scanner) IntN(size int) []int {
	v := make([]int, size)
	for i := 0; i < size; i++ {
		v[i] = s.Int()
	}
	return v
}

func (s *Scanner) IntN2(size int) ([]int, []int) {
	v1 := make([]int, size)
	v2 := make([]int, size)
	for i := 0; i < size; i++ {
		v1[i] = s.Int()
		v2[i] = s.Int()
	}
	return v1, v2
}

func (s *Scanner) IntN3(size int) ([]int, []int, []int) {
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

func (s *Scanner) IntN4(size int) ([]int, []int, []int, []int) {
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

func (s Scanner) Float() float64 {
	s.Scan()
	v, err := strconv.ParseFloat(s.Text(), 64)
	if err != nil {
		panic(err)
	}
	return v
}

type FenwickTree struct {
	data []int
}

func NewFenwickTree(n int) *FenwickTree {
	return &FenwickTree{
		data: make([]int, n+1),
	}
}

func (f *FenwickTree) Add(p, x int) {
	p++
	for l := len(f.data); p < l; {
		f.data[p-1] += x
		p += p & -p
	}
}

func (f *FenwickTree) Sum(l, r int) int {
	return f.sum(r) - f.sum(l)
}

func (f *FenwickTree) sum(r int) int {
	s := 0
	for r > 0 {
		s += f.data[r-1]
		r -= r & -r
	}
	return s
}
