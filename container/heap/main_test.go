package heap

import (
    "fmt"
    "testing"
    "github.com/BismarckDD/golang/sort"
)

type IntHeap []int
func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) {
    *h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
    old := *h
    size := len(old)
    x := old[size - 1]
    *h = old[0 : size - 1]
    return x
}

func (heap IntHeap) verify(t *testing.T, i int) {
    t.Helper()
    limit := heap.Len()
    j1 := (i << 1) + 1
    j2 := (i << 1) + 2
    if j1 < limit {
        if heap.Less(j1, i) {
            t.Errorf("heap invariant invalidated [%d] = %d > [%d] = %d", i, heap[i], j1, heap[j1])
            return
        }
        heap.verify(t, j1)
    }
    if j2 < limit {
        if heap.Less(j2, i) {
            t.Errorf("heap invariant invalidated [%d] = %d > [%d] = %d", i, heap[i], j1, heap[j2])
            return
        }
        heap.verify(t, j2)
    }
}

type IntGHeap []int
func (h IntGHeap) Len() int           { return len(h) }
func (h IntGHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntGHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntGHeap) Push(x interface{}) {
    *h = append(*h, x.(int))
}

func (h *IntGHeap) Pop() interface{} {
    old := *h
    size := len(old)
    x := old[size - 1]
    *h = old[0 : size - 1]
    return x
}

func (heap IntGHeap) verify(t *testing.T, i int) {
    t.Helper()
    limit := heap.Len()
    j1 := (i << 1) + 1
    j2 := (i << 1) + 2
    if j1 < limit {
        if heap.Less(i, j1) {
            t.Errorf("heap invariant invalidated [%d] = %d > [%d] = %d", i, heap[i], j1, heap[j1])
            return
        }
        heap.verify(t, j1)
    }
    if j2 < limit {
        if heap.Less(i, j2) {
            t.Errorf("heap invariant invalidated [%d] = %d > [%d] = %d", i, heap[i], j1, heap[j2])
            return
        }
        heap.verify(t, j2)
    }
}

func TestMain(t *testing.T) {

    ints := []int { 10, 9, 8, 7, 6, 5, 4, 3, 2, 1 }
    sort.SortInt(ints)
    fmt.Println(ints)

    // array := [10]int { 10, 9, 8, 7, 6, 5, 4, 3, 2, 1 }
    // sort.Ints([]int(array))
    // fmt.Println(array)

    strings := []string { "Hello World", "Hello Go", "ABCDEFG" }
    sort.SortString(strings)
    fmt.Println(strings)

    // list := containter.list

    h := &IntHeap{2, 1, 5, 100, 3, 6, 4, 5}
    Init(h)
    Push(h, 3)
    Fix(h, 3)
    fmt.Printf("minimum: %d\n", (*h)[0])
    for h.Len() > 0 {
        fmt.Printf("%d ", Pop(h))
    } 
    fmt.Println("")

    g := &IntGHeap{2, 1, 5, 100, 3, 6, 4, 5}
    Init(g)
    Push(g, 3)
    Fix(g, 3)
    fmt.Printf("maximum: %d\n", (*g)[0])
    for g.Len() > 0 {
        fmt.Printf("%d ", Pop(g))
    }
    fmt.Println("")
}

