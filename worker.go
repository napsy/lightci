package main

import (
    "fmt"
    "netchan"
)

type Message struct {
    Type int
}

func RunChannel() *chan int{

    exporter, err := netchan.NewExporter("tcp", "127.0.0.1:7000")
    s := make(chan int)
    if err != nil {
        panic(err.String())
    }
    srv_netchan := make(chan Message)
    err = exporter.Export("LCIMessenger", srv_netchan, netchan.Recv)
    if err != nil {
        panic(err.String())
    }
    go func() {
        for {
            msg := <-srv_netchan
            fmt.Printf("Received message, type = %d\n", msg.Type)
            msg_resp := Message{10}
            srv_netchan <- msg_resp

        }
    }()
    return &s
}

func main() {
    fmt.Println("lightci worker")
    s := RunChannel()
    <-*s

}
