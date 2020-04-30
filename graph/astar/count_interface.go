package astar

// CountWeight count weight interface
type CountWeight interface {
	Weight(graph *Graph, tile *Tile, ptile *Tile)
}

// CountCost count cost interface
type CountCost interface {
	Cost(graph *Graph, tile *Tile, ptile *Tile)
}

// SimpleCost simple cost and no attr
type SimpleCost struct {
}

// Cost ca
func (cost *SimpleCost) Cost(graph *Graph, tile *Tile, ptile *Tile) {
	tile.Cost = ptile.Cost + 1
}

// SimpleWeight simple weight and no attr
type SimpleWeight struct {
}

// Weight ca
func (cost *SimpleWeight) Weight(graph *Graph, tile *Tile, ptile *Tile) {
	_, end := graph.GetTarget()
	absY := abs(tile.Y - end.Y)
	absX := abs(tile.X - end.X)
	tile.Weight = -(absX + absY + tile.Cost)
}
