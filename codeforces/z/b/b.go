package main

import (
	"bufio"
	"fmt"
	"os"
)

var in = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func init() {
	in.Split(bufio.ScanWords)
}
func exit() {
	out.Flush()
}
func nextBytes() []byte {
	in.Scan()
	return in.Bytes()
}
func nextString() string {
	in.Scan()
	return in.Text()
}
func main() {
	defer exit()
	fmt.Fprintf(out, "%v \n", nextBytes())
}
