package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(H int64, W int64, S []string) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	ai, aj := int64(-1), int64(-1)
	bi, bj := int64(-1), int64(-1)
	for i := int64(0); i < H; i++ {
		for j := int64(0); j < W; j++ {
			if S[i][j] == 'o' {
				if ai == -1 && aj == -1 {
					ai, aj = i, j
				} else {
					bi, bj = i, j
				}
			}
		}
	}
	fmt.Fprintln(w, Abs(ai-bi)+Abs(aj-bj))
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
	S := make([]string, H)
	for i := int64(0); i < H; i++ {
		scanner.Scan()
		S[i] = scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	solve(H, W, S)
}
