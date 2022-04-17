package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const MOD = 998244353

func solve(A int64, B int64, C int64) {
	fmt.Println(Multiply(MOD, A*(A+1)/2, B*(B+1)/2, C*(C+1)/2))
}

func Multiply(mod int64, ints ...int64) int64 {
	result := ints[0] % mod
	for i := 1; i < len(ints); i++ {
		result = ints[i] % mod * result % mod
	}
	return result
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var A int64
	scanner.Scan()
	A, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	var B int64
	scanner.Scan()
	B, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	var C int64
	scanner.Scan()
	C, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	solve(A, B, C)
}
