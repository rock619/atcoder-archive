package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const yes = "Yes"
const no = "No"

func solve(o Printer, N int, Q int, S []byte, L []int, R []int) {
	rollingHash := NewRollingHash(N)
	hash := rollingHash.Of(string(S))
	hashR := rollingHash.Of(string(Reversed(S)))
	for i := 0; i < Q; i++ {
		if rollingHash.Slice(hash, L[i]-1, R[i]) == rollingHash.Slice(hashR, N-R[i], N-L[i]+1) {
			o.l(yes)
		} else {
			o.l(no)
		}
	}
}

func Reversed(s []byte) []byte {
	r := make([]byte, len(s))
	for left, right := 0, len(s)-1; left <= right; left, right = left+1, right-1 {
		r[left], r[right] = s[right], s[left]
	}
	return r
}

const (
	rollingHashMask30      uint64 = (1 << 30) - 1
	rollingHashMask31      uint64 = (1 << 31) - 1
	rollingHashMod         uint64 = (1 << 61) - 1
	rollingHashPositivizer uint64 = rollingHashMod * ((1 << 3) - 1)
)

type RollingHash struct {
	base uint64
	pows []uint64
}

func NewRollingHash(maxLen int) *RollingHash {
	rand.Seed(time.Now().UnixNano())
	minBase := 129
	h := &RollingHash{
		base: uint64(rand.Intn(math.MaxInt32-minBase) + minBase),
		pows: make([]uint64, maxLen+1),
	}

	h.pows[0] = 1
	for i := 1; i < len(h.pows); i++ {
		h.pows[i] = h.CalcMod(h.Mul(h.pows[i-1], h.base))
	}
	return h
}

func (h *RollingHash) Of(s string) []uint64 {
	hash := make([]uint64, len(s)+1)
	for i := 0; i < len(s); i++ {
		hash[i+1] = h.CalcMod(h.Mul(hash[i], h.base) + uint64(s[i]))
	}
	return hash
}

func (h *RollingHash) Slice(hash []uint64, low, high int) uint64 {
	return h.CalcMod(hash[high] + rollingHashPositivizer - h.Mul(hash[low], h.pows[high-low]))
}

func (h *RollingHash) Mul(l, r uint64) uint64 {
	lu := l >> 31
	ld := l & rollingHashMask31
	ru := r >> 31
	rd := r & rollingHashMask31
	middleBit := ld*ru + lu*rd
	return ((lu * ru) << 1) + ld*rd + ((middleBit & rollingHashMask30) << 31) + (middleBit >> 30)
}

func (h *RollingHash) CalcMod(v uint64) uint64 {
	result := (v & rollingHashMod) + (v >> 61)
	if result >= rollingHashMod {
		return result - rollingHashMod
	}
	return result
}

func main() {
	sc := NewScanner()
	N := sc.Int()
	Q := sc.Int()
	S := sc.Bytes()
	L := make([]int, Q)
	R := make([]int, Q)
	for i := 0; i < Q; i++ {
		L[i] = sc.Int()
		R[i] = sc.Int()
	}
	out := NewPrinter()
	solve(out, N, Q, S, L, R)
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
