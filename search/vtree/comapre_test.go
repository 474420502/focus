package treelist

import (
	"bytes"
	"strconv"
	"testing"
)

var compsize = 500000

func BenchmarkCase1(t *testing.B) {
	var l [][]byte
	for i := 0; i < compsize; i++ {
		istr := []byte(strconv.Itoa(i))
		l = append(l, istr)
		if len(l) > 500 {
			for _, x := range l[len(l)-100:] {
				if bytes.Compare(istr, x) > 0 {
					istr = nil
				}
			}
		}

	}
}

func BenchmarkCase2(t *testing.B) {
	var l [][]byte

	for i := 0; i < compsize; i++ {
		istr := []byte(strconv.Itoa(i))
		l = append(l, istr)
		if len(l) > 500 {
			for _, x := range l[len(l)-100:] {
				if CompatorByte(istr, x) > 0 {
					istr = nil
				}
			}
		}
	}
}
