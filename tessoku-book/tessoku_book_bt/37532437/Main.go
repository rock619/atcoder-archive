package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func UpdateMax(max *int, v ...int) {
	for i := 0; i < len(v); i++ {
		if v[i] > *max {
			*max = v[i]
		}
	}
}

func solve(o Printer, H, W, K int, c [][]byte) {
	originCount := 0
	hCounts := make([]int, H+1)
	for i := 1; i <= H; i++ {
		for j := 1; j <= W; j++ {
			if c[i-1][j-1] == '.' {
				hCounts[i]++
			} else {
				originCount++
			}
		}
	}

	max := 0
	for i := 0; i < (1 << H); i++ {
		usedCount := 0
		hUsed := make([]bool, H+1)
		count := 0
		for j := 1; j <= H; j++ {
			if i&(1<<(j-1)) != 0 {
				hUsed[j] = true
				usedCount++
				count += hCounts[j]
			}
		}

		rest := K - usedCount
		if rest < 0 {
			continue
		}

		wCounts := make([]int, W+1)
		for j := 1; j <= H; j++ {
			if hUsed[j] {
				continue
			}
			for k := 1; k <= W; k++ {
				if c[j-1][k-1] == '.' {
					wCounts[k]++
				}
			}
		}

		sort.Sort(sort.Reverse(sort.IntSlice(wCounts)))
		for j := 0; j < rest; j++ {
			count += wCounts[j]
		}
		UpdateMax(&max, count)
	}
	o.l(originCount + max)
}

func main() {
	sc := NewScanner()
	H, W, K := sc.Int(), sc.Int(), sc.Int()
	c := make([][]byte, H)
	for i := range c {
		c[i] = sc.Bytes()
	}
	out := NewPrinter()
	solve(out, H, W, K, c)
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
