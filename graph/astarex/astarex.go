package astarex

import (
	"unsafe"

	"github.com/474420502/focus/tree/heap"
)

// Tile imp
type Tile interface {
	GetCost() int
	SetCost(int)

	GetWeight() int
	SetWeight(int)

	Key() []byte

	GetAttribute() interface{} //byte
}

// INeighbor the neighbor of tile
type INeighbor interface {
	GetNeighbor(astar *AstarEx, cur Tile) []Tile
}

// IHandler the Handle of tile
type IHandler interface {
	HandleAttribute(astar *AstarEx, current Tile) bool
}

// ICount the cost of tile
type ICount interface {
	CountNeighborCost(astar *AstarEx, neighbor, current Tile)
	CountNeighborWeight(astar *AstarEx, neighbor, current Tile)
}

// AstarEx astar interface
type AstarEx struct {
	start Tile
	end   Tile

	counted  map[string]bool
	openHeap *heap.Tree

	INeighbor
	IHandler
	ICount
}

// search astar search path
func (graph *AstarEx) search() bool {

	defer func() {
		graph.openHeap.Clear()
	}()

	if graph.start == nil {
		panic("not set start point")
	}

	if graph.end == nil {
		panic("not set end point")
	}

	// startTile := graph.Tiles[graph.start.Y][graph.start.X]
	// graph.start.Attr = startTile.Attr
	// startTile.IsCount = true
	// startTile.Attr = START

	// endTile := graph.Tiles[graph.end.Y][graph.end.X]
	// graph.end.Attr = endTile.Attr
	// endTile.IsCount = false
	// endTile.Attr = END

	ebkey := graph.end.Key()
	ekey := *(*string)(unsafe.Pointer(&ebkey))

	graph.openHeap.Put(graph.start)

	for {
		if itile, ok := graph.openHeap.Pop(); ok {
			tile := itile.(Tile)
			bkey := tile.Key()
			key := *(*string)(unsafe.Pointer(&bkey))

			if key == ekey {

				// graph.pathlist = make([]Path, 0)

				// var path Path
				// path = append(path, tile)
				// if multi {
				// 	graph.multiPath(tile, startTile, path)
				// 	sort.Sort(graph.pathlist)
				// } else {
				// 	graph.singlePath(tile, startTile, path)
				// }

				return true
			}

			for _, ntile := range graph.GetNeighbor(graph, tile) {

				if graph.HandleAttribute(graph, tile) {

					if _, ok := graph.counted[key]; !ok {
						graph.CountNeighborCost(graph, ntile, tile)
						graph.CountNeighborWeight(graph, ntile, tile)
						graph.counted[key] = true
						// 处理ntile权值
						graph.openHeap.Put(ntile)
					}
				}
			}

		} else {
			// log.Println("path can not found")
			break
		}
	}

	return false
}
