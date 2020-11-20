package main

func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	sum := ((maxChoosableInteger + 1) * (maxChoosableInteger)) / 2
	b := 0
	one := 1
	two := 0
	who := 1
	for {
		if who == 1 {
			twon := 0
			// two
			for idx := 0; idx != two; idx++ {
				twon += (idx + 1)
			}
			onen := 0
			for idx := 0; idx != one; idx++ {
				onen += (maxChoosableInteger - idx)
			}
			N := twon + onen
			if N > sum {
				return false
			}
			if desiredTotal <= N && desiredTotal >= b {
				return true
			}
			b = N + 1
			two++
			who = 2
		} else {
			onen := 0
			for idx := 0; idx != one; idx++ {
				onen += (idx + 1)
			}
			twon := 0
			for idx := 0; idx != two; idx++ {
				twon += (maxChoosableInteger - idx)
			}
			N := onen + twon
			if N > sum {
				return false
			}
			if desiredTotal >= b && desiredTotal <= N {
				return false
			}
			b = N + 1
			one++
			who = 1
		}
	}
}
