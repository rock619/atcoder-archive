package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

const yes = "Yes"
const no = "No"

type TwoSAT struct {
	n      int
	g      *SCCGraph
	Answer []bool
}

func NewTwoSAT(n int) *TwoSAT {
	return &TwoSAT{
		n:      n,
		g:      NewSCCGraph(2*n, make([]Edge, 0)),
		Answer: make([]bool, n),
	}
}

func (ts *TwoSAT) AddClause(i int, f bool, j int, g bool) {
	switch {
	case f && g:
		ts.g.AddEdge(Edge{From: 2 * i, To: 2*j + 1})
		ts.g.AddEdge(Edge{From: 2 * j, To: 2*i + 1})
	case !f && g:
		ts.g.AddEdge(Edge{From: 2*i + 1, To: 2*j + 1})
		ts.g.AddEdge(Edge{From: 2 * j, To: 2 * i})
	case f && !g:
		ts.g.AddEdge(Edge{From: 2 * i, To: 2 * j})
		ts.g.AddEdge(Edge{From: 2*j + 1, To: 2*i + 1})
	default:
		ts.g.AddEdge(Edge{From: 2*i + 1, To: 2 * j})
		ts.g.AddEdge(Edge{From: 2*j + 1, To: 2 * i})
	}
}

func (ts *TwoSAT) Satisfiable() bool {
	_, ids := ts.g.sccIDs()
	for i := 0; i < ts.n; i++ {
		if ids[2*i] == ids[2*i+1] {
			return false
		}
		ts.Answer[i] = ids[2*i] < ids[2*i+1]
	}
	return true
}

type Edge struct {
	From, To int
}

type CSR struct {
	Start []int
	Elist []int
}

func NewCSR(n int, edges []Edge) *CSR {
	start := make([]int, n+1)
	elist := make([]int, len(edges))
	for _, e := range edges {
		start[e.From+1]++
	}
	for i := 1; i <= n; i++ {
		start[i] += start[i-1]
	}
	counter := make([]int, len(start))
	copy(counter, start)
	for _, e := range edges {
		elist[counter[e.From]] = e.To
		counter[e.From]++
	}
	return &CSR{
		Start: start,
		Elist: elist,
	}
}

type SCCGraph struct {
	N     int
	edges []Edge
}

func NewSCCGraph(n int, edges []Edge) *SCCGraph {
	return &SCCGraph{
		N:     n,
		edges: edges,
	}
}

func (g *SCCGraph) AddEdge(e Edge) {
	g.edges = append(g.edges, e)
}

func (g *SCCGraph) SCC() [][]int {
	groupNum, ids := g.sccIDs()

	counts := make([]int, groupNum)
	for _, id := range ids {
		counts[id]++
	}

	groups := make([][]int, groupNum)
	for i := range groups {
		groups[i] = make([]int, 0, counts[i])
	}
	for i := 0; i < g.N; i++ {
		groups[ids[i]] = append(groups[ids[i]], i)
	}
	return groups
}

func (g *SCCGraph) sccIDs() (groupNum int, ids []int) {
	csr := NewCSR(g.N, g.edges)

	nowOrd := 0
	visited := make([]int, 0, g.N)
	low := make([]int, g.N)
	ord := make([]int, g.N)
	for i := range ord {
		ord[i] = -1
	}
	ids = make([]int, g.N)

	var dfs func(int)
	dfs = func(v int) {
		ord[v], low[v] = nowOrd, nowOrd
		nowOrd++
		visited = append(visited, v)

		for i := csr.Start[v]; i < csr.Start[v+1]; i++ {
			to := csr.Elist[i]
			if ord[to] != -1 {
				if ord[to] < low[v] {
					low[v] = ord[to]
				}
				continue
			}

			dfs(to)
			if low[to] < low[v] {
				low[v] = low[to]
			}
		}

		if low[v] == ord[v] {
			for {
				u := visited[len(visited)-1]
				visited = visited[:len(visited)-1]
				ord[u] = g.N
				ids[u] = groupNum
				if u == v {
					break
				}
			}
			groupNum++
		}
	}

	for i := 0; i < g.N; i++ {
		if ord[i] == -1 {
			dfs(i)
		}
	}
	for i, id := range ids {
		ids[i] = groupNum - 1 - id
	}

	return groupNum, ids
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func solve(o Printer, N, D int, X, Y []int) {
	ts := NewTwoSAT(N)
	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			if Abs(X[i]-X[j]) < D {
				ts.AddClause(i, false, j, false)
			}
			if Abs(X[i]-Y[j]) < D {
				ts.AddClause(i, false, j, true)
			}
			if Abs(Y[i]-X[j]) < D {
				ts.AddClause(i, true, j, false)
			}
			if Abs(Y[i]-Y[j]) < D {
				ts.AddClause(i, true, j, true)
			}
		}
	}

	if !ts.Satisfiable() {
		o.l(no)
		return
	}
	o.l(yes)
	for i := 0; i < N; i++ {
		if ts.Answer[i] {
			o.l(X[i])
		} else {
			o.l(Y[i])
		}
	}
}

func main() {
	sc := NewScanner()
	N, D := sc.Int(), sc.Int()
	X, Y := sc.IntN2(N)
	out := NewPrinter()
	solve(out, N, D, X, Y)
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
