package main

import (
	"time"

	"github.com/nitish2580/lambda/network"
)

//server
//Transport =>tcp,udp
// Block
// tx
//keypair

func main() {
	trLocal := network.NewLocalTransport("Local")
	trRemote := network.NewLocalTransport("REMOTE")

	trLocal.Connect(trRemote)
	trRemote.Connect(trLocal)

	go func(){
		for {
		trRemote.SendMessage(trLocal.Addr(),[]byte("Hello world"))
		time.Sleep(1 * time.Second)
		}
	}()

	opts := network.ServerOpts{
		Transports: []network.Transport{trLocal},
	}

	s:= network.NewServer(opts)
	s.Start()
}