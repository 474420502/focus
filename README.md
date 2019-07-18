# structure

暂时没时间整理, 后期才整理完整

## PriorityQueue

``` golang
pq := pqueuekey.New(compare.Int)
pq.Push(1, 1)
pq.Push(4, 4)
pq.Push(5, 5)
pq.Push(6, 6)
pq.Push(2, 2)        // pq.Values() = [6 5 4 2 1]
value, _ := pq.Pop() // value = 6
t.Error(value)
value, _ = pq.Get(1)        // value = 1 pq.Values() = [5 4 2 1]
value, _ = pq.Get(0)        // value = nil , Get equal to Seach Key
value, _ = pq.Index(0)      // value = 5, compare.Int the order from big to small
values := pq.GetRange(2, 5) // values = [2 4 5]
values = pq.GetRange(5, 2)  // values = [5 4 2]
values3 := pq.GetAround(5) //  values3 = [<nil>, 5, 4]
```
