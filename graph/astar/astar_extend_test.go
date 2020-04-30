package astar

import "testing"

type MyCost struct {
}

const (
	Mountain = byte('m')
	RIVER    = byte('r')
)

// Cost ca
func (cost *MyCost) Cost(graph *Graph, tile *Tile, ptile *Tile) {
	moveCost := 0
	switch tile.Attr {
	case Mountain:
		moveCost = 3
	case PLAIN:
		moveCost = 1
	case RIVER:
		moveCost = 2
	}
	tile.Cost = ptile.Cost + moveCost
}

func TestMyCost(t *testing.T) {

	graph := NewWithTiles(`
s.......
xx..mmmm
....mmmm
.....rre
`)
	graph.SetCountCost(&MyCost{})
	graph.Search()

	var should, result string
	should = `
sooo....
xx.ommmm
...ommmm
...ooooe
`
	result = graph.GetSingleStringTiles()
	if result != should {
		t.Error("result:\n", result, "\nshould:\n", should)
	}

	graph.Clear()
	graph.SetStringTiles(`
s......m
xx..mmmr
...rmmmr
..rrrrre
`)
	graph.Search()
	result = graph.GetSingleStringTiles()
	should = `
sooooooo
xx..mmmo
...rmmmo
..rrrrre
`
	result = graph.GetSingleStringTiles()
	if result != should {
		t.Error("result:\n", result, "\nshould:\n", should)
	}
}
