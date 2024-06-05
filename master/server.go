package master

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/ayush18023/Load_balancer_Fyp/ca"
	"github.com/ayush18023/Load_balancer_Fyp/internal"
	"github.com/ayush18023/Load_balancer_Fyp/internal/auth"
	"github.com/ayush18023/Load_balancer_Fyp/internal/sqlite"
	lru "github.com/hashicorp/golang-lru/v2"
	"google.golang.org/grpc"
)

type MasterServer struct {
	grpcServer *grpc.Server
	httpServer *http.Server
	kReader    *internal.KafkaReader
	// cached
}

var master *MasterServer

func NewMasterServer(port int) {
	cache, err := lru.New[string, Task](128)
	if err != nil {
		log.Fatal("Master server not started")
	}
	Master_.kwriter = &internal.KafkaWriter{
		Writer: internal.KafkaUPAuthWriter("build"),
	}
	Master_.dnsStatus = RUNNING
	Master_.dbDns, err = sqlite.CreateConn()
	if err != nil {
		Master_.dnsStatus = EXITED
	} else {
		Master_.dnsStatus = STARTED
	}
	Master_.cacheDns = cache
	var waitgrp sync.WaitGroup
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal("Master server not started")
	}
	waitgrp.Add(4)
	master = &MasterServer{
		grpcServer: NewMasterGrpcInstance(),
		httpServer: NewMasterHttpInstance(port + 1),
		kReader: &internal.KafkaReader{
			Reader: internal.KafkaUPAuthReader(
				auth.GetKey("UPSTASH_KAFKA_TOPIC"),
				auth.GetKey("UPSTASH_KAFKA_GROUP"),
			),
		},
	}
	go func() {
		defer waitgrp.Done()
		err := master.grpcServer.Serve(lis)
		if err != nil {
			log.Fatal("Master GRPC not started")
		}
	}()
	fmt.Println("Master GRPC server started on port", port)
	go func() {
		defer waitgrp.Done()
		err := master.httpServer.ListenAndServe()
		if err != nil {
			log.Fatal("Master HTTP not started")
		}
	}()
	fmt.Println("Master HTTP server started on port", port+1)
	go func() {
		defer waitgrp.Done()
		master.kReader.ReaderServer(
			Master_.KafkaHandler,
			Master_.KafkaError,
			1,
		)
	}()
	fmt.Println("Kafka Listner started")

	workers := &internal.Background{
		Callback: Master_.Pool,
		Timer:    10 * time.Second,
	}
	go func() {
		defer waitgrp.Done()
		workers.Run()
	}()
	fmt.Println("Background workers started")
	waitgrp.Wait()
	BootMasterServices()
}

func StartCAService(wg *sync.WaitGroup) {
	defer wg.Done()
	var port string = fmt.Sprintf(":%d", 2222)
	var lis net.Listener
	lis, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("CA Service not started")
	}
	ins := ca.NewCAInstance()
	err = ins.Serve(lis)
	if err != nil {
		fmt.Println("CA Service not started")
	}
	Master_.JoinService(port, "CA")
}

func BootMasterServices() {
	ms := []string{"CA"} //read from config
	wg := &sync.WaitGroup{}
	for _, service := range ms {
		wg.Add(1)
		if service == "CA" {
			go StartCAService(wg)
		}
	}
	wg.Wait()
}
