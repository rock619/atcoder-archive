package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	fmt.Fprintln(w, 297)

	for _, i := range []int{1, 100, 10000} {
		if i == 1 {
			fmt.Fprint(w, i)
		} else {
			fmt.Fprintf(w, " %d", i)
		}
		for j := 2; j <= 99; j++ {
			fmt.Fprintf(w, " %d", i*j)
		}
	}
	fmt.Fprintln(w)
}
