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
	for i := 0; i < 10; i++ {
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
	for i := 0; i < 10; i++ {
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
	for i := 0; i < 10; i++ {
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

// ngx
// "memUsage": 2654208,
// "cpuPercent": 0,
// "diskUsage": 187615885
// zk
// "memUsage": 54628352,
// "cpuPercent": 0,
// "diskUsage": 313803452
// grf
// "memUsage": 98250752,
// "cpuPercent": 0,
// "diskUsage": 430257759

// 15.207.20.43:3210-nginx - ngx0:9366873429 (DEPLOY received at 18:37:04)
// 15.207.20.43:3210-nginx - ngx1:2324979144 (DEPLOY received at 18:37:48)
// 15.207.20.43:3210-nginx - ngx4:2476343593 (DEPLOY received at 18:38:57)
// 15.207.20.43:3210-zookeeper - zk1:11921829496 (DEPLOY received at 18:38:01)
// 15.207.20.43:3210-grafana/grafana - grf4:14960913619 (DEPLOY received at 18:39:15)
// 15.207.20.43:3210-grafana/grafana - grf9:2806585333 (DEPLOY received at 18:40:13)
// 15.207.20.43:3210-zookeeper - zk4:2377172955 (DEPLOY received at 18:39:00)
// 15.207.20.43:3210-zookeeper - zk9:2797576690 (DEPLOY received at 18:40:16)

// 13.232.212.234:3210-zookeeper - zk0:11962570013 (DEPLOY received at 18:37:16)
// 13.232.212.234:3210-nginx - ngx2:9940003941 (DEPLOY received at 18:38:11)
// 13.232.212.234:3210-zookeeper - zk3:2363826233 (DEPLOY received at 18:38:39)
// 13.232.212.234:3210-grafana/grafana - grf3:14705476424 (DEPLOY received at 18:38:54)
// 13.232.212.234:3210-nginx - ngx6:2517118231 (DEPLOY received at 18:39:43)
// 13.232.212.234:3210-zookeeper - zk6:2472803867 (DEPLOY received at 18:39:46)
// 13.232.212.234:3210-grafana/grafana - grf6:2650197849 (DEPLOY received at 18:39:49)
// 13.232.212.234:3210-nginx - ngx7:2907620101 (DEPLOY received at 18:39:52)

// 13.235.70.21:3210-grafana/grafana - grf0:14658712132 (DEPLOY received at 18:37:31)
// 13.235.70.21:3210-grafana/grafana - grf2:2727119817 (DEPLOY received at 18:38:26)
// 13.235.70.21:3210-nginx - ngx3:9517234403 (DEPLOY received at 18:38:36)
// 13.235.70.21:3210-grafana/grafana - grf8:3143826076 (DEPLOY received at 18:40:04)
// 13.235.70.21:3210-nginx - ngx8:2666467860 (DEPLOY received at 18:40:07)
// 13.235.70.21:3210-zookeeper - zk5:12358018123 (DEPLOY received at 18:39:37)
// 13.235.70.21:3210-nginx - ngx9:2494893303 (DEPLOY received at 18:40:10)

// 15.207.16.145:3210-grafana/grafana - grf1:13909264197 (DEPLOY received at 18:37:46)
// 15.207.16.145:3210-zookeeper - zk2:12180324693 (DEPLOY received at 18:38:23)
// 15.207.16.145:3210-nginx - ngx5:9091952740 (DEPLOY received at 18:39:24)
// 15.207.16.145:3210-grafana/grafana - grf5:2414755092 (DEPLOY received at 18:39:40)
// 15.207.16.145:3210-grafana/grafana - grf7:2344583866 (DEPLOY received at 18:39:54)
// 15.207.16.145:3210-zookeeper - zk7:2494620259 (DEPLOY received at 18:39:57)
// 15.207.16.145:3210-zookeeper - zk8:2914829219 (DEPLOY received at 18:40:01)
