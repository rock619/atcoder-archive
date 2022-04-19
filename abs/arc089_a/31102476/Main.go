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

func solve(N int64, t []int64, x []int64, y []int64) {
	currentT, currentX, currentY := int64(0), int64(0), int64(0)
	for i := 0; i < len(x); i++ {
		if rest := (t[i] - currentT) - (Abs(x[i]-currentX) + Abs(y[i]-currentY)); rest < 0 || rest%2 == 1 {
			fmt.Println(NO)
			return
		}

		currentT, currentX, currentY = t[i], x[i], y[i]
	}
	fmt.Println(YES)
}

func Abs(n int64) int64 {
	if n < 0 {
		return -n
	}
	return n
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	N, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	t := make([]int64, N)
	x := make([]int64, N)
	y := make([]int64, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		t[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
		scanner.Scan()
		x[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
		scanner.Scan()
		y[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	solve(N, t, x, y)
}
