package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(a int64, b int64, c int64, s string) {
	fmt.Printf("%d %s\n", a+b+c, s)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	a, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	scanner.Scan()
	b, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	scanner.Scan()
	c, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	scanner.Scan()
	s := scanner.Text()
	solve(a, b, c, s)
}
