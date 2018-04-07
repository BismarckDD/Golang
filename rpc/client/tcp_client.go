package main

import (
    "fmt"
    "log"
    "net/rpc"

    "github.com/BismarckDD/golang/test"
)

var (
    serverAddress = "localhost"
)

func main() {


    client, client_err := rpc.Dial("tcp", serverAddress + ":1234")
    if client_err != nil {
        log.Fatal("client error: ", client_err)
    }

    args := &test.Args{7, 8}

    log.Printf("Connection has been established, start to call rpc service.")
    var reply int
    call_err := client.Call("Arith.Multiply", args, &reply)
    if call_err != nil {
        log.Fatal("call error: ", call_err)
    }
    fmt.Printf("Arith %d * %d = %d\n", args.A, args.B, reply)

}
