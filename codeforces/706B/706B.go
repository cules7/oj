package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
var prices []int
var d int
var mone []int

func main() {
	defer exit()

	n = nextInt()
	prices = make([]int, n+1, n+1)

	for i := 1; i != n+1; i++ {
		prices[i] = nextInt()
	}
	sort.Ints(prices[1:])

	d := nextInt()
	mone = make([]int, d+1, d+1)
	for i := 1; i != d+1; i++ {
		target := nextInt()
		l := 0
		r := n
		for {
			if l == r {
				break
			}
			mid := (l + r + 1) / 2
			midv := prices[mid]
			if midv > target {
				r = mid - 1
			} else {
				l = mid
			}
		}
		fmt.Fprintln(out, l)
	}

}
