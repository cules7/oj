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
func nextInt64() int64 {
	in.Scan()
	ret, _ := strconv.ParseInt(in.Text(), 10, 64)
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
var allTime int64
var walkTime int64

func main() {
	defer exit()
	num = nextInt64()
	allTime = 4 * 60
	walkTime = nextInt64()

	if walkTime+5 > allTime {
		fmt.Fprintf(out, "%d\n", 0)
		return
	}

	l := int64(1)
	r := num
	for {
		if l == r {
			break
		}
		mid := int64((l + r + 1) / 2)
		midv := (allTime - 5*(mid+(mid)*(mid-1)/2) - walkTime) >= 0
		if midv {
			l = mid
		} else {
			r = mid - 1
		}
	}
	fmt.Fprintf(out, "%d\n", l)
}
