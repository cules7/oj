package main

import "fmt"

// MAX32 is
var MAX32 int32 = 2147483647

// MIN32 is
var MIN32 int32 = -2147483648

// MAX64 is
var MAX64 int64 = 9223372036854775807

// MIN64 is
var MIN64 int64 = -9223372036854775808

func ri32(pn *int32) {
	fmt.Scanf("%d", pn)
}
func rai32(nums int32) []int32 {
	re := make([]int32, nums+1, nums+1)
	for idx := int32(1); idx != nums+1; idx++ {
		ri32(&re[idx])
	}
	return re
}

func ri64(pn *int64) {
	fmt.Scanf("%d", pn)
}
func rai64(nums int64) []int64 {
	re := make([]int64, nums+1, nums+1)
	for idx := int64(1); idx != nums+1; idx++ {
		ri64(&re[idx])
	}
	return re
}

func rstr(ps *string) {
	fmt.Scanf("%s", ps)
}
func rastr(nums int32) []string {
	re := make([]string, nums+1, nums+1)
	for idx := int32(1); idx != nums+1; idx++ {
		rstr(&re[idx])
	}
	return re
}

// OUT is
var OUT string

func oi32(i int32) {
	OUT += fmt.Sprintf("%d", i)
}
func oi64(i int64) {
	OUT += fmt.Sprintf("%d", i)
}

func oi32c(i int32, c string) {
	OUT += fmt.Sprintf("%d%s", i, c)
}
func oi64c(i int64, c string) {
	OUT += fmt.Sprintf("%d%s", i, c)
}

func oi32n(i int32) {
	oi32c(i, "\n")
}
func oi64n(i int64) {
	oi64c(i, "\n")
}

func oi32space(i int32) {
	oi32c(i, " ")
}
func oi64space(i int64) {
	oi64c(i, " ")
}

func ostring(s string) {
	OUT += s
}
func ob(b byte) {
	OUT += fmt.Sprintf("%c", b)
}

func p() {
	fmt.Print(OUT)
}

// 2 3 5 6
var N []int64

func min3(a, b, c int64) int64 {
	abxiao := int64(0)
	if N[a] <= N[b] {
		abxiao = N[a]
	} else {

		abxiao = N[b]
	}
	if abxiao <= N[c] {
		return abxiao
	}
	return N[c]
}
func main() {
	N = rai64(4)
	a := min3(1, 3, 4)
	var count int64
	count += (a * 256)
	sheng2 := N[1] - a
	if sheng2 <= N[2] {
		count += (sheng2 * 32)
		oi64n(count)
		p()
		return
	}
	count += (N[2] * 32)
	oi64n(count)
	p()
}
