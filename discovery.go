package main

import (
	"log"
	"net/url"
	"time"
)

func EchoExchange() {
	if len(MyMeta.peers) == 0 {
		ser := InitEchoServer(Port)
		defer ser.GracefulStop()
		InitDiscovery()
	}
}

func InitDiscovery() {
	for _, server := range ServerList {
		serverURL, err := url.Parse(server)
		if err != nil {
			log.Fatal(err)
		}
		MyMeta.AppendPeer(serverURL, State(""))
	}
	retries := 0
	allReady := false
	for retries < 3 && !allReady {
		if DiscoveredAll() {
			allReady = true
		} else {
			time.Sleep(10 * time.Second)
		}
	}
}

func DiscoveredAll() bool {
	var successEcho bool = true
	for i, peer := range MyMeta.peers {
		res, err := GrpcEchoPool(peer.URL.RawPath)
		if err != nil {
			successEcho = false
		} else {
			MyMeta.peers[i].state = State(res.State)
		}
	}
	return successEcho
}
