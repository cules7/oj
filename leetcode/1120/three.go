package main

import (
	"fmt"
)

var nodeNP []byte
var max int
var d int

func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	if desiredTotal == 0 {
		return true
	}
	var sum int
	for idx := 1; idx != maxChoosableInteger+1; idx++ {
		sum += idx
	}
	if sum <= desiredTotal {
		return false
	}
	max = maxChoosableInteger
	d = desiredTotal
	nodeNP = make([]byte, 1<<uint(maxChoosableInteger)) // 'N' 'P' '\0'
	dp(0)
	if nodeNP[0] == 'N' {
		return true
	}
	return false
}

func dp(node int) {
	if nodeNP[node] != 0 {
		return
	}
	// use d to tag 'P' as terminal
	var child []int
	var zong int
	for i := 0; i != max; i++ {
		yiwei := node >> uint(i)
		if (yiwei & 1) == 0 { // no use
			child = append(child, i)
		} else { //use
			zong += (i + 1)
		}
		if zong >= d {
			nodeNP[node] = 'P'
			return
		}
	}
	// scan child
	var last byte = 'P'
	for _, v := range child {
		cnode := node + 1<<uint(v)
		dp(cnode)
		cNP := nodeNP[cnode]
		if cNP == 'P' {
			last = 'N'
		}
	}
	nodeNP[node] = last
}
func main() {
	mmm := 5
	ddd := 50
	b := canIWin(mmm, ddd)
	fmt.Println(mmm, ddd, b)
}
