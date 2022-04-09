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

type Vector struct {
	X, Y int64
}

func (v Vector) GTE(v2 Vector) bool {
	if v.X == v2.X {
		return v.Y >= v2.Y
	}
	return v.X > v2.X
}

func (v Vector) Sub(v2 Vector) Vector {
	return Vector{
		X: v.X - v2.X,
		Y: v.Y - v2.Y,
	}
}

func Cross(a, b Vector) int64 {
	return a.X*b.Y - a.Y*b.X
}

func solve(x []int64, y []int64) {
	v1, v2, v3, v4 := Vector{x[0], y[0]}, Vector{x[1], y[1]}, Vector{x[2], y[2]}, Vector{x[3], y[3]}
	if v1.GTE(v2) {
		v1, v2 = v2, v1
	}
	if v3.GTE(v4) {
		v3, v4 = v4, v3
	}

	v12 := v2.Sub(v1)
	v13 := v3.Sub(v1)
	v14 := v4.Sub(v1)
	v31 := v1.Sub(v3)
	v32 := v2.Sub(v3)
	v34 := v4.Sub(v3)

	if Cross(v12, v13) == 0 && Cross(v12, v14) == 0 && Cross(v34, v31) == 0 && Cross(v34, v32) == 0 {
		if Min(v2, v4).GTE(Max(v1, v3)) {
			fmt.Println(YES)
		} else {
			fmt.Println(NO)
		}
		return
	}

	if OppositeSign(Cross(v12, v13), Cross(v12, v14)) && OppositeSign(Cross(v34, v31), Cross(v34, v32)) {
		fmt.Println(YES)
	} else {
		fmt.Println(NO)
	}
}

func Max(a, b Vector) Vector {
	if a.GTE(b) {
		return a
	}
	return b
}

func Min(a, b Vector) Vector {
	if a.GTE(b) {
		return b
	}
	return a
}

func OppositeSign(a, b int64) bool {
	return a >= 0 && b <= 0 || a <= 0 && b >= 0
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	x := make([]int64, 4)
	y := make([]int64, 4)
	for i := int64(0); i < 4; i++ {
		scanner.Scan()
		x[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
		scanner.Scan()
		y[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	solve(x, y)
}
