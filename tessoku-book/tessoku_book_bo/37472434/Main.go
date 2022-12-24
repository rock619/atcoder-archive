package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
	"strconv"
)

type Edge struct {
	To, Weight int
}

type PriorityQueueValue = int

type PriorityQueueItem struct {
	Value    PriorityQueueValue
	Priority int
}

type PriorityQueue struct {
	pq *priorityQueue
	n  int
}

func NewPriorityQueue(capacity int) *PriorityQueue {
	pq := make(priorityQueue, 0, capacity)
	heap.Init(&pq)
	return &PriorityQueue{
		pq: &pq,
		n:  0,
	}
}

func (pq *PriorityQueue) Empty() bool {
	return pq.pq.Len() == 0
}

func (pq *PriorityQueue) Push(x *PriorityQueueItem) {
	heap.Push(pq.pq, &priorityQueueItem{
		PriorityQueueItem: x,
		Index:             pq.n,
	})
	pq.n++
}

func (pq *PriorityQueue) Pop() (x *PriorityQueueItem, ok bool) {
	if pq.Empty() {
		return x, false
	}
	pq.n--
	return heap.Pop(pq.pq).(*priorityQueueItem).PriorityQueueItem, true
}

type priorityQueueItem struct {
	*PriorityQueueItem
	Index int
}

type priorityQueue []*priorityQueueItem

func (pq priorityQueue) Len() int { return len(pq) }

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].Priority > pq[j].Priority
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index, pq[j].Index = i, j
}

func (pq *priorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*priorityQueueItem)
	item.Index = n
	*pq = append(*pq, item)
}

func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.Index = -1
	*pq = old[0 : n-1]
	return item
}

func solve(o Printer, N int, M int, A []int, B []int, C []int) {
	g := make([][]Edge, N+1)
	for i := range A {
		g[A[i]] = append(g[A[i]], Edge{To: B[i], Weight: C[i]})
		g[B[i]] = append(g[B[i]], Edge{To: A[i], Weight: C[i]})
	}
	pq := NewPriorityQueue(N - 1)
	for _, next := range g[1] {
		pq.Push(&PriorityQueueItem{Value: next.To, Priority: -next.Weight})
	}

	seen := make([]bool, N+1)
	seen[1] = true
	result := 0
	for !pq.Empty() {
		x, _ := pq.Pop()
		if seen[x.Value] {
			continue
		}
		seen[x.Value] = true
		result += -x.Priority
		for _, next := range g[x.Value] {
			pq.Push(&PriorityQueueItem{Value: next.To, Priority: -next.Weight})
		}
	}

	o.l(result)
}

func main() {
	sc := NewScanner()
	N := sc.Int()
	M := sc.Int()
	A := make([]int, M)
	B := make([]int, M)
	C := make([]int, M)
	for i := 0; i < M; i++ {
		A[i] = sc.Int()
		B[i] = sc.Int()
		C[i] = sc.Int()
	}
	out := NewPrinter()
	solve(out, N, M, A, B, C)
	if err := out.w.Flush(); err != nil {
		panic(err)
	}
}

type Scanner struct {
	*bufio.Scanner
}

func NewScanner() *Scanner {
	s := bufio.NewScanner(os.Stdin)
	s.Buffer(make([]byte, 4096), math.MaxInt64)
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

type Printer interface {
	// p fmt.Print
	p(a ...interface{})
	// f fmt.Printf
	f(format string, a ...interface{})
	// l fmt.Println
	l(a ...interface{})
}

type printer struct {
	w *bufio.Writer
}

func NewPrinter() *printer {
	return &printer{bufio.NewWriter(os.Stdout)}
}

func (p *printer) p(a ...interface{}) {
	fmt.Fprint(p.w, a...)
}

func (p *printer) f(format string, a ...interface{}) {
	fmt.Fprintf(p.w, format, a...)
}

func (p *printer) l(a ...interface{}) {
	fmt.Fprintln(p.w, a...)
}
