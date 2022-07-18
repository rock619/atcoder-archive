package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Student struct {
	ID, A, B, Sum int
	Pass          bool
}

func solve(N int, X int, Y int, Z int, A []int, B []int) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	students := make([]Student, N)
	for i := range A {
		students[i] = Student{
			ID:  i + 1,
			A:   A[i],
			B:   B[i],
			Sum: A[i] + B[i],
		}
	}
	sort.Slice(students, func(i, j int) bool {
		if students[i].A == students[j].A {
			return students[i].ID < students[j].ID
		}
		return students[i].A > students[j].A
	})
	count := 0
	for i := range students {
		if count == X {
			break
		}
		students[i].Pass = true
		count++
	}

	sort.Slice(students, func(i, j int) bool {
		if students[i].B == students[j].B {
			return students[i].ID < students[j].ID
		}
		return students[i].B > students[j].B
	})

	count = 0
	for i := range students {
		if count == Y {
			break
		}
		if students[i].Pass {
			continue
		}
		students[i].Pass = true
		count++
	}

	sort.Slice(students, func(i, j int) bool {
		if students[i].Sum == students[j].Sum {
			return students[i].ID < students[j].ID
		}
		return students[i].Sum > students[j].Sum
	})

	count = 0
	for i := range students {
		if count == Z {
			break
		}
		if students[i].Pass {
			continue
		}
		students[i].Pass = true
		count++
	}

	sort.Slice(students, func(i, j int) bool {
		return students[i].ID < students[j].ID
	})

	for _, s := range students {
		if s.Pass {
			fmt.Fprintln(w, s.ID)
		}
	}
}

func main() {
	s := NewScanner()
	N := s.Int()
	X := s.Int()
	Y := s.Int()
	Z := s.Int()
	A := make([]int, N)
	for i := 0; i < N; i++ {
		A[i] = s.Int()
	}
	B := make([]int, N)
	for i := 0; i < N; i++ {
		B[i] = s.Int()
	}
	solve(N, X, Y, Z, A, B)
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
