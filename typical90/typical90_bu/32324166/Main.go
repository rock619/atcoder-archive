package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const mod = 1000000007

type DFS struct {
	graph   [][]int
	dp      [][]int
	visited []bool
	c       []byte
}

func (s *DFS) Do(v int) {
	v1, v2 := 1, 1
	for _, child := range s.graph[v] {
		if s.visited[child] {
			continue
		}
		s.visited[v] = true
		s.Do(child)

		if s.c[v] == 'a' {
			v1 *= s.dp[child][0] + s.dp[child][2]
		} else {
			v1 *= s.dp[child][1] + s.dp[child][2]
		}
		v2 *= s.dp[child][0] + s.dp[child][1] + 2*s.dp[child][2]
		v1 %= mod
		v2 %= mod
	}

	if s.c[v] == 'a' {
		s.dp[v][0] = v1
	} else {
		s.dp[v][1] = v1
	}
	s.dp[v][2] = (v2 - v1 + mod) % mod
}

func solve(N int, c []byte, a []int, b []int) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	dp := make([][]int, N)
	for i := range dp {
		dp[i] = make([]int, 3)
	}

	graph := make([][]int, N)
	for i := 0; i < N-1; i++ {
		ai, bi := a[i]-1, b[i]-1
		graph[ai] = append(graph[ai], bi)
		graph[bi] = append(graph[bi], ai)
	}

	visited := make([]bool, N)
	visited[0] = true
	dfs := &DFS{
		graph:   graph,
		dp:      dp,
		visited: visited,
		c:       c,
	}
	dfs.Do(0)

	fmt.Fprintln(w, dfs.dp[0][2])
}

func main() {
	s := NewScanner()

	N := s.Int()
	c := s.ByteN(N)
	a, b := s.Ints2(N - 1)

	solve(N, c, a, b)
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

type Stack []int

func (s Stack) Size() int {
	return len(s)
}

func (s Stack) Empty() bool {
	return s.Size() == 0
}

func (s *Stack) Push(v int) {
	*s = append(*s, v)
}

func (s *Stack) Pop() int {
	if s.Empty() {
		panic("*Stack.Pop(): stack is empty")
	}
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v
}
