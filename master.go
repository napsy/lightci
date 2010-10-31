package main

import (
    "fmt"
    "netchan"
)

const (
    MsgServerInfo = iota
    MsgAddProject = iota
    MsgQueryStatus = iota
    MsgForceBuild = iota
)

type Message struct {
    Type int
}

func RunChannel() *chan int{
    exporter, err := netchan.NewImporter("tcp", "127.0.0.1:7000")
    s := make(chan int)
    if err != nil {
        panic(err.String())
    }
    srv_netchan := make(chan Message)
    err = exporter.Import("LCIMessenger", srv_netchan, netchan.Send)
    if err != nil {
        panic(err.String())
    }
    go func() {
        for {
            msg := Message{MsgServerInfo }
            srv_netchan <- msg
            break
        }
    }()
    return &s
}

func main() {
    fmt.Println("lightci - master")
    s := RunChannel()
    <-*s
}
