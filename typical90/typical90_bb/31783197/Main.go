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

func solve(N, M int64, K []int64, R [][]int64) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	graph := make(map[int64][]int64, N+M)
	for i := int64(0); i < M; i++ {
		graph[-(i + 1)] = R[i]

		for j := int64(0); j < K[i]; j++ {
			graph[R[i][j]] = append(graph[R[i][j]], -(i + 1))
		}
	}

	dist := make(map[int64]int64, len(graph))
	queue := Queue(make([]int64, 0, 1))
	queue.Enqueue(1)
	dist[1] = 0
	for !queue.Empty() {
		v := queue.Dequeue()
		for _, c := range graph[v] {
			if _, ok := dist[c]; ok {
				continue
			}
			dist[c] = dist[v] + 1
			queue.Enqueue(c)
		}
	}

	for i := int64(1); i <= N; i++ {
		if d, ok := dist[i]; ok {
			fmt.Fprintln(w, d/2)
		} else {
			fmt.Fprintln(w, -1)
		}
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
	scanner.Scan()
	M, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	K := make([]int64, M)
	R := make([][]int64, M)
	for i := int64(0); i < M; i++ {
		scanner.Scan()
		K[i], err = strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
		R[i] = make([]int64, K[i])
		for j := int64(0); j < K[i]; j++ {
			scanner.Scan()
			R[i][j], err = strconv.ParseInt(scanner.Text(), 10, 64)
			if err != nil {
				panic(err)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	solve(N, M, K, R)
}
