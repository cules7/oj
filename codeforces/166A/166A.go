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
func nextInt64() int64 {
	in.Scan()
	ret, _ := strconv.ParseInt(in.Text(), 10, 64)
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

type chengji struct {
	fenshu  int
	shijian int
}
type teams []chengji

func (ts teams) Len() int {
	return len(ts)
}
func (ts teams) Less(i, j int) bool {
	if ts[i].fenshu < ts[j].fenshu {
		return false
	}
	if ts[i].fenshu > ts[j].fenshu {
		return true
	}
	return ts[i].shijian < ts[j].shijian
}
func (ts teams) Swap(i, j int) {
	ts[i], ts[j] = ts[j], ts[i]
}

var num int
var kk int
var liushui teams

func main() {
	defer exit()
	num = nextInt()
	kk = nextInt()
	liushui = make(teams, num+1, num+1)
	for i := 1; i != num+1; i++ {
		liushui[i].fenshu = nextInt()
		liushui[i].shijian = nextInt()
	}
	sort.Sort(liushui[1:])
	// two direction binary search begin
	l := kk
	r := num
	for {
		if l == r {
			break
		}
		mid := (l + r + 1) / 2
		midv := (liushui[mid].fenshu == liushui[kk].fenshu) && (liushui[mid].shijian == liushui[kk].shijian)
		if midv {
			l = mid
		} else {
			r = mid - 1
		}
	}
	up := l

	l = 1
	r = kk
	for {
		if l == r {
			break
		}
		mid := (l + r) / 2
		midv := (liushui[mid].fenshu == liushui[kk].fenshu) && (liushui[mid].shijian == liushui[kk].shijian)
		if midv {
			r = mid
		} else {
			l = mid + 1
		}
	}
	down := l
	fmt.Fprintf(out, "%d\n", up-down+1)
}
