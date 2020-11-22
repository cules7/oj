package main

import "strings"
import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func ans(arr string, n int) int {
	if n == 1 || n == 0 {
		return 0
	}
	lenofarr := len(arr)
	if lenofarr < n {
		return -1 // fuck
	} else if lenofarr < 2*n {
		sum := strings.Count(arr, "1")
		toone := lenofarr - sum
		tozero := sum
		return min(toone, tozero)
	}

	n_1sum := make([]int, lenofarr+1)
	initsum := strings.Count(arr[n:2*n-1], "1")
	n_1sum[2*n-1] = initsum
	for i := (2 * n); i != lenofarr+1; i++ {
		if arr[i-1] == '1' {
			initsum++
		}
		if arr[i-n] == '1' {
			initsum--
		}
		n_1sum[i] = initsum
	}

	frontnsum := strings.Count(arr[0:n], "1")
	N0 := make([]int, lenofarr+1)
	N1 := make([]int, lenofarr+1)
	if arr[n] == '1' {
		N0[n+1] = 1 + frontnsum
		N1[n+1] = n - frontnsum
	} else {
		N0[n+1] = frontnsum
		N1[n+1] = 1 + n - frontnsum
	}
	for i := n + 2; i != 2*n; i++ {
		if arr[i-1] == '1' {
			N0[i] = 1 + N0[i-1]
			N1[i] = N1[i-1]
		} else {
			N0[i] = N0[i-1]
			N1[i] = 1 + N1[i-1]
		}
	}

	B0 := make([]int, lenofarr+1)
	B1 := make([]int, lenofarr+1)

	minfront := min(frontnsum, n-frontnsum)

	B0[2*n-1] = n_1sum[2*n-1] + minfront
	B1[2*n-1] = n - n_1sum[2*n-1] - 1 + minfront

	for i := 2 * n; i != lenofarr+1; i++ {
		if arr[i-1] == '1' {
			N0[i] = min(N0[i-1], B0[i-1]) + 1
			N1[i] = min(N1[i-1], B1[i-1])
		} else {
			N0[i] = min(N0[i-1], B0[i-1])
			N1[i] = min(N1[i-1], B1[i-1]) + 1
		}
		B0[i] = min(N0[i-n+1], N1[i-n+1]) + n_1sum[i]
		B1[i] = min(N0[i-n+1], N1[i-n+1]) + n - n_1sum[i] - 1
	}
	// fmt.Println(n_1sum)
	// fmt.Println(N0)
	// fmt.Println(N1)
	// fmt.Println(B0)
	// fmt.Println(B1)
	return min(N0[lenofarr], N1[lenofarr])
}
func reverseString(s string) string {
	runes := []rune(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}

func main() {
	s := "10100010"
	n := 3
	a := ans(s, n)
	b := ans(reverseString(s), n)
	fmt.Println(a, b)

	s = "10100010"
	n = 2
	a = ans(s, n)
	b = ans(reverseString(s), n)
	fmt.Println(a, b)

}
