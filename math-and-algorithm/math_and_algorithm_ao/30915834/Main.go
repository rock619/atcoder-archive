package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	YES = "Yes"
	NO  = "No"
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

func solve(N int64, M int64, A []int64, B []int64) {
	graph := make([][]int64, N)
	for i, a := range A {
		graph[a-1] = append(graph[a-1], B[i]-1)
		graph[B[i]-1] = append(graph[B[i]-1], a-1)
	}

	results := make([]bool, N)
	visited := make([]bool, N)
	for {
		start := notVisited(visited)
		if start == -1 {
			fmt.Println(YES)
			return
		}

		visited[start] = true
		stack := NewStack(1)
		stack.Push(start)
		for stack.Size() > 0 {
			pos := stack.Pop()

			for _, next := range graph[pos] {
				if visited[next] {
					if results[next] == results[pos] {
						fmt.Println(NO)
						return
					}
					continue
				}

				results[next] = !results[pos]
				visited[next] = true
				stack.Push(next)
			}
		}
	}
}

func notVisited(visited []bool) int64 {
	for i, v := range visited {
		if !v {
			return int64(i)
		}
	}
	return -1
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
	A := make([]int64, M)
	B := make([]int64, M)
	for i := int64(0); i < M; i++ {
		scanner.Scan()
		A[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
		scanner.Scan()
		B[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	solve(N, M, A, B)
}
