package main

import (
	"bufio"
	"fmt"
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
var piles []int
var t int

func main() {
	defer exit()

	n = nextInt()
	piles = make([]int, n+1, n+1)
	piles[1] = 1
	for i := 2; i != n+1; i++ {
		piles[i] = piles[i-1] + nextInt()
	}
	t = nextInt()
	t = nextInt()
	for i := 1; i != t+1; i++ {
		ts := nextInt()
		l := 1
		r := n
		for {
			if l >= r {
				break
			}
			mid := (l + r + 1) / 2
			midv := piles[mid]
			if ts < midv {
				r = mid - 1
			} else {
				l = mid
			}
		}
		fmt.Fprintln(out, l)
	}
}
