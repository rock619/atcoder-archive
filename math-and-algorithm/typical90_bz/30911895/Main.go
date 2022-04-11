package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Stack struct {
	s []int64
}

func NewStack(cap int64) *Stack {
	return &Stack{
		s: make([]int64, 0, cap),
	}
}

func (s Stack) Size() int64 {
	return int64(len(s.s))
}

func (s *Stack) Push(e int64) {
	s.s = append(s.s, e)
}

func (s *Stack) Pop() int64 {
	if s.Size() == 0 {
		panic("Stack.Pop(): stack is empty")
	}
	e := s.s[len(s.s)-1]
	s.s = s.s[:len(s.s)-1]
	return e
}

func solve(N int64, M int64, a []int64, b []int64) {
	graph := make([]int64, N)
	for i := range a {
		if a[i]-1 < b[i]-1 {
			graph[b[i]-1]++
		} else {
			graph[a[i]-1]++
		}
	}

	count := 0
	for _, v := range graph {
		if v == 1 {
			count++
		}
	}
	fmt.Println(count)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var N int64
	scanner.Scan()
	N, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	var M int64
	scanner.Scan()
	M, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	a := make([]int64, M)
	b := make([]int64, M)
	for i := int64(0); i < M; i++ {
		scanner.Scan()
		a[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
		scanner.Scan()
		b[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	solve(N, M, a, b)
}
