package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func solve(A int64, B int64, H int64, M int64) {
	fmt.Println(Dist(
		float64(A)*XOfHourHand(H, M),
		float64(B)*XOfMinuteHand(M),
		float64(A)*YOfHourHand(H, M),
		float64(B)*YOfMinuteHand(M),
	))
}

func XOfHourHand(hour, min int64) float64 {
	return math.Cos(math.Pi * float64(180-(60*hour+min)) / 360)
}

func YOfHourHand(hour, min int64) float64 {
	return math.Sin(math.Pi * float64(180-(60*hour+min)) / 360)
}

func XOfMinuteHand(min int64) float64 {
	return math.Cos(math.Pi * float64(15-min) / 30)
}

func YOfMinuteHand(min int64) float64 {
	return math.Sin(math.Pi * float64(15-min) / 30)
}

func Dist(x1, x2, y1, y2 float64) float64 {
	return math.Sqrt((x1-x2)*(x1-x2) + (y1-y2)*(y1-y2))
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
	var H int64
	scanner.Scan()
	H, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	var M int64
	scanner.Scan()
	M, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	solve(A, B, H, M)
}
