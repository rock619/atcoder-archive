package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(h []int, w []int) {
	out := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := out.Flush(); err != nil {
			panic(err)
		}
	}()

	row1 := make([][]int, 0, 900)
	for i := 1; i <= h[0]-2; i++ {
		for j := 1; j <= h[0]-i-1; j++ {
			k := h[0] - i - j
			if k < 1 || k > 28 {
				continue
			}
			row1 = append(row1, []int{i, j, k})
		}
	}

	row2 := make([][]int, 0, 900)
	for i := 1; i <= h[1]-2; i++ {
		for j := 1; j <= h[1]-i-1; j++ {
			k := h[1] - i - j
			if k < 1 || k > 28 {
				continue
			}
			row2 = append(row2, []int{i, j, k})
		}
	}

	row3 := make([][]int, 0, 900)
	for i := 1; i <= h[2]-2; i++ {
		for j := 1; j <= h[2]-i-1; j++ {
			k := h[2] - i - j
			if k < 1 || k > 28 {
				continue
			}
			row3 = append(row3, []int{i, j, k})
		}
	}

	count := 0
	for _, r1 := range row1 {
		for _, r2 := range row2 {
			for _, r3 := range row3 {
				if r1[0]+r2[0]+r3[0] != w[0] || r1[1]+r2[1]+r3[1] != w[1] || r1[2]+r2[2]+r3[2] != w[2] {
					continue
				}
				count++
			}
		}
	}

	fmt.Fprintln(out, count)
}

func main() {
	s := NewScanner()

	h := make([]int, 3)
	for i := 0; i < 3; i++ {
		h[i] = s.Int()
	}
	w := make([]int, 3)
	for i := 0; i < 3; i++ {
		w[i] = s.Int()
	}

	solve(h, w)
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
