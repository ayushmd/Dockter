package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"sync"
)

func makeRequest(name string) {
	url := "http://ayushd.cloud/api/buildraw"
	var jsonStr = []byte(fmt.Sprintf(`{
"name":"%s",
"gitlink":"https://github.com/johnpapa/node-hello.git",
"branch":"master",
"buildCmd":"npm i",
"startCmd":"npm start",
"runtimeEnv":"Node",
"runningPort":"3000",
}`, name))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
}

func Client() {
	wg := sync.WaitGroup{}
	for i := 0; i < 30; i++ {
		wg.Add(1)
		name := fmt.Sprintf("node%d", i)
		go func() {
			makeRequest(name)
			wg.Done()
		}()
	}
	wg.Wait()
}
