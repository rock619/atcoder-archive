package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Interval struct {
	L, R int
}

func solve(N int, L []int, R []int) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	diffs := make([]int, 200001)
	for i := range L {
		diffs[L[i]]++
		diffs[R[i]]--
	}

	current := 0
	results := make([]Interval, 0, N)
	for i := 1; i < len(diffs); i++ {
		if current == 0 && diffs[i] > 0 {
			results = append(results, Interval{L: i})
		} else if current > 0 && current+diffs[i] == 0 {
			results[len(results)-1].R = i
		}
		current += diffs[i]
	}

	// intervals := make([]Interval, N)
	// for i := range L {
	// 	intervals[i] = Interval{L: L[i], R: R[i]}
	// }

	// sort.Slice(intervals, func(i, j int) bool {
	// 	if intervals[i].R == intervals[j].R {
	// 		return intervals[i].L < intervals[j].L
	// 	}
	// 	return intervals[i].R < intervals[j].R
	// })

	// results := make([]Interval, 1, N)
	// results[0] = intervals[0]
	// for i := 1; i < N; i++ {
	// 	if intervals[i].L <= results[len(results)-1].R {
	// 		results[len(results)-1].R = intervals[i].R
	// 		if intervals[i].L < results[len(results)-1].L {
	// 			results[len(results)-1].L = intervals[i].L
	// 		}
	// 	} else {
	// 		results[len(results)-1].R = intervals[i-1].R
	// 		results = append(results, intervals[i])
	// 	}
	// }

	for _, r := range results {
		fmt.Fprintln(w, r.L, r.R)
	}
}

func main() {
	s := NewScanner()

	N := s.Int()
	L := make([]int, N)
	R := make([]int, N)
	for i := 0; i < N; i++ {
		L[i] = s.Int()
		R[i] = s.Int()
	}

	solve(N, L, R)
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

func (s Scanner) Float() float64 {
	s.Scan()
	v, err := strconv.ParseFloat(s.Text(), 64)
	if err != nil {
		panic(err)
	}
	return v
}
