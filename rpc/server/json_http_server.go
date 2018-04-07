package main

import (
    "log"
    "net"
    "net/rpc"
    "net/rpc/jsonrpc"
    . "github.com/BismarckDD/golang/test"
)

func main() {

    // Arith有两个接口
    arith := new(Arith)
    // 注册对象到rpc上
    // server := rpc.NewServer()
    // server.Register(arith)
    rpc.Register(arith)

    // listener监听tcp:1234端口
    listener, listener_err := net.Listen("tcp", ":1234")
    if listener_err != nil {
        log.Fatal("listen error: ", listener_err)
    }

    go func() {
        conCount := 0
        for {
            conn, conn_err := listener.Accept()
            if conn_err != nil {
                log.Fatal("Accept Error:,", conn_err)
                continue
            }
            conCount++
            log.Printf("Receive Client Connection %d\n", conCount)
            go jsonrpc.ServeConn(conn)
        }
    }()
    // 开球一个新的gorountine

    dead_loop := make(chan int)
    i, ok := <-dead_loop 
    if !ok {
        println(i)
    }
    defer close(dead_loop)
}
