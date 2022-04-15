package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(H int64, W int64, A [][]int64) {
	rowSums := make([]int64, H)
	colSums := make([]int64, W)
	for row := 0; row < len(rowSums); row++ {
		for col := 0; col < len(colSums); col++ {
			n := A[row][col]
			rowSums[row] += n
			colSums[col] += n
		}
	}

	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	for row := 0; row < len(rowSums); row++ {
		for col := 0; col < len(colSums); col++ {
			if col != 0 {
				fmt.Fprint(w, " ")
			}
			fmt.Fprint(w, rowSums[row]+colSums[col]-A[row][col])
		}
		fmt.Fprintln(w)
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var H int64
	scanner.Scan()
	H, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	var W int64
	scanner.Scan()
	W, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	A := make([][]int64, H)
	for i := int64(0); i < H; i++ {
		A[i] = make([]int64, W)
	}
	for i := int64(0); i < H; i++ {
		for j := int64(0); j < W; j++ {
			scanner.Scan()
			A[i][j], _ = strconv.ParseInt(scanner.Text(), 10, 64)
		}
	}
	solve(H, W, A)
}
