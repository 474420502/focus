package astarex

import "testing"

// GetCost() int
// SetCost(int)

// GetWeight() int
// SetWeight(int)

// IsCount() bool
// SetCount(bool)

// GetAttribute() interface{} //byte
type Gird struct {
	bits   []byte
	x, y   int
	weight int
	cost   int
}

func (gird *Gird) GetCost() int {
	return gird.cost
}

func (gird *Gird) SetCost(cost int) {
	gird.cost = cost
}

func (gird *Gird) GetWeight() int {
	return gird.weight
}
func (gird *Gird) SetWeight(w int) {
	gird.SetWeight(w)
}

func (gird *Gird) Key() []byte {
	return gird.bits
}

func (gird *Gird) GetAttribute() interface{} { //byte
	return nil
}

// Clone get bit value by x, y, dimX
func (gird *Gird) Clone() *Gird {
	other := &Gird{}
	other.bits = make([]byte, len(gird.bits))

	other.x = gird.x
	other.y = gird.y
	copy(other.bits, gird.bits)
	return other
}

// SwapValue get bit value by x, y, dimX
func (gird *Gird) SwapValue(dimX, x, y int) {
	dstmsize := (y*dimX + x)
	srcmsize := (gird.y*dimX + gird.x)

	gird.bits[dstmsize], gird.bits[srcmsize] = gird.bits[srcmsize], gird.bits[dstmsize]

	gird.y = y
	gird.x = x
}

type SquareUp struct {
	dimX, dimY int
	tsize      int
	AstarEx
}

func NewSquareUp(dimX, dimY int) *SquareUp {
	su := &SquareUp{dimX: dimX, dimY: dimY}
	su.tsize = dimX * dimY
	return su
}

func (su *SquareUp) GetNeighbor(astar *AstarEx, cur Tile) []Tile {
	var result []Tile

	gird := cur.(*Gird)

	for _, offset := range [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {

		x, y := gird.x+offset[0], gird.y+offset[1]

		if x < 0 || y < 0 || x >= su.dimX || y >= su.dimY {
			continue
		}

		neighbor := gird.Clone()
		neighbor.SwapValue(su.dimX, x, y)

		result = append(result, neighbor)
	}

	return result
}

func (su *SquareUp) HandleAttribute(astar *AstarEx, current Tile) bool {
	return true
}

func (su *SquareUp) CountNeighborCost(astar *AstarEx, neighbor, current Tile) {
	neighbor.SetCost(current.GetCost() + 1)
}

func (su *SquareUp) CountNeighborWeight(astar *AstarEx, neighbor, current Tile) {
	ngird := neighbor.(*Gird)
	egird := su.end.(*Gird)

	for y := 0; y < su.dimY; y++ {
		continuously := 1
		for x := 0; x < su.dimX; x++ {
			msize := (y*su.dimX + x)
			if ngird.bits[msize] == egird.bits[msize] {
				ngird.weight += continuously
				continuously++
			} else {
				continuously = 1
			}
			if continuously >= su.dimX {
				ngird.weight += continuously * su.tsize
			}
		}
	}

	ngird.weight -= ngird.cost
}

func TestCase1(t *testing.T) {
	NewSquareUp(3, 3)
}
