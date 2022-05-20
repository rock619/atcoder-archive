package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(Q int64, t []int64, x []int64) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	cards1 := make([]int64, 0, Q)
	cards2 := make([]int64, 0, Q)
	for i := int64(0); i < Q; i++ {
		switch {
		case t[i] == 1:
			cards1 = append(cards1, x[i])
		case t[i] == 2:
			cards2 = append(cards2, x[i])
		case x[i] <= int64(len(cards1)):
			fmt.Fprintln(w, cards1[int64(len(cards1))-x[i]])
		default:
			fmt.Fprintln(w, cards2[x[i]-int64(len(cards1))-1])
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
	Q, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	t := make([]int64, Q)
	x := make([]int64, Q)
	for i := int64(0); i < Q; i++ {
		scanner.Scan()
		t[i], err = strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
		scanner.Scan()
		x[i], err = strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	solve(Q, t, x)
}
