package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"strconv"
)

func solve(H int64, W int64, P [][]int64) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	max := int64(0)
	for i := int64(1); i < (1 << H); i++ {
		m := make(map[int64]int64, W)
		for j := int64(0); j < W; j++ {
			same := true
			n := int64(0)
			for k := int64(0); k < H; k++ {
				if i&(1<<k) != 0 {
					if n == 0 {
						n = P[k][j]
						continue
					}
					if n != P[k][j] {
						same = false
						break
					}
				}
			}
			if same {
				m[n]++
			}
		}
		vMax := int64(0)
		for _, v := range m {
			if v > vMax {
				vMax = v
			}
		}
		if result := vMax * int64(bits.OnesCount64(uint64(i))); result > max {
			max = result
		}
	}
	fmt.Fprintln(w, max)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1048576
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)

	var err error
	scanner.Scan()
	H, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	scanner.Scan()
	W, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	P := make([][]int64, H)
	for i := int64(0); i < H; i++ {
		P[i] = make([]int64, W)
	}
	for i := int64(0); i < H; i++ {
		for j := int64(0); j < W; j++ {
			scanner.Scan()
			P[i][j], err = strconv.ParseInt(scanner.Text(), 10, 64)
			if err != nil {
				panic(err)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	solve(H, W, P)
}
