package main

import (
	"fmt"
)

func qiansanxing(nums []int) bool {
	yi := nums[0]
	er := nums[1]
	san := nums[2]
	if san > yi && san < er {
		return true
	}
	return false
}

// xiao da
func findG(max, min, xin int) (int, int) {
	if max == min {
		if xin == max {
			return max, min
		}
		if xin < max {
			return xin, xin
		}
		if xin > max {
			return max, xin
		}
	}
	if max > min {
		if xin == min {
			return min, max
		}
		if xin == max {
			return min, max
		}
		if xin > max {
			return min, xin
		}
		if xin < min {
			return min, max
		}
	}
	return min, max
}
func findmm(nums []int) (min, max int) {
	yi := nums[0]
	er := nums[1]
	san := nums[2]
	if yi >= er {
		max = er
		min = er
	} else {
		min = yi
		max = er
	}
	min, max = findG(max, min, san)
	return min, max
}
func find132pattern(nums []int) bool {
	l := len(nums)
	if l < 3 {
		return false
	}
	if qiansanxing(nums) == true {
		return true
	}
	min, max := findmm(nums)
	idx := 3
	for {
		if idx == l {
			return false
		}
		v := nums[idx]
		if v < max && v > min {
			return true
		}
		min, max = findG(max, min, v)
		idx++
	}
}
func main() {
	min, max := findG(1, 0, -4)
	fmt.Println(min, max)
}
