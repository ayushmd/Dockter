package main

import (
	"flag"
	"fmt"
	"sync"

	// "os"
	"github.com/ayush18023/Load_balancer_Fyp/builder"
	"github.com/ayush18023/Load_balancer_Fyp/master"
	"github.com/ayush18023/Load_balancer_Fyp/worker"
)

var Port int
var State string

func main() {
	var join string
	flag.IntVar(&Port, "port", 3000, "Port to serve")
	flag.StringVar(&State, "state", "", "Load balanced backends, use commas to separate")
	flag.StringVar(&join, "join", "", "Join a master")
	flag.Parse()
	fmt.Println("The port is: ", Port)
	fmt.Println("The state is: ", State)
	if join != "" {
		// spin up worker or builder
		var waitgrp sync.WaitGroup
		waitgrp.Add(1)
		go func() {
			defer waitgrp.Done()
			if State == "WORKER" {
				worker.NewWorkerServer(Port)
			} else {
				builder.NewBuilderServer(Port)
			}
		}()
		worker.Worker_.JoinMaster(join)
		waitgrp.Wait()
	} else {
		// spin up master
		master.NewMasterServer(Port)
	}
	// if len(serverList) == 0 {
	// 	log.Fatal("Please provide one or more backends to load balance")
	// }

	// pars, _ := url.Parse("localhost:6000")
	// master.Master_.ServerPool = append(master.Master_.ServerPool, &master.Backend{
	// 	URL:            pars,
	// 	State:          State,
	// 	IsAlive:        true,
	// 	CurrentConnect: 1,
	// 	CpuUsage:       20.04,
	// 	MemUsage:       4.8,
	// 	DiskUsage:      4.56,
	// })

	// EchoExchange()
	// for _, tok := range tokens {

	// }
	// fmt.Println("hello")
	// args := os.Args[1:]
	// var app *gin.Engine
	// if args[0] == "master" {
	// 	app = master.CreateServer()
	// } else if args[0] == "worker" {
	// 	app = worker.CreateServer()
	// } else if args[0] == "builder" {
	// 	app = builder.CreateServer()
	// }
	// app.Run()
	// write := utils.KafkaUPAuthWriter("build")
	// read := utils.KafkaUPAuthReader("build", "mygroup")
	// defer write.Close()
	// k := utils.Kafka{
	// 	Writer: write,
	// 	Reader: read,
	// }
	// defer k.CloseReader()
	// defer k.CloseWriter()
	// // for i := 0; i < 10; i++ {
	// // 	k.Write(
	// // 		[]byte("TESTKEY"+fmt.Sprint(i)),
	// // 		[]byte("TESTVALUE"+fmt.Sprint(i)),
	// // 	)
	// // }
	// k.ReaderServer(
	// 	Caller,
	// 	ErrorCallback,
	// 	int64(1),
	// )
	// s := middleware.NewMiddlewareInstance(3000)
	// s.ListenAndServe()
	// worker.NewWorkerServer(8000)
}
