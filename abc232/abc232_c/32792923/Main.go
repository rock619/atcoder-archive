package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

const yes = "Yes"
const no = "No"

func solve(N int, M int, A []int, B []int, C []int, D []int) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	graphA := make([][]int, N)
	for i := 0; i < M; i++ {
		ai, bi := A[i]-1, B[i]-1
		graphA[ai] = append(graphA[ai], bi)
		graphA[bi] = append(graphA[bi], ai)
	}

	indexes := make([]int, N)
	for i := range indexes {
		indexes[i] = i
	}
	p := NewPermutation(indexes)
	for {
		graphB := make([][]int, N)
		idxes := p.Current()
		for i := 0; i < M; i++ {
			ci, di := idxes[C[i]-1], idxes[D[i]-1]
			graphB[ci] = append(graphB[ci], di)
			graphB[di] = append(graphB[di], ci)
		}

		if EqualGraph(graphA, graphB) {
			fmt.Fprintln(w, yes)
			return
		}

		if !p.Next() {
			break
		}
	}
	fmt.Fprintln(w, no)
}

type Permutation struct {
	indexes []int
	s       []int
}

func NewPermutation(s []int) *Permutation {
	indexes := make([]int, len(s))
	for i := range indexes {
		indexes[i] = i
	}
	return &Permutation{
		indexes: indexes,
		s:       s,
	}
}

func (p *Permutation) Current() []int {
	result := make([]int, len(p.s))
	for i, n := range p.indexes {
		result[i] = p.s[n]
	}
	return result
}

func (p *Permutation) Next() bool {
	left := len(p.indexes) - 2
	for left >= 0 && p.indexes[left] >= p.indexes[left+1] {
		left--
	}
	if left < 0 {
		return false
	}

	right := len(p.indexes) - 1
	for p.indexes[left] >= p.indexes[right] {
		right--
	}

	p.indexes[left], p.indexes[right] = p.indexes[right], p.indexes[left]
	left++
	right = len(p.indexes) - 1
	for left < right {
		p.indexes[left], p.indexes[right] = p.indexes[right], p.indexes[left]
		left++
		right--
	}

	return true
}

func EqualGraph(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		sort.Ints(a[i])
		sort.Ints(b[i])
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}

func main() {
	s := NewScanner()

	N := s.Int()
	M := s.Int()
	A := make([]int, M)
	B := make([]int, M)
	for i := 0; i < M; i++ {
		A[i] = s.Int()
		B[i] = s.Int()
	}
	C := make([]int, M)
	D := make([]int, M)
	for i := 0; i < M; i++ {
		C[i] = s.Int()
		D[i] = s.Int()
	}

	solve(N, M, A, B, C, D)
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
