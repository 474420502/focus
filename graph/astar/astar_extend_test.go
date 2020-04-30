package astar

import "testing"

type MyCost struct {
}

const (
	Marsh    = byte('M')
	Mountain = byte('m')
	RIVER    = byte('r')
)

// Cost ca
func (cost *MyCost) Cost(graph *Graph, tile *Tile, ptile *Tile) {
	moveCost := 0
	switch tile.Attr {
	case Marsh:
		moveCost = 6
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

func TestMultiSearch(t *testing.T) {
	graph := NewWithTiles(`
s.......
xx..mmmm
....mmmm
.....rre
`)
	graph.SetCountCost(&MyCost{})
	graph.SearchMulti()

	result := []string{
		`
soo.....
xxo.mmmm
..o.mmmm
..oooooe
`,
		`
soo.....
xxo.mmmm
..oommmm
...ooooe
`,
		`
soo.....
xxoommmm
...ommmm
...ooooe
`,
		`
sooo....
xx.ommmm
...ommmm
...ooooe
`,
	}

	for i, p := range graph.GetMultiPath() {
		if result[i] != graph.GetStringTiles(p) {
			t.Error(graph.GetSteps(p))
			t.Error(graph.GetStringTiles(p))
		}

	}

	graph.Clear()

	graph.SetStringTiles(`
s......m
xx..mmmr
...rmmmr
..rrrrre
`)

	result = []string{
		`
soooooom
xx..mmoo
...rmmmo
..rrrrre
`,
		`
sooooooo
xx..mmmo
...rmmmo
..rrrrre
`,
	}

	graph.SearchMulti()
	for i, p := range graph.GetMultiPath() {
		if result[i] != graph.GetStringTiles(p) {
			t.Error(graph.GetSteps(p))
			t.Error(graph.GetStringTiles(p))
		}

	}
}

func TestMultiSearchDifferentSteps(t *testing.T) {
	graph := NewWithTiles(`
s..xmrrr
.x....xm
.xxxxxx.
..Mrr...
.xxxxxxe
`)
	graph.SetCountCost(&MyCost{})
	graph.SearchMulti()

	result := []string{
		`
s..xmrrr
ox....xm
oxxxxxx.
oooooooo
.xxxxxxe
`,
		`
sooxmooo
.xooooxo
.xxxxxxo
..Mrr..o
.xxxxxxe
`,
	}

	pl := graph.GetMultiPath()

	if graph.GetSteps(pl[0]) == graph.GetSteps(pl[1]) {
		t.Error(graph.GetSteps(pl[0]), graph.GetSteps(pl[1]))
	}

	for i, p := range pl {
		if result[i] != graph.GetStringTiles(p) {
			t.Error(graph.GetSteps(p))
			t.Error(graph.GetStringTiles(p))
		}
	}
}
