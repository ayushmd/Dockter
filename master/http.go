package master

import (
	"fmt"
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

func HasSubdomain(host string) bool {
	splits := strings.Split(host, ".")
	return len(splits) == 3
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func Handler(w http.ResponseWriter, r *http.Request) {
	// keys := strings.Split(r.URL.Path, "/")
	// task, ok := Master_.cacheDns.Get(keys[0])
	// actualpth := strings.Join(keys[1:], "/")
	// if ok {
	// 	http.Redirect(w, r, "http://"+task.URL.Host+actualpth, http.StatusMovedPermanently)
	// } else {
	// 	router.Run(w, r)
	// }
	enableCors(&w)
	log.Println(r.Method, r.Host, r.URL.Host)
	if HasSubdomain(r.Host) {
		DynamicRouter(w, r)
	} else {
		router.Run(w, r)
	}
}

func DynamicRouter(w http.ResponseWriter, r *http.Request) {
	subdo := strings.Split(r.Host, ".")[0]
	log.Println("Request to subdomain ", subdo)
	task, ok := Master_.cacheDns.Get(subdo)
	var Hostip string = ""
	var Hostport string = ""
	if ok {
		Hostip = strings.Split(task.URL.Host, ":")[0]
		Hostport = task.Hostport

	} else if Master_.dbDns != nil {
		row := Master_.GetDnsRecord(subdo)
		var Subdomain string
		var HostIp string
		var HostPort string
		var RunningPort string
		var ImageName string
		var ContainerID string
		err := row.Scan(&Subdomain, &HostIp, &HostPort, &RunningPort, &ImageName, &ContainerID)
		if err == nil {
			Hostip = HostIp
			Hostport = HostPort
		}
	}
	if Hostip != "" && Hostport != "" {
		deps := fmt.Sprintf("http://%s:%s", Hostip, Hostport)
		log.Printf("Redirecting %s to %s", r.Host, deps)
		http.Redirect(w, r, deps, http.StatusMovedPermanently)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Page not found")
	}
}

func NewMasterHttpInstance(port int) *http.Server {
	LoadApi(router)
	LoadUI(router)
	return &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: http.HandlerFunc(Handler),
	}
}
