package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	yes = "Yes"
	no  = "No"
)

func solve(a int64, b int64, c int64) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	if c == 1 {
		fmt.Println(no)
		return
	}

	v := int64(1)
	for i := int64(1); i <= b; i++ {
		if a/c < v {
			fmt.Println(yes)
			return
		}
		v *= c
	}

	fmt.Println(no)
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
