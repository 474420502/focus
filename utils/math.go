package utils

var bit = uint(32 << (^uint(0) >> 63))
var bitsub1 = bit - 1

// AbsInt
func AbsInt(n int) uint {
	y := n >> bitsub1
	return uint((n ^ y) - y)
}
