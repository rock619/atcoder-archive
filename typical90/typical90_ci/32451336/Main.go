package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const inf = "Infinity"

type Edge struct {
	From, To, Cost int
}

func solve(N, P, K int, A [][]int) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	graph := make([][]Edge, N)
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			graph[i] = append(graph[i], Edge{From: i, To: j, Cost: A[i][j]})
		}
	}

	l := 0
	for left, right, center := 0, P+2, P/2+1; left < right; center = (left + right) / 2 {
		dists := floydWarshall(graph, center)
		if countPairsLTE(dists, P) <= K {
			if right == center {
				l = center
				break
			}
			right = center
		} else {
			if left == center {
				l = center
				break
			}
			left = center
		}
	}

	r := 0
	for left, right, center := 0, P+2, P/2+1; left < right; center = (left + right) / 2 {
		dists := floydWarshall(graph, center)
		if countPairsLTE(dists, P) < K {
			if right == center {
				r = center
				break
			}
			right = center
		} else {
			if left == center {
				r = center
				break
			}
			left = center
		}
	}

	switch lInf, rInf := l > P, r > P; {
	case lInf && rInf:
		fmt.Fprintln(w, 0)
	case rInf:
		fmt.Fprintln(w, inf)
	default:
		fmt.Fprintln(w, r-l)
	}
}

func floydWarshall(graph [][]Edge, x int) [][]int {
	dists := make([][]int, len(graph))
	for i := range graph {
		dists[i] = make([]int, len(graph[i]))
		for j := range graph[i] {
			dists[i][j] = graph[i][j].Cost
			if dists[i][j] == -1 {
				dists[i][j] = x
			}
		}
	}

	for k := range graph {
		for i := range graph {
			for j := range graph {
				UpdateMin(&dists[i][j], dists[i][k]+dists[k][j])
			}
		}
	}
	return dists
}

func countPairsLTE(dists [][]int, P int) int {
	count := 0
	for i := range dists {
		for j := i + 1; j < len(dists); j++ {
			if dists[i][j] <= P {
				count++
			}
		}
	}
	return count
}

func UpdateMin(min *int, v int) {
	if v < *min {
		*min = v
	}
}

func main() {
	s := NewScanner()
	N := s.Int()
	P := s.Int()
	K := s.Int()
	A := s.IntsN(N, N)

	solve(N, P, K, A)
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
