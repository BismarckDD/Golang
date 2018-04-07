package main

import (
    "log"
    "net"
    "net/http"
    "net/rpc"

    . "github.com/BismarckDD/golang/test"
)

func main() {

    // Arith有两个接口
    arith := new(Arith)
    // arith1 := new(Arith)
    // arith2 := new(Arith)
    // 注册对象到rpc上
    rpc.Register(arith)
    rpc.Register(arith1)
    rpc.Register(arith2)
    // rpc接收http服务
    rpc.HandleHTTP()

    // listener监听tcp:1234端口
    listener, server_err := net.Listen("tcp", ":1234")
    if server_err != nil {
        log.Fatal("listen error: ", server_err)
    }

    // 开球一个新的gorountine
    go http.Serve(listener, nil)

    for i := 1; i == 1; {
    }

}
