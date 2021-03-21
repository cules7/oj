package main

import (
	"bufio"
	"os"
	"strconv"
)

var in = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func init() {
	in.Split(bufio.ScanWords)
}
func exit() {
	out.Flush()
}

// custom
func nextInt() int {
	in.Scan()
	ret, _ := strconv.Atoi(in.Text())
	return ret
}

/*
 */
var n int
var sites []int
var vs []int

func time(leftIdx, rightIdx int) float64 {
	return float64(sites[rightIdx]-sites[leftIdx]) / float64(vs[rightIdx]-vs[leftIdx])
}

func main() {
	defer exit()

	n = nextInt()
	sites = make([]int, n+1, n+1)

	for i := 1; i != n+1; i++ {
		sites[i] = nextInt()
	}

	for i := 1; i != n+1; i++ {
		vs[i] = nextInt()
	}

	var left int
	var right int
	if sites[1] > sites[2] {
		left = 2
		right = 1
	} else {
		left = 1
		right = 2
	}
	dptime := time(left, right)
	for i := 3; i != n+1; i++ {

	}
}
