package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func solve(a_x int64, a_y int64, b_x int64, b_y int64, c_x int64, c_y int64) {
	bax, bay := a_x-b_x, a_y-b_y
	bcx, bcy := c_x-b_x, c_y-b_y
	cax, cay := a_x-c_x, a_y-c_y
	cbx, cby := b_x-c_x, b_y-c_y

	pattern := 2
	if bax*bcx+bay*bcy < 0 {
		pattern = 1
	} else if cax*cbx+cay*cby < 0 {
		pattern = 3
	}

	ans := float64(0)
	switch pattern {
	case 1:
		ans = math.Sqrt(float64(bax*bax + bay*bay))
	case 3:
		ans = math.Sqrt(float64(cax*cax + cay*cay))
	default:
		s := math.Abs(float64(bax*cay - bay*cax))
		bc := math.Sqrt(float64(bcx*bcx + bcy*bcy))
		ans = s / bc
	}

	fmt.Println(ans)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var a_x int64
	scanner.Scan()
	a_x, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	var a_y int64
	scanner.Scan()
	a_y, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	var b_x int64
	scanner.Scan()
	b_x, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	var b_y int64
	scanner.Scan()
	b_y, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	var c_x int64
	scanner.Scan()
	c_x, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	var c_y int64
	scanner.Scan()
	c_y, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	solve(a_x, a_y, b_x, b_y, c_x, c_y)
}
