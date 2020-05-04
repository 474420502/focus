package bitmap

// NewBitmap create bit map
func NewBitmap(dx, dy int) *Bitmap {
	bm := &Bitmap{}
	bm.bits = make([]byte, (dx*dy+7)/8)
	return bm
}

// Clone clone bitmap
func (bm *Bitmap) Clone() *Bitmap {
	other := &Bitmap{}
	other.bits = make([]byte, len(bm.bits))
	copy(other.bits, bm.bits)
	return other
}

// Bitmap not with the info of dimension
type Bitmap struct {
	// dimX, dimY int
	bits []byte
}

// GetBytes get bitmap all bytes(bit)
func (bm *Bitmap) GetBytes() []byte {
	return bm.bits
}

// SetBytes set bitmap byts
func (bm *Bitmap) SetBytes(offset int, src []byte) {
	copy(bm.bits[offset:len(bm.bits)], src)

}

// GetBitBySize get bit value by offset size
func (bm *Bitmap) GetBitBySize(msize int) int {
	a := msize / 8
	b := msize % 8
	if bm.bits[a]&(1<<b) > 0 {
		return 1
	}
	return 0
}

// GetBit get bit value by x, y, dimX
func (bm *Bitmap) GetBit(dimX int, x, y int) int {
	msize := y*dimX + x
	a := msize / 8
	b := msize % 8
	if bm.bits[a]&(1<<b) > 0 {
		return 1
	}
	return 0
}

// SetBit set bit value by x, y, dimX
func (bm *Bitmap) SetBit(dimX int, x, y int, v int) {
	msize := y*dimX + x
	a := msize / 8
	b := msize % 8
	if v > 0 {
		bm.bits[a] |= (1 << b)
	} else {
		bm.bits[a] &^= (1 << b)
	}
}

// SetBitBySize set bit value by offset size
func (bm *Bitmap) SetBitBySize(msize int, v int) {
	a := msize / 8
	b := msize % 8
	if v > 0 {
		bm.bits[a] |= (1 << b)
	} else {
		bm.bits[a] &^= (1 << b)
	}
}
