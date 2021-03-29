package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var in = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func init() {
	in.Split(bufio.ScanWords)
	buf := make([]byte, 2500000, 2500000)
	in.Buffer(buf, 0)
}
func exit() {
	out.Flush()
}

// custom
func nextInt() int {
	in.Scan()
	ret, _ := strconv.Atoi(in.Text())
	return ret
}

func nextString() string {
	in.Scan()
	return in.Text()
}
func nextBytes() []byte {
	in.Scan()
	return in.Bytes()
}

/*
 */
var source []byte
var target []byte
var asource []byte
var order []int

func containsSubsequence(exclude map[int]struct{}) int {
	j := 0
	for i := 0; i < len(source); i++ {
		if _, ok := exclude[i+1]; ok {
			continue
		}
		if source[i] == target[j] {
			j++
			if j == len(target) {
				return 1
			}
		}
	}
	return 2
}

func main() {
	defer exit()
	source = nextBytes()
	target = nextBytes()
	asource = make([]byte, len(source), len(source))
	copy(asource, source)
	order = make([]int, len(source)+1, len(source)+1)
	for idx := 1; idx != len(order); idx++ {
		order[idx] = nextInt()
	}
	l := 0
	r := len(source) - len(target)

	for {
		if l == r {
			break
		}
		mid := (l + r + 1) / 2
		ex := make(map[int]struct{})
		for idx := 1; idx != mid+1; idx++ {
			ex[order[idx]] = struct{}{}
		}
		midv := containsSubsequence(ex)
		if midv == 1 {
			l = mid
		} else {
			r = mid - 1
		}
	}
	fmt.Fprintf(out, "%d\n", l)
}
