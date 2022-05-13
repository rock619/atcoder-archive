package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type Pair struct {
	First, Second int64
}

type Edge struct {
	From, To, Cap, Flow int64
}

type MFGraph struct {
	n   int64
	g   [][]edge
	pos []Pair
}

func NewMFGraph(n int64) *MFGraph {
	return &MFGraph{
		n: n,
		g: make([][]edge, n),
	}
}

func (g *MFGraph) AddEdge(from, to, capacity int64) int64 {
	m := int64(len(g.pos))
	g.pos = append(g.pos, Pair{First: from, Second: int64(len(g.g[from]))})
	fromID := int64(len(g.g[from]))
	toID := int64(len(g.g[to]))
	if from == to {
		toID++
	}
	g.g[from] = append(g.g[from], edge{To: to, Rev: toID, Cap: capacity})
	g.g[to] = append(g.g[to], edge{To: from, Rev: fromID, Cap: 0})
	return m
}

func (g *MFGraph) Edge(i int64) Edge {
	e := g.g[g.pos[i].First][g.pos[i].Second]
	re := g.g[e.To][e.Rev]
	return Edge{From: g.pos[i].First, To: e.To, Cap: e.Cap + re.Cap, Flow: re.Cap}
}

func (g *MFGraph) Edges() []Edge {
	m := int64(len(g.pos))
	result := make([]Edge, m)
	for i := int64(0); i < m; i++ {
		result[i] = g.Edge(i)
	}
	return result
}

func (g *MFGraph) Change(i, capacity, flow int64) {
	e := g.g[g.pos[i].First][g.pos[i].Second]
	g.g[g.pos[i].First][g.pos[i].Second].Cap = capacity - flow
	g.g[e.To][e.Rev].Cap = flow
}

func (g *MFGraph) Flow(s, t int64) int64 {
	return g.FlowLimit(s, t, math.MaxInt64)
}

func (g *MFGraph) FlowLimit(s, t, limit int64) int64 {
	flow := int64(0)
	for flow < limit {
		level := g.bfs(s, t)
		if level[t] == -1 {
			break
		}
		iter := make([]int64, g.n)
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

func (g *MFGraph) bfs(s, t int64) []int64 {
	level := make([]int64, g.n)
	for i := range level {
		level[i] = -1
	}
	level[s] = 0
	que := Queue(make([]int64, 0, 1))
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

func (g *MFGraph) dfs(cur, s, limit int64, iter, level []int64) int64 {
	if cur == s {
		return limit
	}
	res := int64(0)
	curLevel := level[cur]
	for itMax := int64(len(g.g[cur])); iter[cur] < itMax; iter[cur]++ {
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

func (g *MFGraph) MinCut(s int64) []bool {
	visited := make([]bool, g.n)
	que := Queue(make([]int64, 0, 1))
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

type edge struct {
	To, Rev, Cap int64
}

type Queue []int64

func (q *Queue) Enqueue(v int64) {
	*q = append([]int64(*q), v)
}

func (q *Queue) Dequeue() int64 {
	v := (*q)[0]
	*q = (*q)[1:]
	return v
}

func (q Queue) Size() int64 {
	return int64(len(q))
}

func (q Queue) Empty() bool {
	return q.Size() == 0
}

func (q *Queue) Clear() {
	*q = (*q)[:0]
}

func Min(ints ...int64) int64 {
	if len(ints) == 0 {
		panic("Min: len(ints) == 0")
	}
	m := ints[0]
	for i := 1; i < len(ints); i++ {
		if ints[i] < m {
			m = ints[i]
		}
	}
	return m
}

func solve(N int64, M int64, S []string) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	g := NewMFGraph(N*M + 2)
	s, t := N*M, N*M+1

	// s -> even / odd -> t
	for i := int64(0); i < N; i++ {
		for j := int64(0); j < M; j++ {
			if S[i][j] == '#' {
				continue
			}
			v := i*M + j
			if (i+j)%2 == 0 {
				g.AddEdge(s, v, 1)
			} else {
				g.AddEdge(v, t, 1)
			}
		}
	}

	// even -> odd
	for i := int64(0); i < N; i++ {
		for j := int64(0); j < M; j++ {
			if (i+j)%2 != 0 || S[i][j] == '#' {
				continue
			}
			v0 := i*M + j
			if i > 0 && S[i-1][j] == '.' {
				v1 := (i-1)*M + j
				g.AddEdge(v0, v1, 1)
			}
			if j > 0 && S[i][j-1] == '.' {
				v1 := i*M + (j - 1)
				g.AddEdge(v0, v1, 1)
			}
			if i+1 < N && S[i+1][j] == '.' {
				v1 := (i+1)*M + j
				g.AddEdge(v0, v1, 1)
			}
			if j+1 < M && S[i][j+1] == '.' {
				v1 := i*M + (j + 1)
				g.AddEdge(v0, v1, 1)
			}
		}
	}

	fmt.Fprintln(w, g.Flow(s, t))

	result := make([][]byte, N)
	for i := range result {
		result[i] = []byte(S[i])
	}
	edges := g.Edges()
	for _, e := range edges {
		if e.From == s || e.To == t || e.Flow == 0 {
			continue
		}
		i0, j0 := e.From/M, e.From%M
		i1, j1 := e.To/M, e.To%M

		switch {
		case i0 == i1+1:
			result[i1][j1] = 'v'
			result[i0][j0] = '^'
		case j0 == j1+1:
			result[i1][j1] = '>'
			result[i0][j0] = '<'
		case i0 == i1-1:
			result[i0][j0] = 'v'
			result[i1][j1] = '^'
		default:
			result[i0][j0] = '>'
			result[i1][j1] = '<'
		}
	}

	for i := int64(0); i < N; i++ {
		fmt.Fprintln(w, string(result[i]))
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1048576
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)

	var err error
	scanner.Scan()
	N, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	scanner.Scan()
	M, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	S := make([]string, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		S[i] = scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	solve(N, M, S)
}
