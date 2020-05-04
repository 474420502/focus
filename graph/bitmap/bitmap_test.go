package bitmap

import (
	"encoding/binary"
	"testing"
)

func TestCase1(t *testing.T) {
	dimX, dimY := 3, 3
	bm := NewBitmap(dimX, dimY)
	for y := 0; y < dimY; y++ {
		for x := 0; x < dimX; x++ {

			bm.SetBit(dimX, y, x, 1)
		}
	}

	if len(bm.bits) != 2 {
		t.Error(len(bm.bits))
	}

	for y := 0; y < dimY; y++ {
		for x := 0; x < dimX; x++ {
			if bm.GetBit(dimX, y, x) != 1 {
				t.Error("y:", y, "x:", x, "value !=", 1)
			}

			bm.SetBit(dimX, y, x, 0)

			if bm.GetBit(dimX, y, x) != 0 {
				t.Error("y:", y, "x:", x, "value !=", 1)
			}
		}
	}
}

func TestCase2(t *testing.T) {
	dimX, dimY := 4, 4
	bm := NewBitmap(dimX, dimY)
	for y := 0; y < dimY; y++ {
		for x := 0; x < dimX; x++ {
			msize := y*dimX + x
			bm.SetBitBySize(msize, 1)
		}
	}

	if len(bm.bits) != 2 {
		t.Error(len(bm.bits))
	}

	for y := 0; y < dimY; y++ {
		for x := 0; x < dimX; x++ {
			msize := y*dimX + x
			if bm.GetBitBySize(msize) != 1 {
				t.Error("y:", y, "x:", x, "value !=", 1)
			}

			bm.SetBitBySize(msize, 0)

			if bm.GetBitBySize(msize) != 0 {
				t.Error("y:", y, "x:", x, "value !=", 1)
			}

		}
	}
}

func TestCase3(t *testing.T) {
	dimX, dimY := 8, 8
	bm := NewBitmap(dimX, dimY)

	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, ^uint64(0))
	bm.SetBytes(0, buf)

	if binary.BigEndian.Uint64(bm.GetBytes()) != ^uint64(0) {
		t.Error("SetBytes GetBytes error")
	}

	for y := 0; y < dimY; y++ {
		for x := 0; x < dimX; x++ {
			msize := y*dimX + x
			if bm.GetBitBySize(msize) != 1 {
				t.Error("y:", y, "x:", x, "value !=", 1)
			}
		}
	}

	nbm := bm.Clone()
	if string(nbm.GetBytes()) != string(bm.GetBytes()) {
		t.Error("Clone error")
	}
}
