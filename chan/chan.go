package main
// main function must be placed within main package
// in each folder, there must no be more than one file containing main function.

import (
    "fmt"
    "time"
)

func sendChan(cl chan int, len int) {
    fmt.Println("sendChan_enter")
    for i := 0; i < len; i++ {
        cl <- i                  //cl的存储第4个数据的时候，会阻塞当前goruntine，直到其它goruntine取走一个或多个数据
        fmt.Println("# ", i)
    }
    fmt.Println("sendChan_end")
}

func getChan(cl chan int, len int) {
    fmt.Println("getChan_enter")
    for i := 0; i < len; i++ {
        data := <-cl
        fmt.Println("$ ", data)  //当cl的数据为空时，阻塞当前goruntine，直到新的数据写入cl
    }
    fmt.Println("getChan_end")
}

func main() {
    cl := make(chan int, 3)      // 写入3个元素都不会阻塞当前goroutine, 存储个数达到4的时候会阻塞

    go sendChan(cl, 10)
    go getChan(cl, 11)

    time.Sleep(time.Second)
}
