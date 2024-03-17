package master

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type Router struct {
	GetMapper    map[string]func(http.ResponseWriter, *http.Request)
	PostMapper   map[string]func(http.ResponseWriter, *http.Request)
	PutMapper    map[string]func(http.ResponseWriter, *http.Request)
	DeleteMapper map[string]func(http.ResponseWriter, *http.Request)
}

func (r *Router) Get(
	path string,
	callback func(http.ResponseWriter, *http.Request),
) {
	if r.GetMapper == nil {
		r.GetMapper = make(map[string]func(http.ResponseWriter, *http.Request))
	}
	r.GetMapper[path] = callback
}

func (r *Router) Post(
	path string,
	callback func(http.ResponseWriter, *http.Request),
) {
	if r.PostMapper == nil {
		r.PostMapper = make(map[string]func(http.ResponseWriter, *http.Request))
	}
	r.PostMapper[path] = callback
}

func (r *Router) Put(
	path string,
	callback func(http.ResponseWriter, *http.Request),
) {
	if r.PutMapper == nil {
		r.PutMapper = make(map[string]func(http.ResponseWriter, *http.Request))
	}
	r.PutMapper[path] = callback
}

func (r *Router) Delete(
	path string,
	callback func(http.ResponseWriter, *http.Request),
) {
	if r.DeleteMapper == nil {
		r.DeleteMapper = make(map[string]func(http.ResponseWriter, *http.Request))
	}
	r.DeleteMapper[path] = callback
}

func (r *Router) Run(
	w http.ResponseWriter, req *http.Request,
) {
	path := req.URL.Path
	var callback func(http.ResponseWriter, *http.Request)
	var ok bool
	switch req.Method {
	case http.MethodGet:
		callback, ok = r.GetMapper[path]
	case http.MethodPost:
		callback, ok = r.PostMapper[path]
	case http.MethodPut:
		callback, ok = r.PutMapper[path]
	case http.MethodDelete:
		callback, ok = r.DeleteMapper[path]
	}
	if ok {
		callback(w, req)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Page not found")
	}
}

var router *Router = &Router{}

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Has ", r.Host)
	hasSubdomain := strings.Contains(r.Host, ".")
	fmt.Println("Has ", hasSubdomain)
	if hasSubdomain {
		DynamicRouter(w, r)
	} else {
		fmt.Println(router.GetMapper)
		router.Run(w, r)
	}
}

func DynamicRouter(w http.ResponseWriter, r *http.Request) {
	subdo := strings.Split(r.Host, ".")[0]
	task, ok := Master_.cacheDns.Get(subdo)
	if ok {
		http.Redirect(w, r, "http://"+task.URL.Host+r.URL.Path, http.StatusMovedPermanently)
	}
}

func NewMasterHttpInstance(port int) *http.Server {

	router.Get("/hi", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "My router response")
	})

	router.Get("/GetServerData", func(w http.ResponseWriter, r *http.Request) {
		servPoolData, err := Master_.GetServerPoolHandler()
		fmt.Println(string(servPoolData))
		if err != nil {
			w.Write([]byte("Error occured"))
		}
		var resp map[string]string = map[string]string{
			"backend": string(servPoolData),
		}
		w.Header().Set("Content-Type", "application/json")
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		w.Write(jsonResp)
	})

	router.Post("/buildraw", func(w http.ResponseWriter, r *http.Request) {
		var task TaskRawRequest
		defer r.Body.Close()
		body, _ := io.ReadAll(r.Body)
		if err := json.Unmarshal(body, &task); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error")
		}
		enc, _ := json.Marshal(task)
		Master_.AddTask("BUILDRAW", enc)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Added to queue")
	})

	router.Post("/buildfile", func(w http.ResponseWriter, r *http.Request) {
		var task TaskFileRequest
		defer r.Body.Close()
		body, _ := io.ReadAll(r.Body)
		if err := json.Unmarshal(body, &task); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error")
		}
		enc, _ := json.Marshal(task)
		Master_.AddTask("BUILDFILE", enc)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Added to queue")
	})

	router.Post("/deploy", func(w http.ResponseWriter, r *http.Request) {
		var task TaskImageRequest
		defer r.Body.Close()
		body, _ := io.ReadAll(r.Body)
		if err := json.Unmarshal(body, &task); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error")
		}
		enc, _ := json.Marshal(task)
		Master_.AddTask("DEPLOY", enc)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Added to queue")
	})

	return &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: http.HandlerFunc(Handler),
	}
}
