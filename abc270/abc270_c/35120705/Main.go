package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type DFS struct {
	N      int
	graph  [][]int
	X, Y   int
	result []int
}

func NewDFS(N, X, Y int, graph [][]int) *DFS {
	return &DFS{
		N:      N,
		graph:  graph,
		X:      X,
		Y:      Y,
		result: nil,
	}
}

func (s *DFS) Do(next int, current []int, depth int) {
	if s.result != nil {
		return
	}
	if next == s.Y {
		s.result = append(current, next)
		return
	}

	current = append(current, next)
	for _, n := range s.graph[next] {
		if depth > 0 && n == current[depth-1] {
			continue
		}
		s.Do(n, current, depth+1)
	}
}

func solve(N int, X int, Y int, U []int, V []int) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	graph := make([][]int, N)
	for i := range U {
		u, v := U[i]-1, V[i]-1
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	x, y := X-1, Y-1
	dfs := NewDFS(N, x, y, graph)
	dfs.Do(x, []int{}, 0)
	for i, v := range dfs.result {
		if i == 0 {
			fmt.Fprint(w, v+1)
		} else {
			fmt.Fprint(w, " ", v+1)
		}
	}
	fmt.Fprintln(w)
}

type StackElement = int

type Stack struct {
	Size int
	s    []StackElement
}

func NewStack(capacity int) *Stack {
	return &Stack{
		s: make([]StackElement, 0, capacity),
	}
}

func (s *Stack) Empty() bool {
	return s.Size == 0
}

func (s *Stack) Clear() {
	s.Size = 0
}

func (s *Stack) Push(x StackElement) {
	if s.Size >= len(s.s) {
		s.s = append(s.s, x)
	} else {
		s.s[s.Size] = x
	}
	s.Size++
}

func (s *Stack) Pop() (x StackElement, ok bool) {
	if s.Empty() {
		return x, false
	}
	s.Size--
	return s.s[s.Size], true
}

func (s *Stack) Peek() (x StackElement, ok bool) {
	if s.Empty() {
		return x, false
	}
	return s.s[s.Size-1], true
}

func (s *Stack) String() string {
	if s.Empty() {
		return "Stack(0)[]"
	}
	var b strings.Builder
	fmt.Fprintf(&b, "Stack(%d)[%v", s.Size, s.s[0])
	for i := 1; i < s.Size; i++ {
		fmt.Fprintf(&b, ", %v", s.s[i])
	}
	fmt.Fprint(&b, "]<-top")
	return b.String()
}

func main() {
	sc := NewScanner()
	N := sc.Int()
	X := sc.Int()
	Y := sc.Int()
	U := make([]int, N-1)
	V := make([]int, N-1)
	for i := 0; i < N-1; i++ {
		U[i] = sc.Int()
		V[i] = sc.Int()
	}
	solve(N, X, Y, U, V)
}

type Scanner struct {
	*bufio.Scanner
}

func NewScanner() *Scanner {
	s := bufio.NewScanner(os.Stdin)
	s.Buffer(make([]byte, 4096), 1048576)
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
