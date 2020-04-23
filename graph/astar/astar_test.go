package astar

import (
	"fmt"
	"log"
	"testing"
)

const VX = 8
const VY = 8

func TestCase1(t *testing.T) {
	g := make([]byte, VX*VY)
	paths := make([]P, 0)
	Tr(P{4, 7}, g, paths)
}

type P struct {
	x, y int
}

func Tr(cur P, g []byte, paths []P) {
	paths = append(paths, cur)
	msize := cur.y*VX + cur.x

	content := "\n"
	for y := 0; y < VY; y++ {
		for x := 0; x < VX; x++ {
			showmsize := y*VX + x
			content += fmt.Sprintf("%03d ", g[showmsize])
		}
		content += "\n"
	}
	log.Println(content)

	Left(cur, msize, g, paths)
	Right(cur, msize, g, paths)
	Up(cur, msize, g, paths)
	Down(cur, msize, g, paths)
}

func Left(cur P, msize int, g []byte, paths []P) {

	leftx := cur.x - 1
	if leftx < 0 {
		return
	}
	ncur := P{leftx, cur.y} // Left

	if g[ncur.y*VX+ncur.x]&0b01000000 > 0 {
		return
	}

	ng := make([]byte, VX*VY)
	copy(ng, g)
	ng[msize] |= 0b01000000

	npaths := make([]P, len(paths))
	copy(npaths, paths)

	Tr(ncur, ng, npaths)
}

func Right(cur P, msize int, g []byte, paths []P) {
	rightx := cur.x + 1
	if rightx >= VX {
		return
	}

	ncur := P{rightx, cur.y} // Left
	if g[ncur.y*VX+ncur.x]&0b01000000 > 0 {
		return
	}

	ng := make([]byte, VX*VY)
	copy(ng, g)
	ng[msize] |= 0b01000000

	npaths := make([]P, len(paths))
	copy(npaths, paths)

	Tr(ncur, ng, npaths)
}

func Up(cur P, msize int, g []byte, paths []P) {
	upy := cur.y + 1
	if upy >= VY {
		return
	}

	ncur := P{cur.x, upy} // Left
	if g[ncur.y*VX+ncur.x]&0b01000000 > 0 {
		return
	}

	ng := make([]byte, VX*VY)
	copy(ng, g)
	ng[msize] |= 0b01000000

	npaths := make([]P, len(paths))
	copy(npaths, paths)

	Tr(ncur, ng, npaths)
}

func Down(cur P, msize int, g []byte, paths []P) {

	downy := cur.y - 1
	if downy < 0 {
		return
	}

	ncur := P{cur.x, downy} // Left
	if g[ncur.y*VX+ncur.x]&0b01000000 > 0 {
		return
	}

	ng := make([]byte, VX*VY)
	copy(ng, g)
	ng[msize] |= 0b01000000

	npaths := make([]P, len(paths))
	copy(npaths, paths)

	Tr(ncur, ng, npaths)
}
