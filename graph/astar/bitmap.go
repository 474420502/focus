package astar

func NewBitmap2D(dx, dy int) *Bitmap2D {
	bm := &Bitmap2D{dimX: dx, dimY: dy}
	bm.bits = make([]byte, (bm.dimX*bm.dimY+7)/8)
	return bm
}

func CopyFrom(other *Bitmap2D) *Bitmap2D {
	bm := &Bitmap2D{dimX: other.dimX, dimY: other.dimY}
	bm.bits = make([]byte, len(other.bits))
	copy(bm.bits, other.bits)
	return bm
}

type Bitmap2D struct {
	dimX, dimY int
	bits       []byte
}

func (bm *Bitmap2D) GetBitBySize(msize int) int {
	a := msize / 8
	b := msize % 8
	if bm.bits[a]&(1<<b) > 0 {
		return 1
	}
	return 0
}

func (bm *Bitmap2D) GetBit(x, y int) int {
	return bm.GetBitBySize(y*bm.dimX + x)
}

func (bm *Bitmap2D) SetBit(x, y int, v int) {
	msize := y*bm.dimX + x
	a := msize / 8
	b := msize % 8
	if v > 0 {
		bm.bits[a] |= (1 << b)
	} else {
		bm.bits[a] &^= (1 << b)
	}
}
