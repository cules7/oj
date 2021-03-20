package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// var file, _ = os.Open("input.txt")
// var outfile, _ = os.Create("output.txt")
// var in = bufio.NewScanner(file)
// var out = bufio.NewWriter(outfile)
var in = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func next() string {
	if !in.Scan() {
		panic("Scan error")
	}
	return in.Text()
}

func nextInt() int {
	ret, _ := strconv.Atoi(next())
	return ret
}

func nextFloat() float64 {
	ret, _ := strconv.ParseFloat(next(), 64)
	return ret
}

func main() {
	defer out.Flush()
	in.Split(bufio.ScanWords)
	n := nextInt()
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = nextInt()
	}
	for i := 1; i < n; i++ {
		a[i] = a[i] + a[i-1]
	}
	m := nextInt()
	for i := 0; i < m; i++ {
		x := nextInt()
		l := 0
		r := n - 1
		for l < r {
			m := l + (r-l)/2
			if x > a[m] {
				l = m + 1
			} else {
				r = m
			}
		}
		fmt.Fprintln(out, r+1)
	}

}
