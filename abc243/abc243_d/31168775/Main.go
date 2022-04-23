package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N int64, X int64, S string) {
	s := Compact(S)
	result := X
	for _, move := range s {
		switch move {
		case 'U':
			result /= 2
		case 'L':
			result *= 2
		default:
			result = result*2 + 1
		}
	}
	fmt.Println(result)
}

func Compact(s string) string {
	b := []byte(s)
	result := make([]byte, 0, len(b))
	for _, r := range b {
		if r == 'U' && len(result) > 0 && result[len(result)-1] != 'U' {
			result = result[:len(result)-1]
			continue
		}
		result = append(result, r)

	}

	return string(result)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 10000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	N, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	scanner.Scan()
	X, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	scanner.Scan()
	S := scanner.Text()
	solve(N, X, S)
}
