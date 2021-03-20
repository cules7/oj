package main

import (
	"fmt"
)

/*
	1223
*/
func findMinimum(n []int, target int) int {
	l := 0
	r := len(n) - 1
	for {
		mid := (l + r) / 2
		midValue := n[mid]
		if midValue < target {
			l = mid + 1
		} else {
			r = mid
		}
		if r-l+1 <= 1 {
			break
		}
	}
	return r
}
func main() {
	test := []int{1, 3, 3, 3, 5}
	fmt.Println(findMinimum(test, 2))
}
