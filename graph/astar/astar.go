package astar

import (
	"bufio"
	"bytes"
	"regexp"
	"sort"

	"github.com/474420502/focus/tree/heap"
)

// AttributeEnum 属性类型
const (
	// SKIP skip set attr. used by SetStringTiles
	SKIP = byte('*')
	// PLAIN  point can be arrived to
	PLAIN = byte('.')
	// BLOCK  point can not be arrived to
	BLOCK = byte('x')
	// START  the start point
	START = byte('s')
	// END  the end point
	END = byte('e')
	// PATH  not contains start and end.
	PATH = byte('o')
)

// Graph Astar struct
type Graph struct {
	dimX, dimY int
	start, end *Point

	pathlist PathList

	Tiles [][]*Tile

	// getNeighbor func(graph *Graph, tile *Tile) []*Tile

	// countCost   func(graph *Graph, tile *Tile, ptile *Tile)
	// countWeight func(graph *Graph, tile *Tile, ptile *Tile)
	neighbor    Neighbor
	countCost   CountCost
	countWeight CountWeight

	openHeap *heap.Tree
}

// Point point x y
type Point struct {
	X, Y int
	Attr byte
}

// Path search astar path
type Path []*Tile

// PathList pathlist
type PathList []Path

func (pl PathList) Less(i, j int) bool {
	if len(pl[i]) < len(pl[j]) {
		return true
	}
	return false
}

func (pl PathList) Swap(i, j int) {
	pl[i], pl[j] = pl[j], pl[i]
}

func (pl PathList) Len() int {
	return len(pl)
}

// Tile node
type Tile struct {
	X, Y    int
	Cost    int
	Weight  int
	IsCount bool

	Attr byte
}

// New create astar
func New(dimX, dimY int) *Graph {
	graph := &Graph{dimX: dimX, dimY: dimY}

	graph.Tiles = make([][]*Tile, graph.dimY)
	for y := 0; y < graph.dimY; y++ {
		xtiles := make([]*Tile, graph.dimX)
		for x := 0; x < graph.dimX; x++ {
			xtiles[x] = &Tile{Y: y, X: x, Attr: PLAIN}
		}
		graph.Tiles[y] = xtiles
	}

	graph.SetNeighbor(&Neighbor4{})
	graph.SetCountCost(&SimpleCost{})
	graph.SetCountWeight(&SimpleWeight{})
	graph.openHeap = heap.New(weightCompare)
	return graph
}

// NewWithTiles create astar
func NewWithTiles(tiles string) *Graph {

	reader := bufio.NewReader(bytes.NewReader([]byte(tiles)))
	var tilebuffer [][]byte
	xMax := 0
	for {
		line, _, err := reader.ReadLine()

		if err != nil {
			break
		}

		if len(line) == 0 {
			continue
		}

		found := regexp.MustCompile("[^\\s]+").FindAll(line, -1)
		if len(found) != 0 {
			buffer := []byte{}

			for _, foundbuf := range found {
				buffer = append(buffer, foundbuf...)
			}

			if xMax < len(buffer) {
				xMax = len(buffer)
			}

			tilebuffer = append(tilebuffer, buffer)
		}

	}

	graph := &Graph{dimX: xMax, dimY: len(tilebuffer)}
	graph.Tiles = make([][]*Tile, graph.dimY)
	for y := 0; y < graph.dimY; y++ {
		xtiles := make([]*Tile, graph.dimX)
		xbuffer := tilebuffer[y]
		for x := 0; x < graph.dimX; x++ {
			if x < len(xbuffer) {
				Attr := xbuffer[x]
				switch Attr {
				case SKIP:
					xtiles[x] = &Tile{Y: y, X: x, Attr: PLAIN}
				case START:
					graph.start = &Point{Y: y, X: x}
					xtiles[x] = &Tile{Y: y, X: x, Attr: Attr}
				case END:
					graph.end = &Point{Y: y, X: x}
					xtiles[x] = &Tile{Y: y, X: x, Attr: Attr}
				default:
					xtiles[x] = &Tile{Y: y, X: x, Attr: Attr}
				}
			} else {
				xtiles[x] = &Tile{Y: y, X: x, Attr: PLAIN}
			}
		}
		graph.Tiles[y] = xtiles
	}

	graph.SetNeighbor(&Neighbor4{})
	graph.SetCountCost(&SimpleCost{})
	graph.SetCountWeight(&SimpleWeight{})

	graph.openHeap = heap.New(weightCompare)
	return graph
}

func weightCompare(x1, x2 interface{}) int {
	p1, p2 := x1.(*Tile), x2.(*Tile)
	if p1.Weight > p2.Weight { // 权重大的优先
		return 1
	}
	return -1
}

