package main

import "sort"

type PX []int

func (px *PX) Len() int {
	return len(*px)
}
func (px *PX) Less(i, j int) bool {
	if (*px)[i] > (*px)[j] {
		return true
	}
	return false
}
func (px *PX) Swap(i, j int) {
	(*px)[i], (*px)[j] = (*px)[j], (*px)[i]
}
func abs(a, b int) int {
	c := a - b
	if c >= 0 {
		return c
	}
	return 0 - c
}
func minMoves2(nums []int) int {
	l := len(nums)
	if l == 1 {
		return 0
	}
	sort.Sort((*PX)(&nums))
	if l%2 == 0 {
		zong := 0
		tar := nums[l/2]
		for _, v := range nums {
			zong += abs(tar, v)
		}

		zanzong := 0
		tar = nums[l/2-1]
		for _, v := range nums {
			zanzong += abs(tar, v)
		}
		if zong <= zanzong {
			return zong
		}
		return zanzong
	}

	tar := nums[l/2]
	zong := 0
	for _, v := range nums {
		zong += abs(tar, v)
	}
	return zong
}
