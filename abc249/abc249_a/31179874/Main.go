package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(A int64, B int64, C int64, D int64, E int64, F int64, X int64) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	distT := int64(0)
	distT += B * A * (X / (A + C))
	if restT := X % (A + C); restT >= A {
		distT += B * A
	} else {
		distT += B * restT
	}

	distA := int64(0)
	distA += E * D * (X / (D + F))
	if restA := X % (D + F); restA >= D {
		distA += E * D
	} else {
		distA += E * restA
	}

	switch diff := distT - distA; {
	case diff > 0:
		fmt.Fprintln(w, "Takahashi")
	case diff < 0:
		fmt.Fprintln(w, "Aoki")
	default:
		fmt.Fprintln(w, "Draw")
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
	A, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	scanner.Scan()
	B, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	scanner.Scan()
	C, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	scanner.Scan()
	D, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	scanner.Scan()
	E, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	scanner.Scan()
	F, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	scanner.Scan()
	X, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	solve(A, B, C, D, E, F, X)
}
