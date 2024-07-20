package main

import (
	"log"
	"time"

	"github.com/fatedier/frp/cmd/frp"
)

type LogListener struct{}

func (*LogListener) Log(log string) {
}

func main() {
	go func() {
		frp.RunFrps("/Users/haidy/GoProjects/frp/conf/frps.toml", true)
		log.Println("frps finished")
	}()

	time.Sleep(1 * time.Second)

	// go func() {
	// 	frp.RunFrpc("", "/Users/haidy/GoProjects/frp/conf/frpc.toml")
	// }()

	// time.Sleep(3 * time.Second)
	// go func() {
	// 	frp.StopFrpc()
	// }()

	time.Sleep(3 * time.Second)
	go func() {
		frp.StopFrps()
	}()

	time.Sleep(1 * time.Second)
}
