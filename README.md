# structure

there is a lot structure that easy by used

## PriorityQueue

``` golang
package main

import (
    "log"

    "github.com/474420502/focus/compare"
    pqueuekey "github.com/474420502/focus/priority_queuekey"
)

func main() {
    pq := New(compare.Int)
    pq.Push(1, 1)
    pq.Push(4, 4)
    pq.Push(5, 5)
    pq.Push(6, 6)
    pq.Push(2, 2) // pq.Values() = [6 5 4 2 1]
    log.Println(pq.Values())
    value, _ := pq.Pop() // value = 6
    log.Println(value)
    value, _ = pq.Get(1) // value = 1 pq.Values() = [5 4 2 1]
    log.Println(value)
    value, _ = pq.Get(0) // value = nil , Get equal to Seach Key
    log.Println(value)
    value, _ = pq.Index(0) // value = 5, compare.Int the order from big to small
    log.Println(value)
    values := pq.GetRange(2, 5) // values = [2 4 5]
    log.Println(values)
    values = pq.GetRange(5, 2) // values = [5 4 2]
    log.Println(values)
    values = pq.GetRange(100, 2) // values = [5 4 2]
    log.Println(values)
    values3 := pq.GetAround(5) // values3 = [<nil>, 5, 4]
    log.Println(values3)

    iter := pq.Iterator() // Next big -> small 
    log.Println(pq.String())
    iter.ToHead()
    iter.Next()
    log.Println(iter.Value())              // Head is maxvalue. true 5
    log.Println(iter.Prev(), iter.Value()) // false 5

    // Prev big -> small
    log.Println(iter.Next(), iter.Value()) // true 4

}
```

## Astar

* Simple

``` golang
package main

import (
    "log"

    "github.com/474420502/focus/graph/astar"
)

func main() {
    a := astar.New(5, 5)
    a.SetTarget(0, 0, 5-1, 5-1)
    if a.Search() {
        log.Println(a.GetSinglePathTiles())
        // soo..
        // ..o..
        // ..o..
        // ..ooo
        // ....e
        for _, p := range a.GetPath() {
            log.Println(p.X, p.Y) // End Point -> Start Point
        }
    }

    a = astar.NewWithTiles(`
    s....
    .xxx.
    .xxx.
    .xxx.
    ....e
    `)
    if a.SearchMulti() { // get multi the path of same cost
        for _, p := range a.GetMultiPath() {
            log.Println(a.GetPathTiles(p))
            // path 1:
            // s....
            // oxxx.
            // oxxx.
            // oxxx.
            // ooooe

            // path 2:
            // soooo
            // .xxxo
            // .xxxo
            // .xxxo
            // ....e
        }
    }
}

```

* custom

```golang
package main

import (
    "log"

    "github.com/474420502/focus/graph/astar"
)

type MyCost struct {
}

const (
    MARSH    = byte('M')
    MOUNTAIN = byte('m')
    RIVER    = byte('r')
)

// Cost ca
func (cost *MyCost) Cost(graph *astar.Graph, tile *astar.Tile, ptile *astar.Tile) {
    moveCost := 0
    switch tile.Attr {
    case MARSH:
        moveCost = 6
    case MOUNTAIN:
        moveCost = 3
    case astar.PLAIN:
        moveCost = 1
    case RIVER:
        moveCost = 2
    }
    tile.Cost = ptile.Cost + moveCost
}

func main() {
    a := astar.NewWithTiles(`
    s..xmrrr
    .x....xm
    .xxxxxx.
    ..Mrr...
    .xxxxxxe
    `)
    a.SetCountCost(&MyCost{})
    a.SearchMulti()

    // result := []string{
    //     `
    // s..xmrrr
    // ox....xm
    // oxxxxxx.
    // oooooooo
    // .xxxxxxe
    // `,
    //     `
    // sooxmooo
    // .xooooxo
    // .xxxxxxo
    // ..Mrr..o
    // .xxxxxxe
    // `,
    // }

    pl := a.GetMultiPath()

    log.Println(a.GetSteps(pl[0]), a.GetSteps(pl[1])) // 12 step 14 step, but they cost is equal

    for _, p := range pl {
        log.Println(a.GetSteps(p))
        log.Println(a.GetPathTiles(p))
    }
}

```
