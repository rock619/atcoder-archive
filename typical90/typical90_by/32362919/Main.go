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

type Point struct {
	X, Y int
}

type Direction struct {
	X, Y, ID int
}

func solve(N, T int, AX, AY, BX, BY []int) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	directions := []Direction{
		{X: 1, Y: 0, ID: 1},
		{X: 1, Y: 1, ID: 2},
		{X: 0, Y: 1, ID: 3},
		{X: -1, Y: 1, ID: 4},
		{X: -1, Y: 0, ID: 5},
		{X: -1, Y: -1, ID: 6},
		{X: 0, Y: -1, ID: 7},
		{X: 1, Y: -1, ID: 8},
	}

	bs := make(map[Point]int)
	for i := range BX {
		bs[Point{X: BX[i], Y: BY[i]}] = i
	}

	graph := NewMFGraph(2*N + 2)
	for i := 0; i < N; i++ {
		graph.AddEdge(0, i+1, 1)
		graph.AddEdge(N+i+1, 2*N+1, 1)

		for _, d := range directions {
			p := Point{
				X: AX[i] + T*d.X,
				Y: AY[i] + T*d.Y,
			}
			if bi, ok := bs[p]; ok {
				graph.AddEdge(i+1, N+bi+1, 1)
			}
		}
	}

	if flow := graph.Flow(0, 2*N+1); flow != N {
		fmt.Fprintln(w, no)
		return
	}

	results := make([]int, N)
	for _, e := range graph.Edges() {
		if e.Flow == 0 || e.From == 0 || e.To == 2*N+1 {
			continue
		}

		ai := e.From - 1
		bi := e.To - N - 1
		results[ai] = directionID(directions, T, AX[ai], AY[ai], BX[bi], BY[bi])
	}

	fmt.Fprintln(w, yes)
	for i, r := range results {
		if i == 0 {
			fmt.Fprint(w, r)
		} else {
			fmt.Fprint(w, " ", r)
		}
	}
	fmt.Fprintln(w)
}

func directionID(directions []Direction, T, ax, ay, bx, by int) int {
	dx, dy := bx-ax, by-ay
	for _, d := range directions {
		if dx == d.X*T && dy == d.Y*T {
			return d.ID
		}
	}
	return 0
}

func main() {
	s := NewScanner()
	N := s.Int()
	T := s.Int()
	AX, AY := s.Ints2(N)
	BX, BY := s.Ints2(N)

	solve(N, T, AX, AY, BX, BY)
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

type Scanner struct {
	*bufio.Scanner
}

func NewScanner() *Scanner {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 4096), 1048576)
	scanner.Split(bufio.ScanWords)
	return &Scanner{
		Scanner: scanner,
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

func (s *Scanner) Ints(size int) []int {
	v := make([]int, size)
	for i := 0; i < size; i++ {
		v[i] = s.Int()
	}
	return v
}

func (s *Scanner) Ints2(size int) ([]int, []int) {
	v1 := make([]int, size)
	v2 := make([]int, size)
	for i := 0; i < size; i++ {
		v1[i] = s.Int()
		v2[i] = s.Int()
	}
	return v1, v2
}

func (s *Scanner) Ints3(size int) ([]int, []int, []int) {
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

func (s *Scanner) Ints4(size int) ([]int, []int, []int, []int) {
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

func (s *Scanner) IntsN(h, w int) [][]int {
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
	que := Queue(make([]int, 0, 1))
	que.Enqueue(s)
	for !que.Empty() {
		v := que.Dequeue()
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
		d := g.dfs(e.To, s, Min(limit-res, g.g[e.To][e.Rev].Cap), iter, level)
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
	que := Queue(make([]int, 0, 1))
	que.Enqueue(s)
	for !que.Empty() {
		p := que.Dequeue()
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

type mfGraphEdge struct {
	To, Rev, Cap int
}

type Queue []int

func (q *Queue) Enqueue(v int) {
	*q = append(*q, v)
}

func (q *Queue) Dequeue() int {
	if q.Empty() {
		panic("*Queue.Dequeue(): empty")
	}
	v := (*q)[0]
	*q = (*q)[1:]
	return v
}

func (q Queue) Size() int {
	return len(q)
}

func (q Queue) Empty() bool {
	return q.Size() == 0
}

func (q *Queue) Clear() {
	*q = (*q)[:0]
}
func Min(v ...int) int {
	switch len(v) {
	case 0:
		panic("Min: len(v) == 0")
	case 1:
		return v[0]
	case 2:
		if v[0] < v[1] {
			return v[0]
		}
		return v[1]
	default:
		m := v[0]
		for i := 1; i < len(v); i++ {
			if v[i] < m {
				m = v[i]
			}
		}
		return m
	}
}
