package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// MAX32 is
var MAX32 int32 = 2147483647

// MIN32 is
var MIN32 int32 = -2147483648

// MAX64 is
var MAX64 int64 = 9223372036854775807

// MIN64 is
var MIN64 int64 = -9223372036854775808

var scanner = bufio.NewScanner(os.Stdin)

var o = bufio.NewWriter(os.Stdout)

func init() {
	// scanner.Buffer()
	scanner.Split(bufio.ScanWords)
}
func end() {
	o.Flush()
}

var n int64
var m int32
var k int32
var x int64
var s int64
var ai []int64
var bi []int64
var ci []int64
var di []int64

func main() {
	defer end()
	n = ri64()             //task
	m = ri32()             // type 1 num
	k = ri32()             // type 2 num
	x = ri64()             // time
	s = ri64()             // mp
	ai = ri64sli(int64(m)) // type 1 xiaoguo
	bi = ri64sli(int64(m)) // type 1 mp
	ci = ri64sli(int64(k)) // type 2 xiaoguo
	di = ri64sli(int64(k)) //type 2 mp
	var last = MAX64
	for idx := int32(1); idx != m+2; idx++ {
		var mpleft int64
		var singletime int64
		if idx == m+1 || s < bi[idx] {
			mpleft = s
			singletime = x
		} else {
			mpleft = s - bi[idx]
			singletime = ai[idx]
		}
		in := bsi64(di, mpleft)
		lefttask := n - ci[in]
		wholetime := int64(singletime) * int64(lefttask)
		if wholetime < last {
			last = wholetime
		}
	}
	oi(last)
	// if n == 1997320486 {
	// 	ostring(" ")
	// 	oi(int64(lastidx))
	// 	ostring(" ")
	// 	oi(lastin)
	// }
	ostring("\n")
}

// IO
func ri32() int32 {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	n32 := int32(n)
	return n32
}
func ri32sli(nums int32) []int32 {
	re := make([]int32, nums+1, nums+1)
	for idx := int32(1); idx != nums+1; idx++ {
		re[idx] = ri32()
	}
	return re
}
func ri64() int64 {
	scanner.Scan()
	n64, _ := strconv.ParseInt(scanner.Text(), 0, 64)
	return n64
}
func ri64sli(nums int64) []int64 {
	re := make([]int64, nums+1, nums+1)
	for idx := int64(1); idx != nums+1; idx++ {
		re[idx] = ri64()
	}
	return re
}

func oi(i int64) {
	o.WriteString(fmt.Sprintf("%d", i))
}
func ostring(s string) {
	o.WriteString(s)
}

// math
func bsi64(s []int64, k int64) int64 {
	i := int64(0)
	j := int64(len(s) - 1)
	for i < j {
		m := (i + j + 1) / 2
		if s[m] <= k {
			i = m
		} else {
			j = m - 1
		}
	}
	return i
}
