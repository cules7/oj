package main

import (
	"sort"
)

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

func findContentChildren(g []int, s []int) int {
	if len(g) == 0 {
		return 0
	}
	if len(s) == 0 {
		return 0
	}
	var pg = (*PX)(&g)
	var ps = (*PX)(&s)
	sort.Sort(pg)
	sort.Sort(ps)
	xiaohaiinit := 0
	var count int
	for igx := range s {
		var con bool
		for xiaohai := xiaohaiinit; xiaohai != len(g); {
			if s[igx] >= g[xiaohai] {
				count++
				xiaohaiinit = xiaohai + 1
				con = true
				break
			}
			xiaohai++
		}
		if con == true {
			continue
		}
		return count
	}
	return count
}

func main() {

}
