package astar

import (
	"bufio"
	"bytes"
	"log"
	"regexp"

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

	path []*Tile

	Tiles       [][]*Tile
	getNeighbor func(graph *Graph, tile *Tile) []*Tile
	openHeap    *heap.Tree
}

// Point point x y
type Point struct {
	X, Y int
	Attr byte
}

// Tile node
type Tile struct {
	X, Y    int
	Cost    int
	Weight  int
	IsCount bool

	Attr byte
}

func (tile *Tile) countCost(graph *Graph, ptile *Tile) {
	tile.Cost = ptile.Cost + 1
}

func (tile *Tile) countWeight(graph *Graph, ptile *Tile) {
	_, end := graph.GetTarget()
	absY := abs(tile.Y - end.Y)
	absX := abs(tile.X - end.X)
	tile.Weight = -(absX + absY + tile.Cost)
}

func weightCompare(x1, x2 interface{}) int {
	p1, p2 := x1.(*Tile), x2.(*Tile)
	if p1.Weight > p2.Weight { // 权重大的优先
		return 1
	}
	return -1
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

	graph.SetNeighborFunc(GetNeighbor4)
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

	graph.SetNeighborFunc(GetNeighbor4)
	graph.openHeap = heap.New(weightCompare)
	return graph
}

// SetNeighborFunc use the function  different directions
func (graph *Graph) SetNeighborFunc(neighborfunc func(graph *Graph, tile *Tile) []*Tile) {
	graph.getNeighbor = neighborfunc
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
				continue
			case END:
				graph.end = &Point{Y: y, X: x}
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

// GetNeighbor8 8向寻址 eight direction
func GetNeighbor8(graph *Graph, tile *Tile) []*Tile {
	var result []*Tile
	for _, neighbor := range [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}, {1, 1}, {-1, 1}, {-1, 1}, {1, -1}} {
		x := tile.X + neighbor[0]
		y := tile.Y + neighbor[1]
		if x < 0 || y < 0 || x >= graph.dimX || y >= graph.dimY {
			continue
		}

		ntile := graph.Tiles[y][x]
		result = append(result, ntile)
	}
	return result
}

// GetNeighbor4 四向寻址 four direction
func GetNeighbor4(graph *Graph, tile *Tile) []*Tile {
	var result []*Tile
	for _, neighbor := range [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
		x := tile.X + neighbor[0]
		y := tile.Y + neighbor[1]
		if x < 0 || y < 0 || x >= graph.dimX || y >= graph.dimY {
			continue
		}

		ntile := graph.Tiles[y][x]
		result = append(result, ntile)
	}
	return result
}

func abs(v int) (ret int) {
	return (v ^ v>>31) - v>>31
}

// GetSteps result == len(path) - 1
func (graph *Graph) GetSteps() int {
	return len(graph.path) - 1 // contains start point so -1
}

// GetTarget start end point
func (graph *Graph) GetTarget() (*Point, *Point) {
	return graph.start, graph.end
}

// GetPath the astar path
func (graph *Graph) GetPath() []*Tile {
	return graph.path // contains start point so -1
}

// GetDimension get dimension info
func (graph *Graph) GetDimension() (int, int) {
	return graph.dimX, graph.dimY // contains start point so -1
}

// GetStringTiles get the string of tiles map info
func (graph *Graph) GetStringTiles() string {
	var content []byte
	content = append(content, '\n')
	for y := 0; y < graph.dimY; y++ {
		for x := 0; x < graph.dimX; x++ {
			content = append(content, graph.Tiles[y][x].Attr)
		}
		content = append(content, '\n')
	}
	return string(content)
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

	var path []*Tile
	for {
		if itile, ok := graph.openHeap.Pop(); ok {
			tile := itile.(*Tile)

			if tile == endTile {
				path = append(path, tile)
				// 回找路径
				for tile != startTile {

					returnTile := tile
					for _, ntile := range graph.getNeighbor(graph, tile) {
						if ntile.IsCount {
							if returnTile.Cost > ntile.Cost {
								returnTile = ntile
							}
						}
					}
					tile = returnTile
					path = append(path, tile)
					tile.Attr = PATH
				}

				startTile.Attr = START
				graph.path = path
				return true
			}

			for _, ntile := range graph.getNeighbor(graph, tile) {
				if ntile.IsCount == false && ntile.Attr != BLOCK {
					ntile.countCost(graph, tile)
					ntile.countWeight(graph, tile)
					ntile.IsCount = true
					// 处理ntile权值
					graph.openHeap.Put(ntile)
				}
			}

		} else {
			log.Println("path can not found")
			break
		}
	}

	return false
}
