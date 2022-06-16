package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type DFS struct {
	cards   []int
	graph   [][]int
	chosen  []int
	current []int
	done    bool
	result  [][][]int
}

func NewDFS(cards []int, graph [][]int) *DFS {
	return &DFS{
		cards:   cards,
		graph:   graph,
		chosen:  make([]int, len(cards)),
		current: make([]int, 0, len(cards)),
		result:  make([][][]int, 8888),
	}
}

func (s *DFS) Do(position, depth int) {
	if s.done {
		return
	}
	if position == len(s.graph) {
		r := make([]int, len(s.current))
		copy(r, s.current)
		s.result[depth] = append(s.result[depth], r)
		if len(s.result[depth]) == 2 {
			s.done = true
		}
		return
	}

	s.Do(position+1, depth)

	if s.chosen[position] == 0 {
		s.current = append(s.current, position)
		for _, next := range s.graph[position] {
			s.chosen[next]++
		}
		s.Do(position+1, depth+s.cards[position])
		for _, next := range s.graph[position] {
			s.chosen[next]--
		}
		s.current = s.current[:len(s.current)-1]
	}
}

func solve(N, Q int, A, X, Y []int) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	graph := make([][]int, N)
	for i := range X {
		xi, yi := X[i]-1, Y[i]-1
		graph[xi] = append(graph[xi], yi)
	}

	dfs := NewDFS(A, graph)
	dfs.Do(0, 0)

	for i := 0; i < 8888; i++ {
		if len(dfs.result[i]) <= 1 {
			continue
		}
		for j := 0; j <= 1; j++ {
			fmt.Fprintln(w, len(dfs.result[i][j]))
			for k, v := range dfs.result[i][j] {
				if k != 0 {
					fmt.Fprint(w, " ")
				}
				fmt.Fprint(w, v+1)
			}
			fmt.Fprintln(w)
		}
	}
}

func main() {
	s := NewScanner()
	N := s.Int()
	Q := s.Int()
	A := s.Ints(N)
	X, Y := s.Ints2(Q)

	solve(N, Q, A, X, Y)
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
