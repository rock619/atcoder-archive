package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type DFS struct {
	S      []string
	m      map[string]struct{}
	used   []bool
	N      int
	result string
}

func NewDFS(N, M int, S, T []string) *DFS {
	m := make(map[string]struct{}, M)
	for _, v := range T {
		m[v] = struct{}{}
	}

	return &DFS{
		S:      S,
		m:      m,
		used:   make([]bool, N),
		N:      N,
		result: "",
	}
}

func (s *DFS) Do(current string, depth int) {
	if s.result != "" {
		return
	}
	if depth == s.N {
		if _, ok := s.m[current]; !ok && len(current) >= 3 && len(current) <= 16 {
			s.result = current
		}
		return
	}
	if depth == 0 {
		for j, v := range s.S {
			s.used[j] = true
			s.Do(v, depth+1)
			s.used[j] = false
		}
	} else {
		for i := 0; len(current)+i < 16; i++ {
			uss := strings.Repeat("_", i+1)
			for j, v := range s.S {
				if s.used[j] {
					continue
				}
				if len(current)+i+1+len(v) > 16 {
					continue
				}

				s.used[j] = true
				s.Do(current+uss+v, depth+1)
				s.used[j] = false
			}
		}
	}

}

func solve(N int, M int, S []string, T []string) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	dfs := NewDFS(N, M, S, T)
	dfs.Do("", 0)
	if dfs.result == "" {
		fmt.Fprintln(w, -1)
	} else {
		fmt.Fprintln(w, dfs.result)
	}
}

func main() {
	sc := NewScanner()
	N := sc.Int()
	M := sc.Int()
	S := make([]string, N)
	for i := 0; i < N; i++ {
		S[i] = sc.Text()
	}
	T := make([]string, M)
	for i := 0; i < M; i++ {
		T[i] = sc.Text()
	}
	solve(N, M, S, T)
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
