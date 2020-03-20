# structure

暂时没时间整理, 后期才整理完整.

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

    iter := pq.Iterator() // Next 大到小 从root节点起始
    log.Println(pq.String())
    // log.Println(iter.Value()) 直接使用会报错,
    iter.ToHead()
    iter.Next()
    log.Println(iter.Value())              // 起始最大值. true 5
    log.Println(iter.Prev(), iter.Value()) // false 5

    // Prev 大到小
    log.Println(iter.Next(), iter.Value()) // true 4

}
```
