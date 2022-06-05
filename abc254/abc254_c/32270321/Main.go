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

func solve(N int, K int, a []int) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	if K == 1 {
		fmt.Fprintln(w, yes)
		return
	}

	slices := make([][]int, K)
	for i := 0; i < K; i++ {
		slices[i] = make([]int, 0, N/K+1)
		for j := i; j < N; j += K {
			slices[i] = append(slices[i], a[j])
		}
		sort.Ints(slices[i])
		// fmt.Fprintln(w, slices[i], i)
	}

	current := 0
	for i := 0; i < N; i++ {
		next := slices[i%K][i/K]
		if next < current {
			fmt.Fprintln(w, no)
			return
		}
		current = next
	}

	fmt.Fprintln(w, yes)
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
	scanner.Scan()
	K, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}
	a := make([]int, N)
	for i := 0; i < N; i++ {
		scanner.Scan()
		a[i], err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	solve(N, K, a)
}