// GetAttr get tiles attribute not contain start end info
func (graph *Graph) GetAttr(x, y int) byte {
	return graph.Tiles[y][x].Attr
}

// SetAttr set tiles attribute
func (graph *Graph) SetAttr(x, y int, attr byte) {
	graph.Tiles[y][x].Attr = attr
}

// SetCountWeight use the function  different weight
func (graph *Graph) SetCountWeight(count CountWeight) {
	graph.countWeight = count
}

// SetCountCost use the function  different cost
func (graph *Graph) SetCountCost(count CountCost) {
	graph.countCost = count
}

// SetNeighbor use the function  different directions
func (graph *Graph) SetNeighbor(neighbor Neighbor) {
	graph.neighbor = neighbor
}

// SetTarget start point end point
func (graph *Graph) SetTarget(sx, sy, ex, ey int) {
	graph.start = &Point{Y: sy, X: sx}
	graph.end = &Point{Y: ey, X: ex}
}

// SetStringTiles if want some tile do nothing, can use SKIP.
func (graph *Graph) SetStringTiles(strtile string) {
	bufreader := bytes.NewReader([]byte(strtile))
	reader := bufio.NewReader(bufreader)

	for y := 0; ; {
		line, _, err := reader.ReadLine()

		if err != nil {
			break
		}

		if len(line) == 0 {
			continue
		}

		for x, i := 0, 0; x < graph.dimX && i < len(line); i++ {
			attr := line[i]
			switch attr {
			case START:
				graph.start = &Point{Y: y, X: x}
				x++
				continue
			case END:
				graph.end = &Point{Y: y, X: x}
				x++
				continue
			case '\t':
				continue
			case ' ':
				continue
			case '\n':
				continue
			}
			if attr != SKIP {
				graph.Tiles[y][x].Attr = attr
			}
			x++
		}
		y++
	}
}

func abs(v int) (ret int) {
	return (v ^ v>>31) - v>>31
}

// GetSteps result == len(path) - 1
func (graph *Graph) GetSteps(path Path) int {
	return len(path) - 1 // contains start point so -1
}

// GetSingleSteps result == len(pathlist[0]) - 1
func (graph *Graph) GetSingleSteps() int {
	return len(graph.pathlist[0]) - 1 // contains start point so -1
}

// GetTarget start end point
func (graph *Graph) GetTarget() (*Point, *Point) {
	return graph.start, graph.end
}

// GetPath the astar path
func (graph *Graph) GetPath() Path {
	return graph.pathlist[0] // contains start point so -1
}

// GetMultiPath get multi  the astar path of same cost
func (graph *Graph) GetMultiPath() []Path {
	return graph.pathlist // contains start point so -1
}

// GetDimension get dimension info
func (graph *Graph) GetDimension() (int, int) {
	return graph.dimX, graph.dimY // contains start point so -1
}

// GetTiles get astar map info not contain target(start, end)
func (graph *Graph) GetTiles() string {
	var data [][]byte = make([][]byte, graph.dimY)

	// content = append(content, '\n')
	for y := 0; y < graph.dimY; y++ {
		xdata := make([]byte, graph.dimX)
		for x := 0; x < graph.dimX; x++ {
			xdata[x] = graph.Tiles[y][x].Attr
		}
		data[y] = xdata
		// content = append(content, '\n')
	}

	var content []byte
	content = append(content, '\n')
	for y := 0; y < graph.dimY; y++ {
		for x := 0; x < graph.dimX; x++ {
			content = append(content, data[y][x])
		}
		content = append(content, '\n')
	}

	return string(content)
}

// GetTilesWithTarget get astar map info with target
func (graph *Graph) GetTilesWithTarget() string {
	var data [][]byte = make([][]byte, graph.dimY)

	// content = append(content, '\n')
	for y := 0; y < graph.dimY; y++ {
		xdata := make([]byte, graph.dimX)
		for x := 0; x < graph.dimX; x++ {
			xdata[x] = graph.Tiles[y][x].Attr
		}
		data[y] = xdata
		// content = append(content, '\n')
	}

	if graph.start != nil {
		data[graph.start.Y][graph.start.X] = START
	}

	if graph.end != nil {
		data[graph.end.Y][graph.end.X] = END
	}

	var content []byte
	content = append(content, '\n')
	for y := 0; y < graph.dimY; y++ {
		for x := 0; x < graph.dimX; x++ {
			content = append(content, data[y][x])
		}
		content = append(content, '\n')
	}

	return string(content)
}

