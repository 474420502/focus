package bitmap

// BitmapEx bitmap with dimension
type BitmapEx struct {
	dimX, dimY int
	bits       []byte
}

// Clone clone bitmap
func (bm *BitmapEx) Clone() *BitmapEx {
	other := &BitmapEx{dimX: bm.dimX, dimY: bm.dimY}
	other.bits = make([]byte, len(bm.bits))
	copy(other.bits, bm.bits)
	return other
}

// NewBitmapWithDimension create bit map
func NewBitmapWithDimension(dx, dy int) *BitmapEx {
	bm := &BitmapEx{dimX: dx, dimY: dy}
	bm.bits = make([]byte, (bm.dimX*bm.dimY+7)/8)
	return bm
}

// GetBytes get bitmap all bytes(bit)
func (bm *BitmapEx) GetBytes() []byte {
	return bm.bits
}

// SetBytes set bitmap byts
func (bm *BitmapEx) SetBytes(offset int, src []byte) {
	copy(bm.bits[offset:len(bm.bits)], src)
}

// GetBitBySize get bit value by offset size
func (bm *BitmapEx) GetBitBySize(msize int) int {
	a := msize / 8
	b := msize % 8
	if bm.bits[a]&(1<<b) > 0 {
		return 1
	}
	return 0
}

// GetBit get bit value by x, y, dimX
func (bm *BitmapEx) GetBit(x, y int) int {
	msize := y*bm.dimX + x
	a := msize / 8
	b := msize % 8
	if bm.bits[a]&(1<<b) > 0 {
		return 1
	}
	return 0
}

// SetBit set bit value by x, y, dimX
func (bm *BitmapEx) SetBit(x, y int, v int) {
	msize := y*bm.dimX + x
	a := msize / 8
	b := msize % 8
	if v > 0 {
		bm.bits[a] |= (1 << b)
	} else {
		bm.bits[a] &^= (1 << b)
	}
}

// SetBitBySize set bit value by offset size
func (bm *BitmapEx) SetBitBySize(msize int, v int) {
	a := msize / 8
	b := msize % 8
	if v > 0 {
		bm.bits[a] |= (1 << b)
	} else {
		bm.bits[a] &^= (1 << b)
	}
}
