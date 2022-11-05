package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const yes = "Yes"
const no = "No"

type Point struct {
	H, W int
}

func solve(o Printer, H, W int, C [][]byte) {
	start := Point{}
	for i := range C {
		for j := range C[i] {
			if C[i][j] == 'S' {
				start.H, start.W = i, j
			}
		}
	}
	if start.H > 0 && C[start.H-1][start.W] == '.' {
		stack := NewStack(4)
		stack.Push(Point{start.H - 1, start.W})
		for !stack.Empty() {
			p, _ := stack.Pop()
			C[p.H][p.W] = 'A'
			for _, d := range [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
				h, w := p.H+d[0], p.W+d[1]
				if h >= 0 && h < H && w >= 0 && w < W && C[h][w] == '.' {
					stack.Push(Point{h, w})
				}
			}
		}
	}
	if start.H < H-1 {
		if C[start.H+1][start.W] == 'A' {
			o.l("Yes")
			return
		}
		if C[start.H+1][start.W] == '.' {
			stack := NewStack(4)
			stack.Push(Point{start.H + 1, start.W})
			for !stack.Empty() {
				p, _ := stack.Pop()
				C[p.H][p.W] = 'B'
				for _, d := range [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
					h, w := p.H+d[0], p.W+d[1]
					if h >= 0 && h < H && w >= 0 && w < W {
						switch C[h][w] {
						case 'A':
							o.l("Yes")
							return
						case '.':
							stack.Push(Point{h, w})
						}
					}
				}
			}
		}
	}
	if start.W > 0 {
		if C[start.H][start.W-1] == 'A' || C[start.H][start.W-1] == 'B' {
			o.l("Yes")
			return
		}
		if C[start.H][start.W-1] == '.' {
			stack := NewStack(4)
			stack.Push(Point{start.H, start.W - 1})
			for !stack.Empty() {
				p, _ := stack.Pop()
				switch C[p.H][p.W] {
				case 'A', 'B':
					o.l("Yes")
					return
				}
				C[p.H][p.W] = 'C'
				for _, d := range [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
					h, w := p.H+d[0], p.W+d[1]
					if h >= 0 && h < H && w >= 0 && w < W {
						switch C[h][w] {
						case 'A', 'B':
							o.l("Yes")
							return
						case '.':
							stack.Push(Point{h, w})
						}
					}
				}
			}
		}
	}

	if start.W < W-1 {
		if C[start.H][start.W+1] == 'A' || C[start.H][start.W+1] == 'B' || C[start.H][start.W+1] == 'C' {
			o.l("Yes")
			return
		}
		if C[start.H][start.W+1] == '.' {
			stack := NewStack(4)
			stack.Push(Point{start.H, start.W + 1})
			for !stack.Empty() {
				p, _ := stack.Pop()
				switch C[p.H][p.W] {
				case 'A', 'B', 'C':
					o.l("Yes")
					return
				}
				C[p.H][p.W] = 'D'

				for _, d := range [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
					h, w := p.H+d[0], p.W+d[1]
					if h >= 0 && h < H && w >= 0 && w < W {
						switch C[h][w] {
						case 'A', 'B', 'C':
							o.l("Yes")
							return
						case '.':
							stack.Push(Point{h, w})
						}
					}
				}
			}
		}
	}

	o.l("No")
}

type StackElement = Point

type Stack struct {
	Size int
	s    []StackElement
}

func NewStack(capacity int) *Stack {
	return &Stack{
		s: make([]StackElement, 0, capacity),
	}
}

func (s *Stack) Empty() bool {
	return s.Size == 0
}

func (s *Stack) Clear() {
	s.Size = 0
}

func (s *Stack) Push(x StackElement) {
	if s.Size >= len(s.s) {
		s.s = append(s.s, x)
	} else {
		s.s[s.Size] = x
	}
	s.Size++
}

func (s *Stack) Pop() (x StackElement, ok bool) {
	if s.Empty() {
		return x, false
	}
	s.Size--
	return s.s[s.Size], true
}

func (s *Stack) Peek() (x StackElement, ok bool) {
	if s.Empty() {
		return x, false
	}
	return s.s[s.Size-1], true
}

func (s *Stack) String() string {
	if s.Empty() {
		return "Stack(0)[]"
	}
	var b strings.Builder
	fmt.Fprintf(&b, "Stack(%d)[%v", s.Size, s.s[0])
	for i := 1; i < s.Size; i++ {
		fmt.Fprintf(&b, ", %v", s.s[i])
	}
	fmt.Fprint(&b, "]<-top")
	return b.String()
}

func main() {
	sc := NewScanner()

	H, W := sc.Int(), sc.Int()
	C := make([][]byte, H)
	for i := range C {
		C[i] = sc.Bytes()
	}
	out := NewPrinter()
	solve(out, H, W, C)
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
