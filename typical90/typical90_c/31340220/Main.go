package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Stack struct {
	s []int
}

func NewStack(cap int) *Stack {
	return &Stack{
		s: make([]int, 0, cap),
	}
}

func (s Stack) Empty() bool {
	return s.Size() == 0
}

func (s Stack) Size() int {
	return len(s.s)
}

func (s *Stack) Push(item int) {
	s.s = append(s.s, item)
}

func (s *Stack) Pop() int {
	if s.Size() == 0 {
		panic("Stack.Pop(): stack is empty")
	}
	item := s.s[len(s.s)-1]
	s.s = s.s[:len(s.s)-1]
	return item
}

type Graph [][]Edge

type Edge struct {
	To int64
}

func (g *Graph) init(index int, result Result) []Result {
	results := make([]Result, len(*g))
	results[index] = result
	return results
}

func (g *Graph) DFS(
	initIndex int,
	initResult Result,
	fn func(current, next []Edge, currentResult Result) Result,
) []Result {
	results := g.init(initIndex, initResult)
	visited := make([]bool, len(*g))
	visited[initIndex] = true
	stack := NewStack(len(*g))
	stack.Push(initIndex)

	for !stack.Empty() {
		currentIndex := stack.Pop()
		for _, next := range (*g)[currentIndex] {
			if visited[next.To] {
				continue
			}

			results[next.To] = fn((*g)[currentIndex], (*g)[next.To], results[currentIndex])
			visited[next.To] = true
			stack.Push(int(next.To))
		}
	}

	return results
}

type Result struct {
	Distance int64
}

func solve(N int64, A []int64, B []int64) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	graph := make([][]Edge, N)
	for i := int64(0); i < N-1; i++ {
		graph[A[i]-1] = append(graph[A[i]-1], Edge{To: B[i] - 1})
		graph[B[i]-1] = append(graph[B[i]-1], Edge{To: A[i] - 1})
	}

	g := Graph(graph)
	results := g.DFS(
		0,
		Result{Distance: 0},
		func(current, next []Edge, currentResult Result) Result {
			return Result{
				Distance: currentResult.Distance + 1,
			}
		},
	)

	maxIndex := 0
	for i, r := range results {
		if r.Distance > results[maxIndex].Distance {
			maxIndex = i
		}
	}

	results = g.DFS(
		maxIndex,
		Result{Distance: 0},
		func(current, next []Edge, currentResult Result) Result {
			return Result{
				Distance: currentResult.Distance + 1,
			}
		},
	)

	max := int64(0)
	for _, r := range results {
		if r.Distance > max {
			max = r.Distance
		}
	}

	fmt.Fprintln(w, max+1)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1048576
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)

	var err error
	scanner.Scan()
	N, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	A := make([]int64, N-1)
	B := make([]int64, N-1)
	for i := int64(0); i < N-1; i++ {
		scanner.Scan()
		A[i], err = strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
		scanner.Scan()
		B[i], err = strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	solve(N, A, B)
}
