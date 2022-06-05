package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Stack []int64

func (s Stack) Size() int64 {
	return int64(len(s))
}

func (s Stack) Empty() bool {
	return s.Size() == 0
}

func (s *Stack) Push(v int64) {
	*s = append(*s, v)
}

func (s *Stack) Pop() int64 {
	if s.Empty() {
		panic("*Stack.Pop(): stack is empty")
	}
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v
}

func UpdateMax(max *int64, v int64) {
	if v > *max {
		*max = v
	}
}

type DFS struct {
	H, W           int64
	c              [][]byte
	using          [][]bool
	startX, startY int64
	maxDepth       int64
}

func NewDFS(H, W int64, c [][]byte, startX, startY int64) *DFS {
	using := make([][]bool, H)
	for i := range using {
		using[i] = make([]bool, W)
	}
	return &DFS{
		H:        H,
		W:        W,
		c:        c,
		using:    using,
		startX:   startX,
		startY:   startY,
		maxDepth: -1,
	}
}

func (dfs *DFS) Do(depth, x, y int64) {
	if dfs.using[y][x] {
		if x == dfs.startX && y == dfs.startY && depth >= 3 {
			UpdateMax(&dfs.maxDepth, depth)
		}
		return
	}

	dfs.using[y][x] = true

	for _, diff := range [][]int64{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
		nextX, nextY := x+diff[0], y+diff[1]
		if nextX < 0 || nextX >= dfs.W || nextY < 0 || nextY >= dfs.H || dfs.c[nextY][nextX] == '#' {
			continue
		}
		dfs.Do(depth+1, nextX, nextY)
	}

	dfs.using[y][x] = false
}

func solve(H, W int64, c [][]byte) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	if H == 1 || W == 1 {
		fmt.Fprintln(w, -1)
		return
	}

	max := int64(-1)
	for i := int64(0); i < H; i++ {
		for j := int64(0); j < W; j++ {
			if c[i][j] == '#' {
				continue
			}
			dfs := NewDFS(H, W, c, j, i)
			dfs.Do(0, j, i)
			UpdateMax(&max, dfs.maxDepth)
		}
	}

	fmt.Fprintln(w, max)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1048576
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)

	var err error
	scanner.Scan()
	H, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	scanner.Scan()
	W, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	c := make([][]byte, H)
	for i := int64(0); i < H; i++ {
		scanner.Scan()
		c[i] = scanner.Bytes()
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	solve(H, W, c)
}
