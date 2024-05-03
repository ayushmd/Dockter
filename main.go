package main

import (
	"flag"
	"fmt"
	"log"
	"sync"

	"github.com/ayush18023/Load_balancer_Fyp/builder"
	"github.com/ayush18023/Load_balancer_Fyp/internal"
	"github.com/ayush18023/Load_balancer_Fyp/internal/auth"
	"github.com/ayush18023/Load_balancer_Fyp/master"
	"github.com/ayush18023/Load_balancer_Fyp/worker"
	"github.com/joho/godotenv"
)

// "os"

var Port int
var State string
var DefaultPort = 3210

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	// builder.CreateImage("ayush", "ayush")
	// builder.BuildContainer("ayush")
	// fmt.Println(builder.FindPort("ayush"))
	// doc := internal.Dockter{}
	// doc.Init()
	// defer doc.Close()
	// doc.BuildNewImage()
	// fmt.Println("Builded")
	// outputlines := doc.ExecuteCommand("ayush", []string{"netstat", "-tuln"})
	// port := internal.FindFirstPort(outputlines)
	// fmt.Println(port)
	// Name := "ayush"
	// GitLink := "https://github.com/johnpapa/node-hello.git"
	// Branch := "master"
	// BuildCmd := "npm i"
	// RunCmd := "npm start"
	// RuntimeEnv := "Node"
	// doc := internal.Dockter{}
	// doc.Init()
	// defer doc.Close()
	// // fmt.Println(doc.ListContainers())
	// containerJSon, _ := doc.ContainerMetrics("962bde4d3c8f25a9b4efdcbce351768dfda657cfd93371e89b7290bfea2897f3")
	// fmt.Println(containerJSon.Memory_stats)
	// fmt.Println(containerJSon.Cpu_stats)
	// iostream.PrintObj(containerJSon.Memory_stats)
	// iostream.PrintObj(containerJSon.Cpu_stats)
	// doc.BuildImage(internal.BuildOptions{
	// 	Context:              GitLink,
	// 	Label:                Name,
	// 	GitBranch:            Branch,
	// 	UseDefaultDockerFile: false,
	// 	DockerfileContent:    builder.Builder_.BuildDockerLayers(Name, BuildCmd, RunCmd, RuntimeEnv, nil),
	// })
	// tag, port, err := builder.Builder_.BuildRaw(
	// 	Name,
	// 	GitLink,
	// 	Branch,
	// 	BuildCmd,
	// 	RunCmd,
	// 	RuntimeEnv,
	// 	nil,
	// )
	// if port != "3000" || err != nil {
	// 	fmt.Printf(`BuildRaw = %q, %v, want 3000, nil`, tag, port, err)
	// }
	// fmt.Println(builder.FindPort(Name))
	var join string
	var generateToken bool
	flag.IntVar(&Port, "port", DefaultPort, "Port to serve")
	flag.StringVar(&State, "state", "MASTER", "Load balanced backends, use commas to separate")
	flag.StringVar(&join, "join", "", "Join a master")
	flag.BoolVar(&generateToken, "generatetoken", false, "Gives the token to join master")
	flag.Parse()
	// fmt.Println("The port is: ", Port)
	// fmt.Println("The state is: ", State)
	if !generateToken {
		var waitgrp sync.WaitGroup
		waitgrp.Add(1)
		if State == "MASTER" {
			go func() {
				defer waitgrp.Done()
				master.NewMasterServer(Port)
			}()
		} else {
			var addr string
			if join != "" {
				addr = auth.ParseAccessToken(join).Address
			}
			// fmt.Println(addr)
			if State == "WORKER" {
				go func() {
					defer waitgrp.Done()
					worker.NewWorkerServer(Port)
				}()
				if join != "" {
					worker.Worker_.JoinMaster(addr)
				}
			} else if State == "BUILDER" {
				go func() {
					defer waitgrp.Done()
					builder.NewBuilderServer(Port)
				}()
				if join != "" {
					builder.Builder_.JoinMaster(addr)
				}
			} else if State == "client" {
				Client()
			} else {
				log.Fatal("State not recognized")
			}
		}
		waitgrp.Wait()
	} else {
		myip, err := internal.GetMyIP()
		if err != nil {
			log.Fatal(err)
		}
		addr := fmt.Sprintf("%s:%d", myip, DefaultPort)
		tokenClaim := auth.MasterClaim{
			Address: addr,
		}
		authToken, err := auth.NewAccessToken(tokenClaim)
		if err != nil {
			log.Fatal("error in genrating jwt")
		}
		fmt.Println("The Generated AuthToken:")
		fmt.Println(authToken)
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
