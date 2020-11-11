package main

import (
	"fmt"
)

func islandPerimeter(grid [][]int) int {
	hang := len(grid)
	if hang == 0 {
		return 0
	}
	lie := len(grid[0])
	var zong int
	for hangidx := range grid {
		for lieidx := range grid[hangidx] {
			if grid[hangidx][lieidx] == 1 {
				// shang xia bian jie
				if hangidx == 0 {
					zong++
					if hangidx == hang-1 {
						zong++
					} else {
						if grid[hangidx+1][lieidx] == 0 {
							zong++
						}
					}
				} else if hangidx == hang-1 {
					zong++
					if grid[hangidx-1][lieidx] == 0 {
						zong++
					}
				} else {
					if grid[hangidx-1][lieidx] == 0 {
						zong++
					}
					if grid[hangidx+1][lieidx] == 0 {
						zong++
					}
				}

				// zuo you bian jie
				if lieidx == 0 {
					zong++
					if lieidx == lie-1 {
						zong++
					} else {

						if grid[hangidx][lieidx+1] == 0 {
							zong++
						}
					}
				} else if lieidx == lie-1 {
					zong++
					if grid[hangidx][lieidx-1] == 0 {
						zong++
					}
				} else {
					if grid[hangidx][lieidx-1] == 0 {
						zong++
					}
					if grid[hangidx][lieidx+1] == 0 {
						zong++
					}
				}
			}

		}
	}
	return zong
}

func main() {
	one := make([]int, 1)

	one[0] = 1

	zong := make([][]int, 1)
	zong[0] = one

	a := islandPerimeter(zong)
	fmt.Println(a)
}