// GetPathTiles get the string of tiles map info
func (graph *Graph) GetPathTiles(path Path) string {
	var data [][]byte = make([][]byte, graph.dimY)

	// content = append(content, '\n')
	for y := 0; y < graph.dimY; y++ {
		xdata := make([]byte, graph.dimX)
		for x := 0; x < graph.dimX; x++ {
			xdata[x] = graph.Tiles[y][x].Attr
		}
		data[y] = xdata
		// content = append(content, '\n')
	}

	for _, t := range path {
		data[t.Y][t.X] = PATH
	}

	data[graph.start.Y][graph.start.X] = START
	data[graph.end.Y][graph.end.X] = END

	var content []byte
	content = append(content, '\n')
	for y := 0; y < graph.dimY; y++ {
		for x := 0; x < graph.dimX; x++ {
			content = append(content, data[y][x])
		}
		content = append(content, '\n')
	}

	return string(content)
}

// GetSinglePathTiles get the string of tiles map info
func (graph *Graph) GetSinglePathTiles() string {
	return graph.GetPathTiles(graph.pathlist[0])
}

// Clear astar 搜索
func (graph *Graph) Clear() {

	for y := 0; y < graph.dimY; y++ {
		for x := 0; x < graph.dimX; x++ {
			tile := graph.Tiles[y][x]
			switch tile.Attr {
			case PATH:
				tile.Attr = PLAIN
			case START:
				tile.Attr = graph.start.Attr
			case END:
				tile.Attr = graph.end.Attr
			}

			tile.Cost = 0
			tile.IsCount = false
			tile.Weight = 0
		}
	}
}

// Search astar search path
func (graph *Graph) Search() bool {
	return graph.search(false)
}

// SearchMulti astar search multi path
func (graph *Graph) SearchMulti() bool {
	return graph.search(true)
}

func (graph *Graph) singlePath(tile *Tile, startTile *Tile, path []*Tile) {
	// 回找路径
	for tile != startTile {

		returnTile := tile
		for _, ntile := range graph.neighbor.GetNeighbor(graph, tile) {
			if ntile.IsCount {
				if returnTile.Cost >= ntile.Cost {
					returnTile = ntile
				}
			}
		}
		tile = returnTile
		path = append(path, tile)
	}

	graph.pathlist = append(graph.pathlist, path)
	return
}

func (graph *Graph) multiPath(tile *Tile, startTile *Tile, path []*Tile) {
	path = append(path, tile)
	// 回找路径
	if tile != startTile {

		var minCostTiles []*Tile
		returnTile := tile
		for _, ntile := range graph.neighbor.GetNeighbor(graph, tile) {
			if ntile.IsCount {
				if returnTile.Cost > ntile.Cost {
					minCostTiles = minCostTiles[0:0]
					returnTile = ntile
					minCostTiles = append(minCostTiles, ntile)
				} else if returnTile.Cost == ntile.Cost {
					minCostTiles = append(minCostTiles, ntile)
				}
			}
		}

		for _, rtile := range minCostTiles {
			npath := make([]*Tile, len(path))
			copy(npath, path)
			graph.multiPath(rtile, startTile, npath)
		}

		// tile.Attr = PATH
	} else {
		graph.pathlist = append(graph.pathlist, path)
	}

	return
}

// search astar search path
func (graph *Graph) search(multi bool) bool {

	defer func() {
		graph.openHeap.Clear()
	}()

	if graph.start == nil {
		panic("not set start point")
	}

	if graph.end == nil {
		panic("not set end point")
	}

	startTile := graph.Tiles[graph.start.Y][graph.start.X]
	graph.start.Attr = startTile.Attr
	startTile.IsCount = true
	startTile.Attr = START

	endTile := graph.Tiles[graph.end.Y][graph.end.X]
	graph.end.Attr = endTile.Attr
	endTile.IsCount = false
	endTile.Attr = END

	graph.openHeap.Put(startTile)

	for {
		if itile, ok := graph.openHeap.Pop(); ok {
			tile := itile.(*Tile)

			if tile == endTile {

				graph.pathlist = make([]Path, 0)

				var path Path
				path = append(path, tile)
				if multi {
					graph.multiPath(tile, startTile, path)
					sort.Sort(graph.pathlist)
				} else {
					graph.singlePath(tile, startTile, path)
				}

				return true
			}

			for _, ntile := range graph.neighbor.GetNeighbor(graph, tile) {
				if ntile.IsCount == false && ntile.Attr != BLOCK {
					graph.countCost.Cost(graph, ntile, tile)
					graph.countWeight.Weight(graph, ntile, tile)
					ntile.IsCount = true
					// 处理ntile权值
					graph.openHeap.Put(ntile)
				}
			}

		} else {
			// log.Println("path can not found")
			break
		}
	}

	return false
}
