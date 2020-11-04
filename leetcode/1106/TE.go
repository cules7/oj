package main

import (
	"fmt"
	"sort"
)

type PX [][]int

func (px *PX) Len() int {
	return len(*px)
}
func (px *PX) Less(i, j int) bool {
	if (*px)[i][0] < (*px)[j][0] {
		return true
	}
	return false
}
func (px *PX) Swap(i, j int) {
	(*px)[i], (*px)[j] = (*px)[j], (*px)[i]
}
func main() {
	a := make([][]int, 3)
	a[0] = make([]int, 1)
	a[1] = make([]int, 1)
	a[2] = make([]int, 1)
	a[0][0] = 4
	a[1][0] = 3
	a[2][0] = 2
	var px = (*PX)(&a)
	sort.Sort(px)
	fmt.Println(a)
}

func haveCommon(a, b []int) bool {
	if a[1] >= b[0] {
		return true
	}
	return false
}

func findMinArrowShots(points [][]int) int {
	lenofwb := len(points)
	if lenofwb == 1 {
		return 1
	}
	var pp = (*PX)(&points)
	sort.Sort(pp)
	var dp = 1
	var needbi = 0
	for idx := 1; idx != lenofwb; idx++ {
		if haveCommon(points[idx], points[needbi]) == true {
			if points[idx][1] < points[needbi][1] {
				needbi = idx
			}
			continue
		} else {
			dp++
			needbi = idx
		}
	}
	return dp
}
