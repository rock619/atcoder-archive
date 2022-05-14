package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const inf int64 = 1 << 62

type State struct {
	X, Y, Dir int64
}

type Deque []State

func (d Deque) Empty() bool {
	return len(d) == 0
}

func (d *Deque) PopFront() State {
	if d.Empty() {
		panic("Deque.PopFront(): deque is empty")
	}
	e := (*d)[0]
	*d = (*d)[1:]
	return e
}

func (d *Deque) PushBack(e State) {
	*d = append(*d, e)
}

func (d *Deque) PushFront(e State) {
	*d = append([]State{e}, *d...)
}

func Min(ints ...int64) int64 {
	if len(ints) == 0 {
		panic("Min: len(ints) == 0")
	}
	m := ints[0]
	for i := 1; i < len(ints); i++ {
		if ints[i] < m {
			m = ints[i]
		}
	}
	return m
}

func solve(H, W, rs, cs, rt, ct int64, S []string) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	sx, sy := rs-1, cs-1
	gx, gy := rt-1, ct-1
	dx := []int64{1, 0, -1, 0}
	dy := []int64{0, 1, 0, -1}

	dist := make([][][]int64, H)
	for i := range dist {
		dist[i] = make([][]int64, W)
		for j := range dist[i] {
			dist[i][j] = make([]int64, 4)
			for k := 0; k < 4; k++ {
				dist[i][j][k] = inf
			}
		}
	}

	deq := Deque(make([]State, 0, 4))
	for i := int64(0); i < 4; i++ {
		dist[sx][sy][i] = 0
		deq.PushBack(State{
			X:   sx,
			Y:   sy,
			Dir: i,
		})
	}

	for !deq.Empty() {
		u := deq.PopFront()
		for i := int64(0); i < 4; i++ {
			tx := u.X + dx[i]
			ty := u.Y + dy[i]
			cost := dist[u.X][u.Y][u.Dir]
			if u.Dir != i {
				cost++
			}
			if tx >= 0 && tx < H && ty >= 0 && ty < W && S[tx][ty] == '.' && dist[tx][ty][i] > cost {
				dist[tx][ty][i] = cost
				s := State{
					X:   tx,
					Y:   ty,
					Dir: i,
				}
				if u.Dir != i {
					deq.PushBack(s)
				} else {
					deq.PushBack(s)
				}
			}
		}
	}

	fmt.Fprintln(w, Min(dist[gx][gy]...))
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
	scanner.Scan()
	rs, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	scanner.Scan()
	cs, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	scanner.Scan()
	rt, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	scanner.Scan()
	ct, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	S := make([]string, H)
	for i := int64(0); i < H; i++ {
		scanner.Scan()
		S[i] = scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	solve(H, W, rs, cs, rt, ct, S)
}
