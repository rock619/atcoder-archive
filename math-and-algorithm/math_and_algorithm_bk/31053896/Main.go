package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Queue struct {
	s []int64
}

func NewQueue(cap int64) *Queue {
	return &Queue{
		s: make([]int64, 0, cap),
	}
}

func (q Queue) Size() int64 {
	return int64(len(q.s))
}

func (q *Queue) Front() int64 {
	if q.Size() == 0 {
		panic("Queue.Front(): queue is empty")
	}
	return q.s[0]
}

func (q *Queue) Enqueue(e int64) {
	q.s = append(q.s, e)
}

func (q *Queue) Dequeue() int64 {
	if q.Size() == 0 {
		panic("Queue.Dequeue(): queue is empty")
	}
	e := q.s[0]
	q.s = q.s[1:]
	return e
}

func solve(N int64, M int64, A []int64, B []int64) {
	graph := make([][]int64, N)
	for i, a := range A {
		graph[a-1] = append(graph[a-1], B[i]-1)
		graph[B[i]-1] = append(graph[B[i]-1], a-1)
	}

	results := make([]int64, N)
	for i := range results {
		results[i] = -1
	}
	results[0] = 0

	queue := NewQueue(1)
	queue.Enqueue(0)

	for queue.Size() > 0 {
		p := queue.Dequeue()

		for _, next := range graph[p] {
			if results[next] != -1 {
				continue
			}
			results[next] = results[p] + 1
			queue.Enqueue(next)
		}
	}

	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	for _, r := range results {
		if r == -1 || r > 120 {
			r = 120
		}
		fmt.Fprintln(w, r)
	}
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
