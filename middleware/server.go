package middleware

import (
	"fmt"
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
	// switch r.Method {
	// case http.MethodGet:
	// 	// Handle GET requests to proxy it
	// 	HandleGet(w, r)
	// case http.MethodPost:
	// 	// Handle POST requests
	// 	callback := PostMapper[r.URL.Path]
	// 	callback(w, r)
	// default:
	// 	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	// }
}

func DynamicRouter(w http.ResponseWriter, r *http.Request) {
	// Your GET request handling logic here
	fmt.Println(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "GET request handled")
}

func NewMiddlewareInstance(port int) http.Server {
	router.Get("/hi", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "My router response")
	})
	return http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: http.HandlerFunc(Handler),
	}
}
