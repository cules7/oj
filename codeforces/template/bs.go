package main

func bsi32(s []int32, i, j, k int32) int32 {
	for {
		if i == j {
			if k <= s[i] {
				return i
			}
			if k > s[i] {
				return i + 1
			}
		}
		centeri := (i + j) / 2
		centerv := s[centeri]
		if k == centerv {
			return centeri
		}
		if k > centerv {
			i = centeri + 1
			continue
		}
		if k < centerv {
			j = centeri - 1
			if j < i {
				j = i
			}
			continue
		}
	}
}

func bsi64(s []int64, k int64) int64 {
	i := int64(0)
	j := int64(len(s) - 1)
	for i < j {
		m := (i + j + 1) / 2
		if s[m] > k {
			j = m - 1
		} else if s[m] <= k {
			i = m
		}
	}
	return i
}
