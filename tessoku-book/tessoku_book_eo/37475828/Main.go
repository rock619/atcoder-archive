package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type IntPair struct {
	First, Second int
}

type MFGraphEdge struct {
	From, To, Cap, Flow int
}

type MFGraph struct {
	n   int
	g   [][]mfGraphEdge
	pos []IntPair
}

func NewMFGraph(n int) *MFGraph {
	return &MFGraph{
		n: n,
		g: make([][]mfGraphEdge, n),
	}
}

func (g *MFGraph) AddEdge(from, to, capacity int) int {
	m := len(g.pos)
	g.pos = append(g.pos, IntPair{First: from, Second: len(g.g[from])})
	fromID := len(g.g[from])
	toID := len(g.g[to])
	if from == to {
		toID++
	}
	g.g[from] = append(g.g[from], mfGraphEdge{To: to, Rev: toID, Cap: capacity})
	g.g[to] = append(g.g[to], mfGraphEdge{To: from, Rev: fromID, Cap: 0})
	return m
}

func (g *MFGraph) Edge(i int) MFGraphEdge {
	e := g.g[g.pos[i].First][g.pos[i].Second]
	re := g.g[e.To][e.Rev]
	return MFGraphEdge{From: g.pos[i].First, To: e.To, Cap: e.Cap + re.Cap, Flow: re.Cap}
}

func (g *MFGraph) Edges() []MFGraphEdge {
	m := len(g.pos)
	result := make([]MFGraphEdge, m)
	for i := 0; i < m; i++ {
		result[i] = g.Edge(i)
	}
	return result
}

func (g *MFGraph) Change(i, capacity, flow int) {
	e := g.g[g.pos[i].First][g.pos[i].Second]
	g.g[g.pos[i].First][g.pos[i].Second].Cap = capacity - flow
	g.g[e.To][e.Rev].Cap = flow
}

func (g *MFGraph) Flow(s, t int) int {
	return g.FlowLimit(s, t, math.MaxInt64)
}

func (g *MFGraph) FlowLimit(s, t, limit int) int {
	flow := 0
	for flow < limit {
		level := g.bfs(s, t)
		if level[t] == -1 {
			break
		}
		iter := make([]int, g.n)
		for flow < limit {
			f := g.dfs(t, s, limit-flow, iter, level)
			if f == 0 {
				break
			}
			flow += f
		}
	}
	return flow
}

func (g *MFGraph) bfs(s, t int) []int {
	level := make([]int, g.n)
	for i := range level {
		level[i] = -1
	}
	level[s] = 0
	que := NewQueue(1)
	que.Enqueue(s)
	for !que.Empty() {
		v, _ := que.Dequeue()
		for _, e := range g.g[v] {
			if e.Cap == 0 || level[e.To] >= 0 {
				continue
			}
			level[e.To] = level[v] + 1
			if e.To == t {
				return level
			}
			que.Enqueue(e.To)
		}
	}
	return level
}

func (g *MFGraph) dfs(cur, s, limit int, iter, level []int) int {
	if cur == s {
		return limit
	}
	res := 0
	curLevel := level[cur]
	for itMax := len(g.g[cur]); iter[cur] < itMax; iter[cur]++ {
		i := iter[cur]
		e := g.g[cur][i]
		if curLevel <= level[e.To] || g.g[e.To][e.Rev].Cap == 0 {
			continue
		}
		d := g.dfs(e.To, s, g.min(limit-res, g.g[e.To][e.Rev].Cap), iter, level)
		if d <= 0 {
			continue
		}
		g.g[cur][i].Cap += d
		g.g[e.To][e.Rev].Cap -= d
		res += d
		if res == limit {
			break
		}
	}
	return res
}

func (g *MFGraph) MinCut(s int) []bool {
	visited := make([]bool, g.n)
	que := NewQueue(1)
	que.Enqueue(s)
	for !que.Empty() {
		p, _ := que.Dequeue()
		visited[p] = true
		for _, e := range g.g[p] {
			if e.Cap != 0 && !visited[e.To] {
				visited[e.To] = true
				que.Enqueue(e.To)
			}
		}
	}
	return visited
}

func (*MFGraph) min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type mfGraphEdge struct {
	To, Rev, Cap int
}

type QueueElement = int

type Queue struct {
	Size int
	head int
	s    []QueueElement
}

func NewQueue(capacity int) *Queue {
	return &Queue{
		s: make([]QueueElement, 0, capacity),
	}
}

func (q *Queue) Empty() bool {
	return q.Size == 0
}

func (q *Queue) Clear() {
	q.Size = 0
}

func (q *Queue) Enqueue(x QueueElement) {
	q.Size++
	if q.Size > len(q.s) {
		q.reorder()
		q.s = append(q.s, x)
		return
	}
	tail := q.head + q.Size - 1
	if tail >= len(q.s) {
		tail -= len(q.s)
	}
	q.s[tail] = x
}

func (q *Queue) Dequeue() (x QueueElement, ok bool) {
	if q.Empty() {
		return x, false
	}
	x = q.s[q.head]
	if q.head++; q.head >= len(q.s) {
		q.head -= len(q.s)
	}
	q.Size--
	return x, true
}

func (q *Queue) Peek() (x QueueElement, ok bool) {
	if q.Empty() {
		return x, false
	}
	return q.s[q.head], true
}

func (q *Queue) reorder() {
	q.s = append(q.s[q.head:], q.s[:q.head]...)
	q.head = 0
}

const inf = 1 << 62

func solve(o Printer, N int, M int, P []int, A []int, B []int) {
	g := NewMFGraph(N + 2)
	s, t := 0, N+1
	offset := 0
	for i := 0; i < N; i++ {
		if p := P[i]; p >= 0 {
			g.AddEdge(s, i+1, p)
			offset += p
		} else {
			g.AddEdge(i+1, t, -p)
		}
	}

	for i := range A {
		g.AddEdge(A[i], B[i], inf)
	}

	result := offset - g.Flow(s, t)
	o.l(result)
}

func main() {
	sc := NewScanner()
	N := sc.Int()
	M := sc.Int()
	P := make([]int, N)
	for i := 0; i < N; i++ {
		P[i] = sc.Int()
	}
	A := make([]int, M)
	B := make([]int, M)
	for i := 0; i < M; i++ {
		A[i] = sc.Int()
		B[i] = sc.Int()
	}
	out := NewPrinter()
	solve(out, N, M, P, A, B)
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
