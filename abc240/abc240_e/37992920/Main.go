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

type DFS struct {
	g       [][]int
	leafIDs []int
	lastID  int
}

func NewDFS(g [][]int) *DFS {
	return &DFS{
		g:       g,
		leafIDs: make([]int, len(g)),
		lastID:  0,
	}
}

func (s *DFS) Do(v, parent int) {
	if v != 1 && len(s.g[v]) == 1 {
		s.lastID++
		s.leafIDs[v] = s.lastID
		return
	}

	for _, next := range s.g[v] {
		if next != parent {
			s.Do(next, v)
		}
	}
}

type DFS2 struct {
	g       [][]int
	leafIDs []int
	l, r    []int
}

func NewDFS2(g [][]int, leafIDs []int) *DFS2 {
	return &DFS2{
		g:       g,
		leafIDs: leafIDs,
		l:       make([]int, len(g)),
		r:       make([]int, len(g)),
	}
}

func (s *DFS2) Do(v, parent int) {
	if lID := s.leafIDs[v]; lID > 0 {
		s.l[v] = lID
		s.r[v] = lID
		return
	}
	s.l[v] = 1 << 62
	s.r[v] = -1 << 62
	for _, next := range s.g[v] {
		if next == parent {
			continue
		}
		s.Do(next, v)
		UpdateMin(&s.l[v], s.l[next])
		UpdateMax(&s.r[v], s.r[next])
	}
}

func UpdateMax(max *int, v ...int) {
	for i := 0; i < len(v); i++ {
		if v[i] > *max {
			*max = v[i]
		}
	}
}

func UpdateMin(min *int, v ...int) {
	for i := 0; i < len(v); i++ {
		if v[i] < *min {
			*min = v[i]
		}
	}
}

func solve(o, lg Printer, N int, u []int, v []int) {
	g := make([][]int, N+1)
	for i := range u {
		ui, vi := u[i], v[i]
		g[ui] = append(g[ui], vi)
		g[vi] = append(g[vi], ui)
	}

	dfs := NewDFS(g)
	dfs.Do(1, -1)
	lg.p(g)
	lg.p(dfs.leafIDs)
	dfs2 := NewDFS2(g, dfs.leafIDs)
	dfs2.Do(1, -1)
	lg.p(dfs2.l)
	lg.p(dfs2.r)
	for i := 1; i <= N; i++ {
		o.l(dfs2.l[i], dfs2.r[i])
	}
}

func main() {
	sc := NewScanner()
	N := sc.Int()
	u := make([]int, N-1)
	v := make([]int, N-1)
	for i := 0; i < N-1; i++ {
		u[i] = sc.Int()
		v[i] = sc.Int()
	}
	stdout := bufio.NewWriter(os.Stdout)
	out := NewPrinter(stdout)
	logger := NewLogger()
	solve(out, logger, N, u, v)
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
