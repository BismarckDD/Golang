package main

import (
    "fmt"
    "errors"
)


func div(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}


func test(x int) func() {
    return func() {
        fmt.Println(x)
    }
}


func test2(a, b int) {
    defer fmt.Println("dispose...")
    fmt.Println(a / b)
}


type user struct {
    name string
    age byte
}


type manager struct {
    user
    title string
}


type X int

func (x *X) inc() {
    *x++
}



func main() {

    x := []int {100, 101, 102}

    for i, n := range x {
        fmt.Println(i, ": " , n)
    }

    a, b := 10, 2
    c, err := div(a, b)
    fmt.Println(c, err)

    xx := 100
    f := test(xx)
    f()


    // test2(10, 0)

    // delete x
    mm := make(map[string]int)
    mm["a"] = 1
    xxx, ok := mm["a"]
    fmt.Println(xxx, ok)
    delete(mm, "a")


    var m manager
    m.name = "Tome"
    m.age = 29
    m.title = "CTO"

    fmt.Println(m)


    var y X
    y.inc()
    fmt.Println(y)
}
