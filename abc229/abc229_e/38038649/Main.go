package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
)

const (
	// _     = iota
	debug = iota
)

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

func solve(o, lg Printer, N int, M int, A []int, B []int) {
	dsu := NewDSU(N)
	g := make([][]int, N)
	for i := range A {
		a, b := A[i]-1, B[i]-1
		if a > b {
			a, b = b, a
		}
		g[a] = append(g[a], b)
	}
	lg.f("%v", g)
	result := make([]int, 1, N)
	result[0] = 0
	groups := 0
	for i := N - 1; i >= 1; i-- {
		groups++
		for _, v := range g[i] {
			if !dsu.Same(i, v) {
				dsu.Merge(i, v)
				groups--
			}
		}
		result = append(result, groups)
	}
	for i := N - 1; i >= 0; i-- {
		o.l(result[i])
	}
	lg.p(result)
}

func main() {
	sc := NewScanner()
	N := sc.Int()
	M := sc.Int()
	A := make([]int, M)
	B := make([]int, M)
	for i := 0; i < M; i++ {
		A[i] = sc.Int()
		B[i] = sc.Int()
	}
	stdout := bufio.NewWriter(os.Stdout)
	out := NewPrinter(stdout)
	logger := NewLogger()
	solve(out, logger, N, M, A, B)
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
