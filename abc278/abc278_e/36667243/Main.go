package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type Point struct {
	H, W int
}

func solve(o Printer, H, W, N, h, w int, A [][]int) {
	counts := Make3D(H+1, W+1, N+1, 0)
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			counts[i+1][j+1][A[i][j]]++
		}
	}

	for i := 1; i <= H; i++ {
		for j := 1; j <= W; j++ {
			for k := 1; k <= N; k++ {
				counts[i][j][k] += counts[i-1][j][k]
			}
		}
	}

	for i := 1; i <= H; i++ {
		for j := 1; j <= W; j++ {
			for k := 1; k <= N; k++ {
				counts[i][j][k] += counts[i][j-1][k]
			}
		}
	}

	for k := 0; k <= H-h; k++ {
		for l := 0; l <= W-w; l++ {
			count := 0
			for m := 1; m <= N; m++ {
				if v := counts[H][W][m] - counts[k+h][l+w][m] + counts[k+h][l][m] + counts[k][l+w][m] - counts[k][l][m]; v > 0 {
					count++
				}
			}
			if l == 0 {
				o.p(count)
			} else {
				o.p(" ", count)
			}
		}
		o.l()
	}
}

type ValueType = int

func Make3D(len1, len2, len3 int, value ValueType) [][][]ValueType {
	result := make([][][]ValueType, len1)
	for i := 0; i < len(result); i++ {
		result[i] = make([][]ValueType, len2)
		for j := 0; j < len(result[i]); j++ {
			result[i][j] = make([]ValueType, len3)
			for k := 0; k < len(result[i][j]); k++ {
				result[i][j][k] = value
			}
		}
	}
	return result
}

func main() {
	sc := NewScanner()
	H, W, N, h, w := sc.Int(), sc.Int(), sc.Int(), sc.Int(), sc.Int()
	A := sc.IntNN(H, W)
	out := NewPrinter()
	solve(out, H, W, N, h, w, A)
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
