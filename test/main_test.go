package test

import (
    "testing"
    "fmt"
)

const (
    FROM = 1
    TO = 3
    MED = 2
    LAYER = 7
    N = 8 
)

var count int
type Step struct {
    from, to int
}
var steps [65536]Step

func Hanoi(layer uint, from, to, med int) {

    if layer > 1 {
        Hanoi(layer - 1, from, med, to)
    }

    fmt.Printf("From %d To %d\n", from, to)
    count++
    var step Step
    step.from = from
    step.to = to
    steps[count] = step

    if layer > 1 {
        Hanoi(layer - 1, med, to, from)
    }
}

func HanoiSearch(layer uint, from, to, med, i int) Step  {

    var step Step
    if (1 << layer) <= i {
        step.from = 0
        step.to = 0
        return step
    }

    if i == (1 << (layer - 1)) {
        var step Step
        step.from = from
        step.to = to
        return step
    } else if i > (1 << (layer - 1)) {
        temp := i - (1 << (layer - 1))
        return HanoiSearch(layer - 1, med, to, from, temp)
    } else if i < (1 << (layer - 1)) {
        temp := i
        return HanoiSearch(layer - 1, from, med, to, temp)
    } else {
        return step
    }
}

func TestHanoi(t *testing.T) {
    count = 0
    Hanoi(LAYER, FROM, TO, MED)
    fmt.Printf("Total %d\n", count)
    for i := 1; i < (1 << LAYER); i++ {
        step := HanoiSearch(LAYER, FROM, TO, MED, i)
        if step.from != steps[i].from || step.to != steps[i].to {
            t.Error()
        }
    }
}

var dia []bool
var ind []bool
var ver []bool
var ans map[int]int
func uncover(p_row, p_col, p_N int) {
    dia[p_col - p_row + p_N] = false
    ind[p_col + p_row] = false
    ver[p_col] = false
}

func cover(p_row, p_col, p_N int) {
    dia[p_col - p_row + p_N] = true
    ind[p_col + p_row] = true
    ver[p_col] = true
}

func validate(p_row, p_col, p_N int) bool {
    return !(dia[p_col - p_row + p_N] || ind[p_col + p_row] || ver[p_col])
}
func answer(p_N int) {
    fmt.Printf("The %dth solution:\n", count)
    for i := 0; i < p_N; i++ {
        for j := 0; j < p_N; j++ {
            if (ans[i] == j) {
                print("Q ")
            } else {
                print(". ")
            }
        }
        print("\n");
    }
}

func Queen(p_layer, p_N int) {

    if p_layer == p_N {
        count++
        answer(p_N)
        return
    }

    for i := 0; i < p_N; i++ {
        if validate(p_layer, i, p_N) {
            cover(p_layer, i, p_N)
            ans[p_layer] = i
            Queen(p_layer + 1, p_N)
            // delete(ans, p_layer)
            uncover(p_layer, i, p_N)
        }
    }
}

func TestQueen(t *testing.T) {
    dia = make([]bool, (N << 1))
    ind = make([]bool, (N << 1))
    ver = make([]bool, N)
    ans = make(map[int]int)
    count = 0
    fmt.Println(dia)
    fmt.Println(ind)
    fmt.Println(ver)
    Queen(0, N)
    fmt.Printf("Total: %d\n", count)
}


func Return(i, j int) (h, l int) {
    // 不用return返回值依然能够return
    h, l = j, i
    return 
}

func TestReturn(t *testing.T) {

    i, j := Return(100, 2000)
    fmt.Printf("%d, %d\n", i, j)
}
