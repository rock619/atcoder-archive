package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const mod = 1000000007

type Pair struct {
	Index int64
	Legal bool
}

type DFS struct {
	H, W   int64
	Counts []int64
	States [][]int64
	Used   [][]bool
	M      map[int64]map[int64]Pair
}

func NewDFS(h, w int64) *DFS {
	used := make([][]bool, h+1)
	states := make([][]int64, w+1)
	m := make(map[int64]map[int64]Pair, w+1)
	for i := int64(0); i <= w; i++ {
		states[i] = make([]int64, 1<<w+1)
		m[i] = make(map[int64]Pair)
	}
	for i := range used {
		used[i] = make([]bool, w+1)
	}

	return &DFS{
		H:      h,
		W:      w,
		Counts: make([]int64, w+1),
		States: states,
		Used:   used,
		M:      m,
	}
}

func (dfs *DFS) isLegal(x, y int64) bool {
	dx := []int64{1, 1, 1, 0, -1, -1, -1, 0}
	dy := []int64{-1, 0, 1, 1, 1, 0, -1, -1}
	for i := range dx {
		tx, ty := x+dx[i], y+dy[i]
		if tx < 0 || ty < 0 || tx > dfs.H || ty >= dfs.W {
			continue
		}

		if dfs.Used[tx][ty] {
			return false
		}
	}
	return true
}

func (dfs *DFS) Do(position, depth, str int64) {
	sx, sy := position/dfs.W, position%dfs.W

	if depth == dfs.W+1 {
		index := dfs.Counts[sy]
		legal := dfs.isLegal(sx, sy)
		dfs.States[sy][index] = str
		dfs.M[sy][str] = Pair{Index: index, Legal: legal}
		dfs.Counts[sy]++
		return
	}

	dfs.Do(position+1, depth+1, str)

	if dfs.isLegal(sx, sy) {
		dfs.Used[sx][sy] = true
		dfs.Do(position+1, depth+1, str+(1<<depth))
		dfs.Used[sx][sy] = false
	}
}

func solve(H, W int64, C [][]byte) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	dfs := NewDFS(H, W)
	for i := int64(0); i < W; i++ {
		dfs.Do(i, 0, 0)
	}

	next0, next1 := make([][]int64, W+1), make([][]int64, W+1)
	for i := int64(0); i < W; i++ {
		next0[i], next1[i] = make([]int64, 1<<W+1), make([]int64, 1<<W+1)

		for j := int64(0); j < dfs.Counts[i]; j++ {
			t := dfs.States[i][j]
			t0, t1 := t>>1, (t>>1)+(1<<W)
			next0[i][j] = dfs.M[(i+1)%W][t0].Index

			if dfs.M[i][t].Legal {
				next1[i][j] = dfs.M[(i+1)%W][t1].Index
			} else {
				next1[i][j] = -1
			}
		}
	}

	dp := make([][][]int64, H+1)
	for i := int64(0); i <= H; i++ {
		dp[i] = make([][]int64, W+1)
		for j := int64(0); j <= W; j++ {
			dp[i][j] = make([]int64, dfs.Counts[j])
		}
	}
	dp[0][0][0] = 1

	for i := int64(0); i < H; i++ {
		for j := int64(0); j < W; j++ {
			n1, n2 := i, j+1
			if n2 == W {
				n1++
				n2 = 0
			}

			for k := int64(0); k < dfs.Counts[j]; k++ {
				if dp[i][j][k] == 0 {
					continue
				}

				dp[n1][n2][next0[j][k]] = addMod(dp[n1][n2][next0[j][k]], dp[i][j][k], mod)

				if next1[j][k] != -1 && C[i][j] == '.' {
					dp[n1][n2][next1[j][k]] = addMod(dp[n1][n2][next1[j][k]], dp[i][j][k], mod)
				}
			}
		}
	}

	sum := int64(0)
	for i := int64(0); i < dfs.Counts[0]; i++ {
		sum = addMod(sum, dp[H][0][i], mod)
	}
	fmt.Fprintln(w, sum)
}

func addMod(a, b, mod int64) int64 {
	a += b % mod
	a %= mod
	return a
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1048576
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)

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
	C := make([][]byte, H)
	for i := int64(0); i < H; i++ {
		scanner.Scan()
		C[i] = scanner.Bytes()
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	solve(H, W, C)
}
