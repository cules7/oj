package main

import (
	"bufio"
	"fmt"
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

var lpos *big.Rat
var rpos *big.Rat
var leftnum *big.Rat

var p_meet_b_t *big.Rat
var p_move *big.Rat
var l_now_pos *big.Rat
var r_now_pos *big.Rat
var b_zhui_t *big.Rat
var zong_t *big.Rat

var hehe *big.Rat

func cost(pos *big.Rat) *big.Rat {
	// the number of the left pupils
	leftnum.Sub(n, rongliang)
	// when bus stop to go back, where are the left pupils
	lpos.Quo(pos, vb).Mul(lpos, vp)
	// when bus stop to go back, where are the in-bus pupils
	rpos.Set(pos)
	for leftnum.Sign() > 0 {
		p_meet_b_t.Sub(rpos, lpos).Quo(p_meet_b_t, vhebp)
		p_move.Mul(p_meet_b_t, vp)
		l_now_pos.Add(lpos, p_move)
		r_now_pos.Add(rpos, p_move)
		b_zhui_t.Sub(r_now_pos, l_now_pos).Quo(b_zhui_t, vchabp)
		zong_t.Add(p_meet_b_t, b_zhui_t)
		lpos.Add(lpos, hehe.Mul(vp, zong_t))
		rpos.Add(rpos, hehe.Mul(vp, zong_t))
		leftnum.Sub(leftnum, rongliang)
	}
	return hehe.Sub(l, rpos).Abs(hehe)
}
func answer(pos float64) float64 {
	var zongt = pos / float64(vb)
	// the number of the left pupils
	leftnum := n - rongliang
	// when bus stop to go back, where are the left pupils
	lpos := (pos / float64(vb)) * float64(vp)
	// when bus stop to go back, where are the in-bus pupils
	rpos := pos
	for leftnum > 0 {
		p_meet_b_t := (rpos - lpos) / float64(vp+vb)
		p_move := p_meet_b_t * float64(vp)
		l_now_pos := lpos + p_move
		r_now_pos := rpos + p_move
		b_zhui_t := (r_now_pos - l_now_pos) / float64(vb-vp)
		zong_t := p_meet_b_t + b_zhui_t
		lpos = lpos + float64(vp)*zong_t
		rpos = rpos + float64(vp)*zong_t
		leftnum -= rongliang
		zongt += zong_t
	}

	return zongt
}
func main() {
	defer exit()

	n = big.NewRat(int64(nextInt()), 1)
	l = big.NewRat(int64(nextInt()), 1)
	vp = big.NewRat(int64(nextInt()), 1)
	vb = big.NewRat(int64(nextInt()), 1)
	rongliang = big.NewRat(int64(nextInt()), 1)
	hehe = big.NewRat(0, 1)
	hehe.Add(vp, vb)
	vhebp.Set(hehe)
	hehe.Sub(vb, vp)
	vchabp.Set(hehe)
	if rongliang.Cmp(n) >= 0 {
		fmt.Fprintf(out, "%s\n", l.Quo(l, vb).FloatString(10))
		return
	}
	var cp = big.NewRat(1, 1000000000)
	var left = big.NewRat(1, 100000000)
	var right = big.NewRat(0, 1)
	right.Set(l)
	var diff = big.NewRat(0, 1)
	var s_1_3 = big.NewRat(0, 1)
	var s_1_2 = big.NewRat(0, 1)
	var san = big.NewRat(3, 1)
	var er = big.NewRat(2, 1)
	for {
		fmt.Fprintf(out, "%s %s\n", left.FloatString(20), right.FloatString(20))
		if diff.Sub(right, left).Cmp(cp) < 0 {
			break
		}
		s_1_3.Quo(diff, san)
		s_1_3.Add(left, s_1_3)

		s_1_2.Quo(diff, er)
		s_1_2.Add(left, s_1_2)
		sanf, _ := s_1_3.Float64()
		erf, _ := s_1_2.Float64()
		var c_s_1_3 = cost(sanf)
		var c_s_1_2 = cost(erf)
		if c_s_1_3 < c_s_1_2 {
			right.Set(s_1_2)
		} else {
			left.Set(s_1_3)
		}
	}
	l_pos, _ := left.Float64()
	t := answer(l_pos)
	fmt.Fprintf(out, "%.10f\n", t)
}
