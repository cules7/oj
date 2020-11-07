package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	//"math"
)

func solve() {
	_, s := readInt(), readString()
	a, d := strings.Count(s, "A"), strings.Count(s, "D")
	ans := "Friendship"
	switch {
	case a > d:
		ans = "Anton"
	case a < d:
		ans = "Danik"
	}
	fmt.Println(ans)
}

var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)

func main() {
	const MaxTokenLength = 1024 * 1024
	scanner.Buffer(make([]byte, 0, MaxTokenLength), MaxTokenLength)
	scanner.Split(bufio.ScanWords)
	solve()
}

// Math
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	return x + y - max(x, y)
}

// IO

func readString() string {
	scanner.Scan()
	return scanner.Text()
}

func readInt() int {
	val, _ := strconv.Atoi(readString())
	return val
}

func readIntArray(sz int) []int {
	a := make([]int, sz)
	for i := 0; i < sz; i++ {
		a[i] = readInt()
	}
	return a
}

func readArray(n int) []int {
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = readInt()
	}
	return res
}
