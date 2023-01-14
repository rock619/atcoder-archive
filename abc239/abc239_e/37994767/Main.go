package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
)

const (
	// _     = iota
	debug = iota
)

type DFS struct {
	X []int
	g [][]int
	p [][]int
}

func NewDFS(X []int, g [][]int) *DFS {
	return &DFS{
		X: X,
		g: g,
		p: make([][]int, len(g)),
	}
}

func (s *DFS) Do(v, parent int) {
	s.p[v] = []int{s.X[v-1]}
	if parent != -1 && len(s.g[v]) == 1 {
		return
	}
	for _, next := range s.g[v] {
		if next == parent {
			continue
		}
		s.Do(next, v)
		s.p[v] = append(s.p[v], s.p[next]...)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(s.p[v])))
	if len(s.p[v]) > 20 {
		s.p[v] = s.p[v][:20]
	}
}

func solve(o, lg Printer, N int, Q int, X []int, A []int, B []int, V []int, K []int) {
	g := make([][]int, N+1)
	for i := range A {
		a, b := A[i], B[i]
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}
	dfs := NewDFS(X, g)
	dfs.Do(1, -1)
	lg.p(dfs.p)
	for i := range V {
		v, k := V[i], K[i]-1
		o.l(dfs.p[v][k])
	}
}

func main() {
	sc := NewScanner()
	N := sc.Int()
	Q := sc.Int()
	X := make([]int, N)
	for i := 0; i < N; i++ {
		X[i] = sc.Int()
	}
	A := make([]int, N-1)
	B := make([]int, N-1)
	for i := 0; i < N-1; i++ {
		A[i] = sc.Int()
		B[i] = sc.Int()
	}
	V := make([]int, Q)
	K := make([]int, Q)
	for i := 0; i < Q; i++ {
		V[i] = sc.Int()
		K[i] = sc.Int()
	}
	stdout := bufio.NewWriter(os.Stdout)
	out := NewPrinter(stdout)
	logger := NewLogger()
	solve(out, logger, N, Q, X, A, B, V, K)
	if err := stdout.Flush(); err != nil {
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
	w io.Writer
}

func NewPrinter(w io.Writer) Printer {
	return &printer{w}
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

type logger struct {
	*log.Logger
}

func NewLogger() Printer {
	return &logger{
		log.New(os.Stderr, "", log.Lmicroseconds|log.Lshortfile),
	}
}

func (l *logger) p(a ...interface{}) {
	if debug == 1 {
		l.Logger.Print(a...)
	}
}

func (l *logger) f(format string, a ...interface{}) {
	if debug == 1 {
		l.Logger.Printf(format, a...)
	}
}

func (l *logger) l(a ...interface{}) {
	if debug == 1 {
		l.Logger.Println(a...)
	}
}
