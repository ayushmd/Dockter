package master

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"

	"github.com/ayush18023/Load_balancer_Fyp/internal"
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
	var waitgrp sync.WaitGroup
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal("Master server not started")
	}
	waitgrp.Add(3)
	master = &MasterServer{
		grpcServer: NewMasterGrpcInstance(),
		httpServer: NewMasterHttpInstance(port + 1),
		kReader: &internal.KafkaReader{
			Reader: internal.KafkaUPAuthReader(
				internal.GetKey("UPSTASH_KAFKA_TOPIC"),
				internal.GetKey("UPSTASH_KAFKA_GROUP"),
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
	waitgrp.Wait()
}
