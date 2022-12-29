package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

type V struct {
	Type, Time, ID int
}

type E struct {
	To, Weight int
}

func Make2D(len1, len2, value int) [][]int {
	result := make([][]int, len1)
	for i := range result {
		result[i] = make([]int, len2)
		for j := range result[i] {
			result[i][j] = value
		}
	}
	return result
}

func UpdateMax(max *int, v ...int) {
	for i := 0; i < len(v); i++ {
		if v[i] > *max {
			*max = v[i]
		}
	}
}

func solve(o Printer, N int, M int, K int, A []int, S []int, B []int, T []int) {
	vs := make([]V, 0, 2*(M+N))
	for i := range S {
		vs = append(vs, V{Type: 2, Time: S[i], ID: i + 1})
		vs = append(vs, V{Type: 1, Time: T[i] + K, ID: i + 1})
	}

	for i := 1; i <= N; i++ {
		vs = append(vs, V{Type: 0, Time: -1, ID: i})
		vs = append(vs, V{Type: 0, Time: 1 << 62, ID: i})
	}

	sort.Slice(vs, func(i, j int) bool {
		if vs[i].Time != vs[j].Time {
			return vs[i].Time < vs[j].Time
		}
		if vs[i].Type != vs[j].Type {
			return vs[i].Type < vs[j].Type
		}
		return vs[i].ID < vs[j].ID
	})

	vertS := make([]int, M+1)
	vertT := make([]int, len(vertS))
	for i, v := range vs {
		switch v.Type {
		case 0:
		case 1:
			vertT[v.ID] = i + 1
		case 2:
			vertS[v.ID] = i + 1
		}
	}
	airports := Make2D(N+1, 0, 0)
	for i, v := range vs {
		var j int
		switch v.Type {
		case 0:
			j = v.ID
		case 1:
			j = B[v.ID-1]
		case 2:
			j = A[v.ID-1]
		}
		airports[j] = append(airports[j], i+1)
	}

	g := make([][]E, len(vs)+2)
	for i := 1; i <= M; i++ {
		g[vertT[i]] = append(g[vertT[i]], E{To: vertS[i], Weight: 1})
	}
	for i := 1; i <= N; i++ {
		for j := 0; j < len(airports[i])-1; j++ {
			k, l := airports[i][j], airports[i][j+1]
			g[l] = append(g[l], E{To: k, Weight: 0})
		}
	}
	for i := 1; i <= N; i++ {
		g[airports[i][0]] = append(g[airports[i][0]], E{To: 0, Weight: 0})
		g[len(vs)+1] = append(g[len(vs)+1], E{To: airports[i][len(airports[i])-1], Weight: 0})
	}

	dp := make([]int, len(vs)+2)
	for i := 1; i <= len(vs)+1; i++ {
		for j := 0; j < len(g[i]); j++ {
			UpdateMax(&dp[i], dp[g[i][j].To]+g[i][j].Weight)
		}
	}
	o.l(dp[len(vs)+1])
}

func main() {
	sc := NewScanner()
	N := sc.Int()
	M := sc.Int()
	K := sc.Int()
	A := make([]int, M)
	S := make([]int, M)
	B := make([]int, M)
	T := make([]int, M)
	for i := 0; i < M; i++ {
		A[i] = sc.Int()
		S[i] = sc.Int()
		B[i] = sc.Int()
		T[i] = sc.Int()
	}
	out := NewPrinter()
	solve(out, N, M, K, A, S, B, T)
	if err := out.w.Flush(); err != nil {
		panic(err)
	}
}

type Scanner struct {
	*bufio.Scanner
}

func NewScanner() *Scanner {
	s := bufio.NewScanner(os.Stdin)
	s.Buffer(make([]byte, 4096), math.MaxInt64)
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

type Printer interface {
	// p fmt.Print
	p(a ...interface{})
	// f fmt.Printf
	f(format string, a ...interface{})
	// l fmt.Println
	l(a ...interface{})
}

type printer struct {
	w *bufio.Writer
}

func NewPrinter() *printer {
	return &printer{bufio.NewWriter(os.Stdout)}
}

func (p *printer) p(a ...interface{}) {
	fmt.Fprint(p.w, a...)
}

func (p *printer) f(format string, a ...interface{}) {
	fmt.Fprintf(p.w, format, a...)
}

func (p *printer) l(a ...interface{}) {
	fmt.Fprintln(p.w, a...)
}
