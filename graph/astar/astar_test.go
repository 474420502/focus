package astar

import (
	"testing"
)

// const VX = 8
// const VY = 8

// func TestCase1(t *testing.T) {
// 	g := make([]byte, VX*VY)
// 	paths := make([]Point, 0)
// 	Tr(Point{4, 7}, g, paths)
// }

// func Tr(cur Point, graph []byte, paths []Point) {
// 	paths = append(paths, cur)
// 	msize := cur.y*VX + cur.x

// 	content := "\n"
// 	for y := 0; y < VY; y++ {
// 		for x := 0; x < VX; x++ {
// 			showmsize := y*VX + x
// 			content += fmt.Sprintf("%03d ", graph[showmsize])
// 		}
// 		content += "\n"
// 	}
// 	log.Println(content)

// 	Left(cur, msize, graph, paths)
// 	Right(cur, msize, graph, paths)
// 	Up(cur, msize, graph, paths)
// 	Down(cur, msize, graph, paths)
// }

// func Left(cur Point, msize int, graph []byte, paths []Point) {

// 	leftx := cur.x - 1
// 	if leftx < 0 {
// 		return
// 	}
// 	ncur := Point{leftx, cur.y} // Left

// 	if graph[ncur.y*VX+ncur.x]&0b01000000 > 0 {
// 		return
// 	}

// 	ng := make([]byte, VX*VY)
// 	copy(ng, graph)
// 	ng[msize] |= 0b01000000

// 	npaths := make([]Point, len(paths))
// 	copy(npaths, paths)

// 	Tr(ncur, ng, npaths)
// }

// func Right(cur Point, msize int, graph []byte, paths []Point) {
// 	rightx := cur.x + 1
// 	if rightx >= VX {
// 		return
// 	}

// 	ncur := Point{rightx, cur.y} // Left
// 	if graph[ncur.y*VX+ncur.x]&0b01000000 > 0 {
// 		return
// 	}

// 	ng := make([]byte, VX*VY)
// 	copy(ng, graph)
// 	ng[msize] |= 0b01000000

// 	npaths := make([]Point, len(paths))
// 	copy(npaths, paths)

// 	Tr(ncur, ng, npaths)
// }

// func Up(cur Point, msize int, graph []byte, paths []Point) {
// 	upy := cur.y + 1
// 	if upy >= VY {
// 		return
// 	}

// 	ncur := Point{cur.x, upy} // Left
// 	if graph[ncur.y*VX+ncur.x]&0b01000000 > 0 {
// 		return
// 	}

// 	ng := make([]byte, VX*VY)
// 	copy(ng, graph)
// 	ng[msize] |= 0b01000000

// 	npaths := make([]Point, len(paths))
// 	copy(npaths, paths)

// 	Tr(ncur, ng, npaths)
// }

// func Down(cur Point, msize int, graph []byte, paths []Point) {

// 	downy := cur.y - 1
// 	if downy < 0 {
// 		return
// 	}

// 	ncur := Point{cur.x, downy} // Left
// 	if graph[ncur.y*VX+ncur.x]&0b01000000 > 0 {
// 		return
// 	}

// 	ng := make([]byte, VX*VY)
// 	copy(ng, graph)
// 	ng[msize] |= 0b01000000

// 	npaths := make([]Point, len(paths))
// 	copy(npaths, paths)

// 	Tr(ncur, ng, npaths)
// }

func TestCase2(t *testing.T) {
	graph := New(8, 8)
	graph.SetTarget(3, 7, 0, 0)
	graph.Search()
	t.Error(graph)
}
