package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func Digits(n int) []int {
	result := make([]int, 0, 4)
	for i, x := 0, n; i < 4; i, x = i+1, x/10 {
		result = append(result, x%10)
	}
	return result
}

func ToInt(digits []int) int {
	result := 0
	for i := len(digits) - 1; i >= 0; i-- {
		result *= 10
		result += digits[i]
	}
	return result
}

func SecondPrizes(n int) []int {
	result := make([]int, 0, 36)
	digits := Digits(n)
	for i := 0; i < len(digits); i++ {
		for j := 0; j <= 9; j++ {
			if digits[i] == j {
				continue
			}
			ds := make([]int, len(digits))
			copy(ds, digits)
			ds[i] = j
			result = append(result, ToInt(ds))
		}
	}
	return result
}

func Contains(s []int, n int) bool {
	for _, v := range s {
		if v == n {
			return true
		}
	}
	return false
}

func solve(o Printer, N int, S []int, T []int) {
	result := make([]int, 0, 1)
	for i := 0; i <= 9999; i++ {
		ok := true
		secondPrizes := SecondPrizes(i)
		for j := range S {
			s, t := S[j], T[j]
			switch t {
			case 1:
				ok = s == i
			case 2:
				ok = s != i && Contains(secondPrizes, s)
			default:
				ok = !(s == i || Contains(secondPrizes, s))
			}
			if !ok {
				break
			}
		}
		if ok {
			result = append(result, i)
		}
	}

	if len(result) >= 2 {
		o.l("Can't Solve")
	} else {
		o.f("%04d\n", result[0])
	}
}

func main() {
	sc := NewScanner()
	N := sc.Int()
	S := make([]int, N)
	T := make([]int, N)
	for i := 0; i < N; i++ {
		S[i] = sc.Int()
		T[i] = sc.Int()
	}
	out := NewPrinter()
	solve(out, N, S, T)
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
