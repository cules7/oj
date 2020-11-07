package main

import "strings"

// MAX32 is
var MAX32 int32 = 2147483647

// MIN32 is
var MIN32 int32 = -2147483648

// MAX64 is
var MAX64 int64 = 9223372036854775807

// MIN64 is
var MIN64 int64 = -9223372036854775808

//
var game string
var N int32

func main() {
	ri32(&N)
	rstr(&game)

	A := int32(strings.Count(game, "A"))
	D := N - A
	if A == D {
		ostring("Friendship\n")
		p()
		return
	}
	if A < D {
		ostring("Danik\n")
		p()
		return
	}
	if A > D {
		ostring("Anton\n")
		p()
		return
	}
}
