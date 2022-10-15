package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func solve(o Printer, H, W, rs, cs, N int, r, c []int, Q int, d []byte, l []int) {
	wallsByRow := make(map[int][]int)
	wallsByCol := make(map[int][]int)
	for i := 0; i < N; i++ {
		ri, ci := r[i], c[i]
		if len(wallsByRow[ri]) == 0 {
			wallsByRow[ri] = []int{ci}
		} else {
			wallsByRow[ri] = append(wallsByRow[ri], ci)
		}
		if len(wallsByCol[ci]) == 0 {
			wallsByCol[ci] = []int{ri}
		} else {
			wallsByCol[ci] = append(wallsByCol[ci], ri)
		}
	}
	for k := range wallsByRow {
		sort.Ints(wallsByRow[k])
	}
	for k := range wallsByCol {
		sort.Ints(wallsByCol[k])
	}

	row, col := rs, cs
	for i := 0; i < Q; i++ {
		di, li := d[i], l[i]
		switch di {
		case 'L':
			min := col - li
			if min < 1 {
				min = 1
			}
			index := sort.Search(len(wallsByRow[row]), func(j int) bool {
				return wallsByRow[row][j] >= col
			})
			index--
			if index < 0 || wallsByRow[row][index] < min {
				col = min
				break
			}
			col = wallsByRow[row][index] + 1
		case 'R':
			max := col + li
			if max > W {
				max = W
			}
			index := sort.Search(len(wallsByRow[row]), func(j int) bool {
				return wallsByRow[row][j] > col
			})
			if index >= len(wallsByRow[row]) || wallsByRow[row][index] > max {
				col = max
				break
			}
			col = wallsByRow[row][index] - 1
		case 'U':
			min := row - li
			if min < 1 {
				min = 1
			}
			index := sort.Search(len(wallsByCol[col]), func(j int) bool {
				return wallsByCol[col][j] >= row
			})
			index--
			if index < 0 || wallsByCol[col][index] < min {
				row = min
				break
			}
			row = wallsByCol[col][index] + 1
		case 'D':
			max := row + li
			if max > H {
				max = H
			}
			index := sort.Search(len(wallsByCol[col]), func(j int) bool {
				return wallsByCol[col][j] > row
			})
			if index >= len(wallsByCol[col]) || wallsByCol[col][index] > max {
				row = max
				break
			}
			row = wallsByCol[col][index] - 1
		}
		o.l(row, col)
	}
}

func main() {
	sc := NewScanner()
	H, W, rs, cs := sc.Int(), sc.Int(), sc.Int(), sc.Int()
	N := sc.Int()
	r, c := sc.IntN2(N)
	Q := sc.Int()
	d := make([]byte, Q)
	l := make([]int, Q)
	for i := 0; i < Q; i++ {
		d[i] = sc.Bytes()[0]
		l[i] = sc.Int()
	}
	out := NewPrinter()
	solve(out, H, W, rs, cs, N, r, c, Q, d, l)
	if err := out.w.Flush(); err != nil {
		panic(err)
	}
}

type Scanner struct {
	*bufio.Scanner
}

func NewScanner() *Scanner {
	s := bufio.NewScanner(os.Stdin)
	s.Buffer(make([]byte, 4096), math.MaxInt64)
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

type Printer interface {
	// p fmt.Print
	p(a ...interface{})
	// f fmt.Printf
	f(format string, a ...interface{})
	// l fmt.Println
	l(a ...interface{})
}

type printer struct {
	w *bufio.Writer
}

func NewPrinter() *printer {
	return &printer{bufio.NewWriter(os.Stdout)}
}

func (p *printer) p(a ...interface{}) {
	fmt.Fprint(p.w, a...)
}

func (p *printer) f(format string, a ...interface{}) {
	fmt.Fprintf(p.w, format, a...)
}

func (p *printer) l(a ...interface{}) {
	fmt.Fprintln(p.w, a...)
}
