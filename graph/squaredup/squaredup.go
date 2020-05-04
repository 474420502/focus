package squaredup

import (
	"log"
	"math/rand"
	"time"

	"github.com/474420502/focus/tree/heap"
)

// SquaredUp map
type SquaredUp struct {
	start *Grid
	end   *Grid

	dimY, dimX int
	vlen       int

	steps int

	openHeap *heap.Tree
	counted  map[string]bool
}

func weightCompare(x1, x2 interface{}) int {
	p1, p2 := x1.(*Grid), x2.(*Grid)
	if p1.weight > p2.weight { // 权重大的优先
		return 1
	}
	return -1
}

// New create
func New(dimY, dimX int) *SquaredUp {
	su := &SquaredUp{dimY: dimY, dimX: dimX}
	su.vlen = 1
	su.openHeap = heap.New(weightCompare)
	su.counted = make(map[string]bool)
	// su.grid = NewGird(dimY, dimX, su.vlen)
	return su
}

func (su *SquaredUp) randomGird(gird *Grid, count int) {
	rand.Seed(time.Now().UnixNano())
	var move = [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for ; count > 0; count-- {

		offset := move[rand.Intn(4)]
		x, y := gird.X+offset[0], gird.Y+offset[1]
		if x < 0 || y < 0 || x >= su.dimX || y >= su.dimY {
			continue
		}
		gird.SwapValue(su.dimX, su.vlen, x, y)
	}
}

// Search search path
func (su *SquaredUp) Search() {
	// 1 2 3
	// 4 5 6
	// 7 8 0
	su.steps = 0
	su.openHeap.Put(su.start)

	for {
		if igird, ok := su.openHeap.Pop(); ok {

			gird := igird.(*Grid)

			if gird.Compare(su.end) == 0 {
				content := gird.GetGirdString(su.dimY, su.dimX, su.vlen)
				log.Println(content, "\n", "steps", su.steps)
				break
			}

			// Left
			su.steps++
			// content := gird.GetGirdString(su.dimY, su.dimX, su.vlen)
			// log.Println(content)

			for _, offset := range [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {

				x, y := gird.X+offset[0], gird.Y+offset[1]

				if x < 0 || y < 0 || x >= su.dimX || y >= su.dimY {
					continue
				}

				dst := gird.Clone()
				dst.SwapValue(su.dimX, su.vlen, x, y)

				key := string(dst.GetValues())
				if _, ok := su.counted[key]; !ok {
					su.counted[key] = true

					// cost
					// weight
					dst.cost = gird.cost + 1
					for i, v := range dst.bits {
						if su.end.bits[i] == v {
							dst.weight += 16
						}
					}

					dst.weight -= dst.cost

					su.openHeap.Put(dst)
				}

				// if ntile.IsCount == false && ntile.Attr != BLOCK {
				// 	graph.countCost.Cost(graph, ntile, tile)
				// 	graph.countWeight.Weight(graph, ntile, tile)
				// 	ntile.IsCount = true
				// 	// 处理ntile权值
				// 	graph.openHeap.Put(ntile)
				// }
			}

		} else {
			break
		}
	}
}
