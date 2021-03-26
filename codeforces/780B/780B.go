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

/*
 */
type siteAndV struct {
	s int
	v int
}
type sandv []siteAndV

func (s *sandv) Len() int {
	return len(*s)
}

var n int
var mth sandv

func zuixiaoTime(meetPos float64) float64 {
	var xiansuanyi float64
	if meetPos > float64(mth[1].s) {
		xiansuanyi = (meetPos - float64(mth[1].s)) / float64(mth[1].v)
	} else {
		xiansuanyi = (-meetPos + float64(mth[1].s)) / float64(mth[1].v)
	}
	for i := 2; i != n+1; i++ {
		var zan float64
		if meetPos > float64(mth[i].s) {
			zan = (meetPos - float64(mth[i].s)) / float64(mth[i].v)
		} else {
			zan = (-meetPos + float64(mth[i].s)) / float64(mth[i].v)
		}
		if zan > xiansuanyi {
			xiansuanyi = zan
		}
	}
	return xiansuanyi
}
func main() {
	defer exit()

	n = nextInt()
	mth = make(sandv, n+1, n+1)

	var left = float64(1)
	var right = float64(1000000000)

	for i := 1; i != n+1; i++ {
		mth[i].s = nextInt()
	}

	for i := 1; i != n+1; i++ {
		mth[i].v = nextInt()
	}

	for {
		if right-left < 0.000001 {
			break
		}
		sanfenyi := left + (right-left)/3
		erfenyi := left + (right-left)/2
		s := zuixiaoTime(sanfenyi)
		e := zuixiaoTime(erfenyi)
		if s < e {
			right = erfenyi
		} else {
			left = sanfenyi
		}
	}
	fmt.Fprintf(out, "%f\n", zuixiaoTime(left))
}
