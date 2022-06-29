package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func solve(N int, A []int, B []int) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	graph := make([][]int, N)
	for i := range A {
		ai, bi := A[i]-1, B[i]-1
		graph[ai] = append(graph[ai], bi)
		graph[bi] = append(graph[bi], ai)
	}

	for i := range graph {
		sort.Sort(sort.Reverse(sort.IntSlice(graph[i])))
		// sort.Ints(graph[i])
	}

	prevs := make([]int, N)
	for i := range prevs {
		prevs[i] = -1
	}

	result := make([]int, 0, N)
	stack := Stack(make([]int, 0, N))
	stack.Push(0)
	for i := 0; !stack.Empty(); {
		prev := i
		i = stack.Pop()
		result = append(result, i+1)

		if prevs[i] != -1 {
			continue
		}
		prevs[i] = prev

		for _, next := range graph[i] {
			if next > 0 && prevs[next] == -1 {
				stack.Push(i)
				stack.Push(next)
			}
		}
	}

	for i, r := range result {
		if i == 0 {
			fmt.Fprint(w, r)
		} else {
			fmt.Fprint(w, " ", r)
		}
	}
	fmt.Fprintln(w)
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

func UpdateMin(min *int, v int) {
	if v < *min {
		*min = v
	}
}

func main() {
	s := NewScanner()

	N := s.Int()
	A := make([]int, N-1)
	B := make([]int, N-1)
	for i := 0; i < N-1; i++ {
		A[i] = s.Int()
		B[i] = s.Int()
	}

	solve(N, A, B)
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

func (s *Scanner) Int() int {
	if ok := s.Scan(); !ok {
		panic(s.Err())
	}
	v, err := strconv.Atoi(s.Text())
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
	if ok := s.Scan(); !ok {
		panic(s.Err())
	}

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

func (s Scanner) Float() float64 {
	s.Scan()
	v, err := strconv.ParseFloat(s.Text(), 64)
	if err != nil {
		panic(err)
	}
	return v
}
