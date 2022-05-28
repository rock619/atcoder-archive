package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

const yes = "Yes"
const no = "No"

func solve(a int64, b int64, c int64) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	s := []int64{a, b, c}
	sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
	if b == s[1] {
		fmt.Fprintln(w, yes)
	} else {
		fmt.Fprintln(w, no)
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
	a, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	scanner.Scan()
	b, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	scanner.Scan()
	c, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	solve(a, b, c)
}
