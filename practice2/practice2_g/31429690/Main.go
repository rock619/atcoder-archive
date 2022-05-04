package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Edge struct {
	From, To int64
}

type CSR struct {
	Start []int64
	Elist []int64
}

func NewCSR(n int64, edges []Edge) *CSR {
	start := make([]int64, n+1)
	elist := make([]int64, len(edges))
	for _, e := range edges {
		start[e.From+1]++
	}
	for i := int64(1); i <= n; i++ {
		start[i] += start[i-1]
	}
	counter := make([]int64, len(start))
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
	N     int64
	edges []Edge
}

func NewSCCGraph(n int64) *SCCGraph {
	return &SCCGraph{
		N: n,
	}
}

func (scc *SCCGraph) AddEdge(e Edge) {
	scc.edges = append(scc.edges, e)
}

func (scc *SCCGraph) Do() [][]int64 {
	groupNum, ids := scc.sccIDs()

	counts := make([]int, groupNum)
	for _, id := range ids {
		counts[id]++
	}

	groups := make([][]int64, groupNum)
	for i := range groups {
		groups[i] = make([]int64, 0, counts[i])
	}
	for i := int64(0); i < scc.N; i++ {
		groups[ids[i]] = append(groups[ids[i]], i)
	}
	return groups
}

func (scc *SCCGraph) sccIDs() (groupNum int64, ids []int64) {
	g := NewCSR(scc.N, scc.edges)

	nowOrd := int64(0)
	visited := make([]int64, 0, scc.N)
	low := make([]int64, scc.N)
	ord := make([]int64, scc.N)
	for i := range ord {
		ord[i] = -1
	}
	ids = make([]int64, scc.N)

	var dfs func(int64)
	dfs = func(v int64) {
		ord[v], low[v] = nowOrd, nowOrd
		nowOrd++
		visited = append(visited, v)

		for i := g.Start[v]; i < g.Start[v+1]; i++ {
			to := g.Elist[i]
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
				ord[u] = scc.N
				ids[u] = groupNum
				if u == v {
					break
				}
			}
			groupNum++
		}
	}

	for i := int64(0); i < scc.N; i++ {
		if ord[i] == -1 {
			dfs(i)
		}
	}
	for i, id := range ids {
		ids[i] = groupNum - 1 - id
	}

	return groupNum, ids
}

func solve(N int64, M int64, a []int64, b []int64) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	scc := NewSCCGraph(N)
	for i := int64(0); i < M; i++ {
		scc.AddEdge(Edge{From: a[i], To: b[i]})
	}
	groups := scc.Do()
	fmt.Fprintln(w, len(groups))
	for _, g := range groups {
		fmt.Fprint(w, len(g))
		for _, id := range g {
			fmt.Fprintf(w, " %d", id)
		}
		fmt.Fprintln(w)
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
	a := make([]int64, M)
	b := make([]int64, M)
	for i := int64(0); i < M; i++ {
		scanner.Scan()
		a[i], err = strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
		scanner.Scan()
		b[i], err = strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	solve(N, M, a, b)
}
