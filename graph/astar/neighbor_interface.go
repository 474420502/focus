package astar

// Neighbor 临近类
type Neighbor interface {
	GetNeighbor(graph *Graph, tile *Tile) []*Tile
}

// Neighbor8  eight direction
type Neighbor8 struct {
}

// GetNeighbor must interface
func (neighbor *Neighbor8) GetNeighbor(graph *Graph, tile *Tile) []*Tile {
	var result []*Tile
	for _, neighbor := range [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}, {1, 1}, {-1, 1}, {-1, -1}, {1, -1}} {
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

// Neighbor4  four direction
type Neighbor4 struct {
}

// GetNeighbor four direction
func (neighbor *Neighbor4) GetNeighbor(graph *Graph, tile *Tile) []*Tile {
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
