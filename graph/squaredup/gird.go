package squaredup

import (
	"bytes"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

// func NewBytesmap()

// Gird map
type Gird struct {
	bits   []byte
	X, Y   int
	weight int
	cost   int
}

func NewGird(dimY, dimX int, vlen int) *Gird {
	g := &Gird{}
	g.bits = make([]byte, dimY*dimX*vlen)
	return g
}

// Compare get bit value by x, y, dimX
func (bm *Gird) Compare(other *Gird) int {
	return bytes.Compare(bm.bits, other.bits)
}

// Clone get bit value by x, y, dimX
func (bm *Gird) Clone() *Gird {
	other := &Gird{}
	other.bits = make([]byte, len(bm.bits))

	other.X = bm.X
	other.Y = bm.Y
	copy(other.bits, bm.bits)
	return other
}

// SwapValue get bit value by x, y, dimX
func (bm *Gird) SwapValue(dimX, vlen int, x, y int) {
	dstmsize := (y*dimX + x) * vlen
	srcmsize := (bm.Y*dimX + bm.X) * vlen

	bm.bits[dstmsize], bm.bits[srcmsize] = bm.bits[srcmsize], bm.bits[dstmsize]

	bm.Y = y
	bm.X = x
}

// SetGirdByString set  value by x, y, dimX
func (bm *Gird) SetGirdByString(dimX, vlen int, value string) {

	y := 0

	for _, ystr := range strings.Split(value, "\n") {
		x := 0
		if ystr != "" {
			for _, xstr := range regexp.MustCompile(`\w+`).FindAllString(ystr, -1) {
				if xstr != "" {
					msize := (y*dimX + x) * vlen
					v, err := strconv.ParseUint(xstr, 10, vlen*8)
					if err != nil {
						panic(err)
					}
					if v == 0 {
						bm.Y = y
						bm.X = x
					}
					bm.bits[msize] = byte(v)
					// binary.BigEndian.PutUint16(bm.bits[msize:], uint16(v))
					x++
				}

			}
			y++
		}
	}
}

func (bm *Gird) GetGirdString(dimY, dimX, vlen int) string {

	dnum := strconv.Itoa(int(math.Log10(float64(dimY*dimX)) + 1))

	content := "\n"

	for y := 0; y < dimY; y++ {
		for x := 0; x < dimX; x++ {
			msize := (y*dimX + x) * vlen
			v := bm.bits[msize]
			content += fmt.Sprintf("%"+dnum+"d ", v)
		}

		content += "\n"
	}

	return content
}

// SetValue set  value by x, y, dimX
func (bm *Gird) SetValue(dimX, vlen int, x, y int, value []byte) {
	msize := (y*dimX + x) * vlen
	copy(bm.bits[msize:msize+vlen], value)
}

// GetValue get bit value by x, y, dimX
func (bm *Gird) GetValue(dimX, vlen int, x, y int) []byte {
	msize := (y*dimX + x) * vlen
	return bm.bits[msize : msize+vlen]
}

// GetValues get bit values
func (bm *Gird) GetValues() []byte {
	return bm.bits
}
