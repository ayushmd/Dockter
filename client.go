package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

func makeRawRequest(name string) {
	url := "http://ayushd.cloud/api/buildraw"
	body := fmt.Sprintf(`{"name":"%s","gitlink":"https://github.com/johnpapa/node-hello.git","branch":"master","buildCmd":"npm i","startCmd":"npm start","runtimeEnv":"Node","runningPort":"3000"}`, name)
	// fmt.Println(body)
	jsonStr := []byte(body)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
}

func makeImageRequest(name, image, port string) {
	url := "http://ayushd.cloud/api/deploy"
	body := fmt.Sprintf(`{"name":"%s","dockerImage":"%s","runningPort":"%s"}`, name, image, port)
	// fmt.Println(body)
	jsonStr := []byte(body)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
}

func makeNgxRequest(wg *sync.WaitGroup) {
	for i := 0; i < 1; i++ {
		wg.Add(1)
		name := fmt.Sprintf("ngx%d", i)
		go func(name string) {
			defer wg.Done()
			makeImageRequest(name, "nginx", "80")
		}(name)
		time.Sleep(5 * time.Second)
	}
}

func makeGrfRequest(wg *sync.WaitGroup) {
	for i := 0; i < 1; i++ {
		wg.Add(1)
		name := fmt.Sprintf("grf%d", i)
		go func(name string) {
			defer wg.Done()
			makeImageRequest(name, "grafana/grafana", "3000")
		}(name)
		time.Sleep(5 * time.Second)
	}
}

func makeZkRequest(wg *sync.WaitGroup) {
	for i := 0; i < 1; i++ {
		wg.Add(1)
		name := fmt.Sprintf("zk%d", i)
		go func(name string) {
			defer wg.Done()
			makeImageRequest(name, "zookeeper", "2181")
		}(name)
		time.Sleep(5 * time.Second)
	}
}

func Client() {
	wg := &sync.WaitGroup{}
	wg.Add(3)
	go makeNgxRequest(wg)
	go makeGrfRequest(wg)
	go makeZkRequest(wg)
	wg.Wait()
}
