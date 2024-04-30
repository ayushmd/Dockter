package internal

import (
	"context"
	"crypto/tls"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ayush18023/Load_balancer_Fyp/internal/auth"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/scram"
)

type KafkaWriter struct {
	Writer *kafka.Writer
}

type KafkaReader struct {
	Reader *kafka.Reader
}

func KafkaUPAuthWriter(topic string) *kafka.Writer {
	mechanism, err := scram.Mechanism(
		scram.SHA256,
		auth.GetKey("UPSTASH_KAFKA_USER"),
		auth.GetKey("UPSTASH_KAFKA_PASS"),
	)
	if err != nil {
		panic(err)
	}
	return &kafka.Writer{
		Addr: kafka.TCP(
			auth.GetKey("UPSTASH_KAFKA_ADDR"),
		),
		Topic: topic,
		Transport: &kafka.Transport{
			SASL: mechanism,
			TLS:  &tls.Config{},
		},
	}
}

func (k *KafkaWriter) Write(key, value []byte) {
	k.Writer.WriteMessages(context.Background(), kafka.Message{
		Key:   key,
		Value: value,
	})
}

func KafkaUPAuthReader(topic, group string) *kafka.Reader {
	mechanism, err := scram.Mechanism(
		scram.SHA256,
		auth.GetKey("UPSTASH_KAFKA_USER"),
		auth.GetKey("UPSTASH_KAFKA_PASS"),
	)
	if err != nil {
		panic(err)
	}
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"close-lynx-12441-us1-kafka.upstash.io:9092"},
		GroupID: group,
		Topic:   topic,
		Dialer: &kafka.Dialer{
			SASLMechanism: mechanism,
			TLS:           &tls.Config{},
		},
	})
}

func (k *KafkaReader) ReaderServer(
	callback func(kafka.Message),
	onError func(),
	concurrentLimit int64,
) {
	// fmt.Println(concurrentLimit)
	var counter int64 = int64(0)
	var mu sync.Mutex // Mutex for synchronizing access to counter

	for {
		mu.Lock()
		if counter < concurrentLimit {
			atomic.AddInt64(&counter, 1)
			mu.Unlock()
			go func() {
				// defer func() {
				// 	}()
				message, err := k.Reader.ReadMessage(context.Background())
				fmt.Printf("%s recieved", string(message.Key))
				if err != nil {
					onError()
				}
				callback(message)
				atomic.AddInt64(&counter, -1)
			}()
		} else {
			mu.Unlock()
			time.Sleep(time.Millisecond * 100) // Wait a bit if limit is reached
		}
	}
}

func (k *KafkaReader) CloseReader() {
	k.Reader.Close()
}

func (k *KafkaWriter) CloseWriter() {
	k.Writer.Close()
}
