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
	buf := make([]byte, 2500000, 2500000)
	in.Buffer(buf, 0)
}
func exit() {
	out.Flush()
}

// io
func nextInt() int {
	in.Scan()
	ret, _ := strconv.Atoi(in.Text())
	return ret
}
func nextString() string {
	in.Scan()
	return in.Text()
}
func nextBytes() []byte {
	in.Scan()
	return in.Bytes()
}

/*
 */
var num int64
var left int64
var timeSum []int64

func main() {
	defer exit()
	num = int64(nextInt())
	left = int64(nextInt())
	timeSum = make([]int64, num+1, num+1)

	for i := int64(1); i != num+1; i++ {
		timeSum[i] = timeSum[i-1] + int64(nextInt())
	}
	ans := int64(0)

	for i := int64(1); i != num+1; i++ {
		minusBase := timeSum[i-1]
		// test one
		if timeSum[i]-minusBase > left {
			continue
		}
		l := i
		r := num
		for {
			if l == r {
				break
			}
			mid := (l + r + 1) / 2
			midv := left >= (timeSum[mid] - minusBase)
			if midv {
				l = mid
			} else {
				r = mid - 1
			}
		}
		temp := l - i + 1
		if temp > ans {
			ans = temp
		}
	}
	fmt.Fprintf(out, "%d\n", ans)
}
