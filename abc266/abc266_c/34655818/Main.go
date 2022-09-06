package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const yes = "Yes"
const no = "No"

func solve(A_x int, A_y int, B_x int, B_y int, C_x int, C_y int, D_x int, D_y int) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	s := (A_x-C_x)*(B_y-A_y) - (A_y-C_y)*(B_x-A_x)
	t := (A_x-C_x)*(D_y-A_y) - (A_y-C_y)*(D_x-A_x)
	if s*t > 0 {
		fmt.Fprintln(w, no)
		return
	}
	s = (B_x-D_x)*(A_y-B_y) - (B_y-D_y)*(A_x-B_x)
	t = (B_x-D_x)*(C_y-B_y) - (B_y-D_y)*(C_x-B_x)
	if s*t > 0 {
		fmt.Fprintln(w, no)
		return
	}
	fmt.Fprintln(w, yes)
}

func main() {
	sc := NewScanner()
	A_x := sc.Int()
	A_y := sc.Int()
	B_x := sc.Int()
	B_y := sc.Int()
	C_x := sc.Int()
	C_y := sc.Int()
	D_x := sc.Int()
	D_y := sc.Int()
	solve(A_x, A_y, B_x, B_y, C_x, C_y, D_x, D_y)
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
