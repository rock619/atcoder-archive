package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Point struct {
	X, Y int64
}

type Queue struct {
	s []Point
}

func NewQueue(cap int64) *Queue {
	return &Queue{
		s: make([]Point, 0, cap),
	}
}

func (q Queue) Size() int64 {
	return int64(len(q.s))
}

func (q *Queue) Enqueue(e Point) {
	q.s = append(q.s, e)
}

func (q *Queue) Dequeue() Point {
	if q.Size() == 0 {
		panic("Queue.Dequeue(): queue is empty")
	}
	e := q.s[0]
	q.s = q.s[1:]
	return e
}

func solve(R, C, sy, sx, gy, gx int64, c [][]byte) {
	results := make([][]int64, R)
	for i := range results {
		results[i] = make([]int64, C)
		for j := range results[i] {
			results[i][j] = -1
		}
	}
	c[sy-1][sx-1] = 'x'
	results[sy-1][sx-1] = 0

	queue := NewQueue(1)
	queue.Enqueue(Point{X: sx - 1, Y: sy - 1})

	for queue.Size() > 0 {
		p := queue.Dequeue()

		for _, next := range []Point{{p.X + 1, p.Y}, {p.X - 1, p.Y}, {p.X, p.Y + 1}, {p.X, p.Y - 1}} {
			if v := c[next.Y][next.X]; v == '#' || v != '.' {
				continue
			}
			c[next.Y][next.X] = 'x'
			results[next.Y][next.X] = results[p.Y][p.X] + 1
			queue.Enqueue(Point{next.X, next.Y})
		}
	}

	fmt.Println(results[gy-1][gx-1])
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var R int64
	scanner.Scan()
	R, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	var C int64
	scanner.Scan()
	C, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	var sy int64
	scanner.Scan()
	sy, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	var sx int64
	scanner.Scan()
	sx, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	var gy int64
	scanner.Scan()
	gy, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	var gx int64
	scanner.Scan()
	gx, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	c := make([][]byte, R)
	for i := int64(0); i < R; i++ {
		scanner.Scan()
		c[i] = scanner.Bytes()
	}
	solve(R, C, sy, sx, gy, gx, c)
}
