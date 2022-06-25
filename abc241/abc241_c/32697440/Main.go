package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const yes = "Yes"
const no = "No"

func solve(N int, S [][]byte) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	// 横
	for i := 0; i < N; i++ {
		for j := 0; j+5 < N; j++ {
			white := 0
			for k := j; k <= j+5; k++ {
				if S[i][k] == '.' {
					white++
					if white > 2 {
						break
					}
				}
			}
			if white <= 2 {
				fmt.Fprintln(w, yes)
				return
			}
		}
	}

	// 縦
	for i := 0; i < N; i++ {
		for j := 0; j+5 < N; j++ {
			white := 0
			for k := j; k <= j+5; k++ {
				if S[k][i] == '.' {
					white++
					if white > 2 {
						break
					}
				}
			}
			if white <= 2 {
				fmt.Fprintln(w, yes)
				return
			}
		}
	}

	// 斜め（右下方向）
	for i := 0; i+5 < N; i++ {
		for j := 0; j+5 < N; j++ {
			white := 0
			for k := 0; k <= 5; k++ {
				if S[i+k][j+k] == '.' {
					white++
					if white > 2 {
						break
					}
				}
			}
			if white <= 2 {
				fmt.Fprintln(w, yes)
				return
			}
		}
	}

	// 斜め（左下方向）
	for i := 0; i+5 < N; i++ {
		for j := 5; j < N; j++ {
			white := 0
			for k := 0; k <= 5; k++ {
				if S[i+k][j-k] == '.' {
					white++
					if white > 2 {
						break
					}
				}
			}
			if white <= 2 {
				fmt.Fprintln(w, yes)
				return
			}
		}
	}
	fmt.Fprintln(w, no)
}

func main() {
	s := NewScanner()
	N := s.Int()
	S := s.BytesN(N)

	solve(N, S)
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
	s.Scan()
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
	return s.Scanner.Bytes()
}

func (s *Scanner) BytesN(n int) [][]byte {
	v := make([][]byte, n)
	for i := 0; i < n; i++ {
		b := s.Bytes()
		v[i] = make([]byte, len(b))
		copy(v[i], b)
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
