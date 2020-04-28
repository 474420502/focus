package astar

import (
	"bufio"
	"bytes"

	"github.com/474420502/focus/tree/heap"
)

type Graph struct {
	dimX, dimY int
	start, end *point
	Tiles      [][]*Tile

	openHeap *heap.Tree
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

func costCompare(x1, x2 interface{}) int {
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

	graph.openHeap = heap.New(costCompare)
	return graph
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
			case 'e':
				graph.end = &point{Y: y, X: x}
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

type Neighbor struct {
	offsetX, offsetY int
}

func GetDirection() []Neighbor {
	n := make([]Neighbor, 0)
	return n
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

	tile.Weight = absX + absY
}

// Search astar 搜索
func (graph *Graph) Search() []*Tile {
	graph.openHeap.Put(&Tile{X: graph.start.X, Y: graph.start.Y})

	var path []*Tile
	for {
		if itile, ok := graph.openHeap.Pop(); ok {
			tile := itile.(*Tile)

			if tile.Y == graph.end.Y && tile.X == graph.end.X {

				// 回找路径
				for tile.Y != graph.start.Y && tile.X != graph.start.X {
					path = append(path, tile)
					returnTile := tile
					for _, neighbor := range GetDirection() {
						x := tile.X + neighbor.offsetX
						y := tile.Y + neighbor.offsetY
						ntile := graph.Tiles[y][x]
						if ntile.IsCount {
							if returnTile.Cost > ntile.Cost {
								returnTile = ntile
							}
						}
					}
					tile = returnTile
				}
				path = append(path, tile)
				return path
			}

			for _, neighbor := range GetDirection() {
				x := tile.X + neighbor.offsetX
				y := tile.Y + neighbor.offsetY

				ntile := graph.Tiles[y][x]
				if ntile.IsCount == false && ntile.Attr != 'x' {
					ntile = &Tile{Y: y, X: x}
					ntile.countCost(tile)
					graph.countWeight(ntile)
					// 处理ntile权值
					graph.Tiles[y][x] = ntile
					graph.openHeap.Put(ntile)
					ntile.IsCount = true
				}
			}

		} else {
			break
		}
	}

	return path
}
