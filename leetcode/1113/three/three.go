package main

func lianghemap(a []int, b []int) map[int64]uint16 {
	mzheng := make(map[int64]uint16)
	for _, v := range a {
		for _, w := range b {
			var he = int64(v + w)

			if _, ok := mzheng[he]; ok == true {
				mzheng[he]++
			} else {
				mzheng[he] = 1
			}

		}
	}
	return mzheng
}

func fourSumCount(A []int, B []int, C []int, D []int) int {
	m1 := lianghemap(A, B)
	var zong int
	for _, v := range C {
		for _, w := range D {
			var he = int64(v + w)
			if vv, ok := m1[-he]; ok == true {
				zong += int(vv)
			}
		}
	}

	return zong
}
func main() {

}
