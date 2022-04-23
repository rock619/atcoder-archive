package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

const (
	YES = "Yes"
	NO  = "No"
)

type Person struct {
	X, Y int64
	D    rune
}

func solve(N int64, X []int64, Y []int64, S string) {
	m := make(map[int64][]Person, N)
	for i := int64(0); i < N; i++ {
		p := Person{
			X: X[i],
			Y: Y[i],
			D: rune(S[i]),
		}
		m[p.Y] = append(m[p.Y], p)
	}

	for _, people := range m {
		if len(people) == 1 {
			continue
		}

		sort.Slice(people, func(i, j int) bool {
			return people[i].X < people[j].X
		})

		toRExists := false
		for _, p := range people {
			if toRExists && p.D == 'L' {
				fmt.Println(YES)
				return
			}
			if p.D == 'R' {
				toRExists = true
			}
		}
	}

	fmt.Println(NO)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	N, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	X := make([]int64, N)
	Y := make([]int64, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		X[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
		scanner.Scan()
		Y[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	scanner.Scan()
	S := scanner.Text()
	solve(N, X, Y, S)
}
