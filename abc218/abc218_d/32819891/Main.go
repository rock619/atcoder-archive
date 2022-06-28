package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Point struct {
	X, Y int
}

func solve(N int, x []int, y []int) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	mx := make(map[int][]int)
	my := make(map[int][]int)
	m := make(map[Point]int)
	points := make([]Point, N)
	for i := range points {
		p := Point{X: x[i], Y: y[i]}
		points[i] = p
		mx[p.X] = append(mx[p.X], i)
		my[p.Y] = append(my[p.Y], i)
		m[p] = i
	}
	result := 0
	for i, p := range points {
		if len(mx[p.X]) == 1 || len(my[p.Y]) == 1 {
			continue
		}
		index := sort.Search(len(mx[p.X]), func(j int) bool {
			return mx[p.X][j] > i
		})
		if index >= len(mx[p.X]) {
			continue
		}
		xMatched := mx[p.X][index:]
		index = sort.Search(len(my[p.Y]), func(j int) bool {
			return my[p.Y][j] > i
		})
		if index >= len(my[p.Y]) {
			continue
		}
		yMatched := my[p.Y][index:]
		for _, j := range xMatched {
			for _, k := range yMatched {
				oppositeX, oppositeY := points[k].X, points[j].Y
				if opposite, ok := m[Point{X: oppositeX, Y: oppositeY}]; ok && i < opposite {
					result++
				}
			}
		}
	}
	fmt.Fprintln(w, result)
}

func main() {
	s := NewScanner()

	N := s.Int()
	x := make([]int, N)
	y := make([]int, N)
	for i := 0; i < N; i++ {
		x[i] = s.Int()
		y[i] = s.Int()
	}

	solve(N, x, y)
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
