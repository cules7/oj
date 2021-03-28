package main

import (
	"bufio"
	"fmt"
	"math"
	"math/big"
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

var n *big.Rat
var l *big.Rat
var vp *big.Rat
var vb *big.Rat
var rongliang *big.Rat
var vhebp *big.Rat
var vchabp *big.Rat
var hehe *big.Rat
var haha *big.Rat
var NTime *big.Rat
var initdiff *big.Rat
var lpos *big.Rat
var inittime *big.Rat
var onetime *big.Rat
var backtime *big.Rat
var chasetime *big.Rat

func initAll() {
	vhebp = big.NewRat(0, 1)
	vchabp = big.NewRat(0, 1)
	hehe = big.NewRat(0, 1)
	haha = big.NewRat(0, 1)
	initdiff = big.NewRat(0, 1)
	NTime = big.NewRat(0, 1)
	lpos = big.NewRat(0, 1)
	inittime = big.NewRat(0, 1)
	onetime = big.NewRat(0, 1)
	backtime = big.NewRat(0, 1)
	chasetime = big.NewRat(0, 1)

	hehe.Add(vp, vb)
	vhebp.Set(hehe)
	hehe.Sub(vb, vp)
	vchabp.Set(hehe)

	var ntime, _ = hehe.Sub(n, rongliang).Quo(hehe, rongliang).Float64()
	NTime.SetFloat64(math.Ceil(ntime))
}
func cost(pos *big.Rat) *big.Rat {
	inittime.Quo(pos, vb)
	lpos.Mul(inittime, vp)
	initdiff.Sub(pos, lpos)
	backtime.Quo(initdiff, vhebp)
	hehe.Mul(backtime, vb)
	haha.Mul(backtime, vp)
	hehe.Add(hehe, haha)
	chasetime.Quo(hehe, vchabp)
	onetime.Add(backtime, chasetime)
	hehe.Mul(onetime, NTime)
	hehe.Mul(vp, hehe)
	hehe.Add(hehe, pos)
	hehe.Sub(hehe, l)
	return hehe.Abs(hehe)
}
func answer(pos *big.Rat) *big.Rat {
	inittime.Quo(pos, vb)
	lpos.Mul(inittime, vp)
	initdiff.Sub(pos, lpos)
	backtime.Quo(initdiff, vhebp)
	hehe.Mul(backtime, vb)
	haha.Mul(backtime, vp)
	hehe.Add(hehe, haha)
	chasetime.Quo(hehe, vchabp)
	onetime.Add(backtime, chasetime)
	hehe.Mul(onetime, NTime)
	return hehe.Add(inittime, hehe)
}
func main() {
	defer exit()

	n = big.NewRat(int64(nextInt()), 1)
	l = big.NewRat(int64(nextInt()), 1)
	vp = big.NewRat(int64(nextInt()), 1)
	vb = big.NewRat(int64(nextInt()), 1)
	rongliang = big.NewRat(int64(nextInt()), 1)
	initAll()

	if rongliang.Cmp(n) >= 0 {
		fmt.Fprintf(out, "%s\n", l.Quo(l, vb).FloatString(10))
		return
	}
	var cp = big.NewRat(1, 10000000000)
	var left = big.NewRat(1, 1000000000)
	var right = big.NewRat(0, 1)
	right.Set(l)
	var diff = big.NewRat(0, 1)
	var s_1_3 = big.NewRat(0, 1)
	var c13 = big.NewRat(0, 1)
	var s_1_2 = big.NewRat(0, 1)
	var c12 = big.NewRat(0, 1)
	var san = big.NewRat(3, 1)
	var er = big.NewRat(2, 1)
	for {
		if diff.Sub(right, left).Cmp(cp) < 0 {
			break
		}
		s_1_3.Quo(diff, san)
		s_1_3.Add(left, s_1_3)

		s_1_2.Quo(diff, er)
		s_1_2.Add(left, s_1_2)
		c13.Set(cost(s_1_3))
		c12.Set(cost(s_1_2))

		if c13.Cmp(c12) > 0 {
			left.Set(s_1_3)
		} else {
			right.Set(s_1_2)
		}
	}
	t := answer(left)
	fmt.Fprintf(out, "%s\n", t.FloatString(10))
}
