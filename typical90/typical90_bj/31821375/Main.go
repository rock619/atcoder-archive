package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Queue []int64

func (q *Queue) Enqueue(v int64) {
	*q = append(*q, v)
}

func (q *Queue) Dequeue() int64 {
	if q.Empty() {
		panic("*Queue.Dequeue(): empty")
	}
	v := (*q)[0]
	*q = (*q)[1:]
	return v
}

func (q Queue) Empty() bool {
	return len(q) == 0
}

func Reverse(s []int64) {
	for left, right := 0, len(s)-1; left < right; left, right = left+1, right-1 {
		s[left], s[right] = s[right], s[left]
	}
}

func solve(N int64, A []int64, B []int64) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	graph := make([][]int64, N)
	queue := Queue(make([]int64, 0, N))
	visited := make([]bool, N)
	for i := int64(0); i < N; i++ {
		a, b := A[i]-1, B[i]-1
		graph[a] = append(graph[a], i)
		graph[b] = append(graph[b], i)
		if a == i || b == i {
			queue.Enqueue(i)
			visited[i] = true
		}
	}

	result := make([]int64, 0, N)

	for !queue.Empty() {
		current := queue.Dequeue()
		result = append(result, current+1)
		for _, next := range graph[current] {
			if !visited[next] {
				visited[next] = true
				queue.Enqueue(next)
			}
		}
	}

	if int64(len(result)) != N {
		fmt.Fprintln(w, -1)
		return
	}
	Reverse(result)
	for _, r := range result {
		fmt.Fprintln(w, r)
	}
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
	A := make([]int64, N)
	B := make([]int64, N)
	for i := int64(0); i < N; i++ {
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
