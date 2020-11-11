package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const MAX_MK = 200010

var N int64
var M, K int
var X, S int64
var A [MAX_MK]int64
var B [MAX_MK]int64
var C [MAX_MK]int64
var D [MAX_MK]int64

func lowerBound(mana int64) int64 {
	var left, right int
	left = -10
	right = K + 10
	var result int64
	result = 0
	for right-left > 1 {
		med := (right + left) / 2
		//fmt.Println("med", med)
		if med < 0 {
			left = med
			continue
		}
		if med > K-1 {
			right = med
			continue
		}
		reqMana := D[med]
		if reqMana > mana {
			right = med
		} else {
			left = med
			result = C[med]
		}
	}
	return result
}

func solve() {
	var result int64
	result = 1 << 62
	for i := 0; i < int(M+1); i++ {
		var pre, mana int64
		if i < int(M) && B[i] <= S {
			pre = A[i]
			mana = S - B[i]
		} else {
			pre = X
			mana = S
		}
		var needPotions int64
		needPotions = N
		needPotions -= lowerBound(mana)
		if needPotions <= 0 {
			result = 0
			break
		}

		if needPotions*pre < result {
			lastidx = 
			result = needPotions * pre
		}
		//fmt.Println(i, needPotions, pre)
		//fmt.Println(i, needPotions*pre)
	}
	fmt.Println(result)
}

func main() {
	fsc := NewFastScanner()
	N, M, K = fsc.NextInt64(), fsc.NextInt(), fsc.NextInt()
	X, S = fsc.NextInt64(), fsc.NextInt64()
	for i := 0; i < M; i++ {
		A[i] = fsc.NextInt64()
	}
	for i := 0; i < M; i++ {
		B[i] = fsc.NextInt64()
	}

	for i := 0; i < K; i++ {
		C[i] = fsc.NextInt64()
	}
	for i := 0; i < K; i++ {
		D[i] = fsc.NextInt64()
	}
	solve()
}

//template
type FastScanner struct {
	r   *bufio.Reader
	buf []byte
	p   int
}

func NewFastScanner() *FastScanner {
	rdr := bufio.NewReaderSize(os.Stdin, 2028)
	return &FastScanner{r: rdr}
}

func (s *FastScanner) Next() string {
	s.Pre()
	start := s.p
	for ; s.p < len(s.buf); s.p++ {
		if s.buf[s.p] == ' ' {
			break
		}
	}
	result := string(s.buf[start:s.p])
	s.p++
	return result
}

func (s *FastScanner) NextInt() int {
	val, err := strconv.Atoi(s.Next())
	if err != nil {
		panic(err)
	}
	return val
}
func (s *FastScanner) NextInt64() int64 {
	val, err := strconv.ParseInt(s.Next(), 10, 64)
	if err != nil {
		panic(err)
	}
	return val
}
func (s *FastScanner) Pre() {
	if s.p >= len(s.buf) {
		s.ReadLine()
		s.p = 0
	}
}

func (s *FastScanner) ReadLine() {
	s.buf = make([]byte, 0)
	for {
		l, p, err := s.r.ReadLine()
		if err != nil {
			panic(err.Error())
		}
		s.buf = append(s.buf, l...)
		if !p {
			break
		}
	}
}
