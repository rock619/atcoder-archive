package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type DFS struct {
	H, W     int
	C        [][]byte
	maxDepth int
	used     [][]bool
}

func NewDFS(H, W int, C [][]byte) *DFS {
	used := make([][]bool, H)
	for i := range used {
		used[i] = make([]bool, W)
	}
	return &DFS{
		H:        H,
		W:        W,
		C:        C,
		maxDepth: 1,
		used:     used,
	}
}

func (s *DFS) Do(h, w, depth int) {
	s.used[h][w] = true
	UpdateMax(&s.maxDepth, depth)

	for _, pair := range [2][2]int{{h + 1, w}, {h, w + 1}} {
		nextH, nextW := pair[0], pair[1]
		if nextH >= s.H || nextW >= s.W || s.C[nextH][nextW] == '#' || s.used[nextH][nextW] {
			continue
		}
		s.Do(nextH, nextW, depth+1)
	}
}

func UpdateMax(max *int, v int) {
	if v > *max {
		*max = v
	}
}

func solve(H, W int, C [][]byte) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	dfs := NewDFS(H, W, C)
	dfs.Do(0, 0, 1)
	fmt.Fprintln(w, dfs.maxDepth)
}

func main() {
	s := NewScanner()

	H := s.Int()
	W := s.Int()
	C := s.BytesN(H)

	solve(H, W, C)
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
