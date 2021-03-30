package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

// io
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
var nums int
var fear int

type weapon struct {
	crary int
	force int
}
type weapons []weapon

var w weapons

func (wps weapons) Len() int {
	return len(wps)
}
func (wps weapons) Less(i, j int) bool {
	return wps[i].crary < wps[j].crary
}
func (wps weapons) Swap(i, j int) {
	wps[i], wps[j] = wps[j], wps[i]
}
func main() {
	defer exit()
	nums = nextInt()
	fear = nextInt()
	w = make(weapons, nums+1, nums+1)
	for i := 1; i != nums+1; i++ {
		w[i].crary = nextInt()
		w[i].force = nextInt()
	}
	sort.Sort(w[1:])
	lastLeft := -1
	myLastForce := int64(-1)
	leftForce := int64(0)
	for i := 1; i != nums+1; i++ {
		l := i
		r := nums
		for {
			if l == r {
				break
			}
			mid := (l + r + 1) / 2
			midv := (w[mid].crary - w[i].crary - fear) < 0
			if midv {
				l = mid
			} else {
				r = mid - 1
			}
		}
		var tempForce int64
		if l > lastLeft {
			tempForce = leftForce - int64(w[i-1].force)
			for j := lastLeft + 1; j != l+1; j++ {
				tempForce += int64(w[j].force)
			}
			leftForce = tempForce
		} else {
			tempForce = leftForce - int64(w[i-1].force)
			leftForce = tempForce
		}
		lastLeft = l
		if myLastForce < leftForce {
			myLastForce = leftForce
		}
	}
	fmt.Fprintf(out, "%d\n", myLastForce)
}
