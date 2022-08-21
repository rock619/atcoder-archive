package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

const yes = "Yes"
const no = "No"

type Result struct {
	X, Y, Z, W int
}

func solve(N int, P int, Q int, R int, A []int) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	sumPQR := P + Q + R
	// log.Println(sumPQR)

	sums := make([]int, N+1)
	for i := 1; i <= N; i++ {
		sums[i] = sums[i-1] + A[i-1]
	}
	// log.Println(sums)

	if sums[N] < sumPQR {
		fmt.Fprintln(w, no)
		return
	}

	results := make([]Result, 0)
	minWI := sort.Search(len(sums), func(i int) bool {
		return sums[i] >= sumPQR
	})
	// log.Println(minWI)
	for i := minWI; i < len(sums); i++ {
		xi := sort.Search(i, func(j int) bool {
			return sums[i]-sums[j] < sumPQR
		})
		xi--
		if sums[i]-sums[xi] == sumPQR {
			results = append(results, Result{X: xi, W: i})
		}
		// log.Println(i, xi)
	}

	for i, r := range results {
		yi := sort.Search(r.W, func(j int) bool {
			return sums[j]-sums[r.X] > P
		})
		yi--
		if sums[yi]-sums[r.X] == P {
			results[i].Y = yi
		}
	}

	for _, r := range results {
		if r.Y == 0 {
			continue
		}

		zi := sort.Search(r.W, func(j int) bool {
			return sums[j]-sums[r.Y] > Q
		})
		zi--
		if sums[zi]-sums[r.Y] == Q && sums[r.W]-sums[zi] == R {
			// log.Println(r.X, r.Y, zi, r.W)
			fmt.Fprintln(w, yes)
			return
		}
	}
	fmt.Fprintln(w, no)
}

func main() {
	sc := NewScanner()
	N := sc.Int()
	P := sc.Int()
	Q := sc.Int()
	R := sc.Int()
	A := make([]int, N)
	for i := 0; i < N; i++ {
		A[i] = sc.Int()
	}
	solve(N, P, Q, R, A)
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
