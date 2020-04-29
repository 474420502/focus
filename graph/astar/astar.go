package astar

import (
	"bufio"
	"bytes"
	"log"

	"github.com/474420502/focus/tree/heap"
)

type Graph struct {
	dimX, dimY  int
	start, end  *point
	Tiles       [][]*Tile
	getNeighbor func(graph *Graph, tile *Tile) []*Tile
	openHeap    *heap.Tree

	isDebug     bool
	debugString string
}

type point struct {
	X, Y int
}

type Tile struct {
	X, Y    int
	Cost    int
	Weight  int
	IsCount bool

	Attr byte
}

func weightCompare(x1, x2 interface{}) int {
	p1, p2 := x1.(*Tile), x2.(*Tile)
	if p1.Weight > p2.Weight { // 权重大的优先
		return 1
	}
	return -1
}

func New(dimX, dimY int) *Graph {
	graph := &Graph{dimX: dimX, dimY: dimY}

	graph.Tiles = make([][]*Tile, graph.dimY)
	for y := 0; y < graph.dimY; y++ {
		xtiles := make([]*Tile, graph.dimX)
		for x := 0; x < graph.dimX; x++ {
			xtiles[x] = &Tile{Y: y, X: x, Attr: '.'}
		}
		graph.Tiles[y] = xtiles
	}

	graph.SetNeighborFunc(GetNeighbor4)
	graph.openHeap = heap.New(weightCompare)
	return graph
}

func (graph *Graph) SetNeighborFunc(neighborfunc func(graph *Graph, tile *Tile) []*Tile) {
	graph.getNeighbor = neighborfunc
}

func (graph *Graph) SetTarget(sx, sy, ex, ey int) {
	graph.start = &point{Y: sy, X: sx}
	graph.end = &point{Y: ey, X: ex}
}

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
			case 's':
				graph.start = &point{Y: y, X: x}
				continue
			case 'e':
				graph.end = &point{Y: y, X: x}
				continue
			case '\t':
				continue
			case ' ':
				continue
			case '\n':
				continue
			}
			graph.Tiles[y][x].Attr = attr
			x++
		}

		y++
	}

}

// GetNeighbor8 8向寻址
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

// GetNeighbor4 四向寻址
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

func (tile *Tile) countCost(ptile *Tile) {
	tile.Cost = ptile.Cost + 1
}

func (graph *Graph) countWeight(tile *Tile) {
	absY := abs(tile.Y - graph.end.Y)
	absX := abs(tile.X - graph.end.X)
	tile.Weight = -(absX + absY + tile.Cost)
}

// Clear astar 搜索
func (graph *Graph) Clear() {
	for y := 0; y < graph.dimY; y++ {
		for x := 0; x < graph.dimX; x++ {
			tile := graph.Tiles[y][x]
			tile.Attr = '.'
			tile.Cost = 0
			tile.IsCount = false
			tile.Weight = 0
		}
	}
}

// Search astar 搜索
func (graph *Graph) Search() []*Tile {
	startTile := graph.Tiles[graph.start.Y][graph.start.X]
	startTile.IsCount = true
	startTile.Attr = 's'

	endTile := graph.Tiles[graph.end.Y][graph.end.X]
	endTile.IsCount = false
	endTile.Attr = 'e'

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
					tile.Attr = 'o'
				}

				startTile.Attr = 's'

				if graph.isDebug {
					var content []byte
					content = append(content, '\n')
					for y := 0; y < graph.dimY; y++ {
						for x := 0; x < graph.dimX; x++ {
							content = append(content, graph.Tiles[y][x].Attr)
						}
						content = append(content, '\n')
					}
					graph.debugString = string(content)
				}
				return path
			}

			for _, ntile := range graph.getNeighbor(graph, tile) {
				if ntile.IsCount == false && ntile.Attr != 'x' {
					ntile.countCost(tile)
					graph.countWeight(ntile)
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

	return path
}
