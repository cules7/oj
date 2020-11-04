package main

func repeatedSubstringPattern(str string) bool {
	l := len(str)
	if l == 0 {
		return false
	}
	if l == 1 {
		return false
	}
	ridx := 1
	var ok bool
	for {
		if ridx == l/2+1 {
			return false
		}
		substr := str[0:ridx]
		if l%ridx != 0 {
			ridx++
			continue
		}
		ok = true
		for idx := range str {
			if str[idx] != substr[idx%ridx] {
				ok = false
				break
			}
		}
		if ok == true {
			return ok
		}
		ridx++
	}
}
func main() {

}
