package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N int, M int, E int, U []int, V []int, Q int, X []int) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	linesToCut := make(map[int]bool)
	for i := 0; i < Q; i++ {
		linesToCut[X[i]] = true
	}

	dsu := NewDSU(N + M)
	for i := N; i < N+M; i++ {
		dsu.Merge(N, i)
	}
	results := make([]int, Q+1)
	for i := 0; i < E; i++ {
		if linesToCut[i+1] {
			continue
		}
		u, v := U[i]-1, V[i]-1
		switch {
		case dsu.Same(u, v):
		case dsu.Same(N, u) && dsu.Same(N, v):
		case dsu.Same(N, u) && !dsu.Same(N, v):
			results[0] += dsu.Size(v)
		case dsu.Same(N, v) && !dsu.Same(N, u):
			results[0] += dsu.Size(u)
		default:
		}
		dsu.Merge(u, v)
	}

	for i := 1; i <= Q; i++ {
		results[i] = results[i-1]
		ei := X[Q-i] - 1

		u, v := U[ei]-1, V[ei]-1
		switch {
		case dsu.Same(u, v):
		case dsu.Same(N, u) && dsu.Same(N, v):
		case dsu.Same(N, u) && !dsu.Same(N, v):
			results[i] += dsu.Size(v)
		case dsu.Same(N, v) && !dsu.Same(N, u):
			results[i] += dsu.Size(u)
		default:
		}
		dsu.Merge(u, v)
	}

	results = results[:Q]
	Reverse(results)
	for _, r := range results {
		fmt.Fprintln(w, r)
	}
}

func Reverse(s []int) {
	for left, right := 0, len(s)-1; left < right; left, right = left+1, right-1 {
		s[left], s[right] = s[right], s[left]
	}
}

type DSU struct {
	parentOrSize []int
}

func NewDSU(n int) *DSU {
	parentOrSize := make([]int, n)
	for i := range parentOrSize {
		parentOrSize[i] = -1
	}

	return &DSU{
		parentOrSize,
	}
}

func (dsu *DSU) Merge(a, b int) int {
	x, y := dsu.Leader(a), dsu.Leader(b)
	if x == y {
		return x
	}
	if -dsu.parentOrSize[x] < -dsu.parentOrSize[y] {
		x, y = y, x
	}
	dsu.parentOrSize[x] += dsu.parentOrSize[y]
	dsu.parentOrSize[y] = x
	return x
}

func (dsu *DSU) Same(a, b int) bool {
	return dsu.Leader(a) == dsu.Leader(b)
}

func (dsu *DSU) Leader(a int) int {
	if dsu.parentOrSize[a] < 0 {
		return a
	}
	dsu.parentOrSize[a] = dsu.Leader(dsu.parentOrSize[a])
	return dsu.parentOrSize[a]
}

func (dsu *DSU) Size(a int) int {
	return -dsu.parentOrSize[dsu.Leader(a)]
}

func (dsu *DSU) Groups() [][]int {
	l := len(dsu.parentOrSize)
	m := make(map[int][]int, l)
	for i := 0; i < l; i++ {
		leader := dsu.Leader(i)
		m[leader] = append(m[leader], i)
	}

	result := make([][]int, 0, len(m))
	for _, v := range m {
		result = append(result, v)
	}
	return result
}

func main() {
	s := NewScanner()
	N := s.Int()
	M := s.Int()
	E := s.Int()
	U := make([]int, E)
	V := make([]int, E)
	for i := 0; i < E; i++ {
		U[i] = s.Int()
		V[i] = s.Int()
	}
	Q := s.Int()
	X := make([]int, Q)
	for i := 0; i < Q; i++ {
		X[i] = s.Int()
	}
	solve(N, M, E, U, V, Q, X)
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
