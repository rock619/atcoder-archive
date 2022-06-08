package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N int) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	fMap := make(map[int]int)
	squares := squaresUnder(1000000000)
	for i := 1; i <= N; i++ {
		maxSquareDivisor := 1
		for _, square := range squares {
			if square > i {
				break
			}
			if i%square == 0 {
				maxSquareDivisor = square
			}
		}

		fMap[i/maxSquareDivisor]++
	}

	result := 0
	for _, v := range fMap {
		result += v * v
	}
	fmt.Fprintln(w, result)
}

func squaresUnder(n int) []int {
	result := make([]int, 0, 1000)
	for i := 2; i*i <= n; i++ {
		result = append(result, i*i)
	}
	return result
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 4096), 1048576)
	scanner.Split(bufio.ScanWords)

	var err error
	scanner.Scan()
	N, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	solve(N)
}
