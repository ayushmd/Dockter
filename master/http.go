package master

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
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

func corsHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Allow requests from any origin with the specified headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// If it's a preflight request, respond with a 200 status code
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Call the next handler in the chain
		next.ServeHTTP(w, r)
	})
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
	log.Println(r.Method, r.Host, r.URL.Host)
	if HasSubdomain(r.Host) {
		DynamicRouter(w, r, r.Host)
	} else if HasSubdomain(r.Referer()) {
		DynamicRouter(w, r, r.Referer())
	} else {
		router.Run(w, r)
	}
}

func ProxyRequest(w http.ResponseWriter, r *http.Request, targetURL *url.URL) {
	// Create a new HTTP request with the same method, URL, and body as the original request
	// targetURL := r.URL
	proxyReq, err := http.NewRequest(r.Method, targetURL.String(), r.Body)
	if err != nil {
		http.Error(w, "Error creating proxy request", http.StatusInternalServerError)
		return
	}

	// Copy the headers from the original request to the proxy request
	for name, values := range r.Header {
		for _, value := range values {
			proxyReq.Header.Add(name, value)
		}
	}

	// Send the proxy request using the custom transport
	resp, err := http.DefaultTransport.RoundTrip(proxyReq)
	if err != nil {
		http.Error(w, "Error sending proxy request", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Copy the headers from the proxy response to the original response
	for name, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(name, value)
		}
	}

	// Set the status code of the original response to the status code of the proxy response
	w.WriteHeader(resp.StatusCode)

	// Copy the body of the proxy response to the original response
	io.Copy(w, resp.Body)
}

func ServeStaticFiles(w http.ResponseWriter, r *http.Request, s3pth string) {
	svc := NewS3()
	input := &s3.HeadObjectInput{
		Bucket: aws.String(os.Getenv("BUCKET_NAME")),
		Key:    aws.String(s3pth),
	}
	metadata, err := svc.HeadObject(input)
	if err != nil {
		// Object doesn't exist or error occurred
		http.NotFound(w, r)
		return
	}
	resp, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(os.Getenv("BUCKET_NAME")),
		Key:    aws.String(s3pth),
	})
	if err != nil {
		// Error occurred while retrieving object
		http.Error(w, "Failed to retrieve HTML page", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Set the appropriate content type for HTML
	w.Header().Set("Content-Type", *metadata.ContentType)

	// Copy the object's content to the response writer
	_, err = io.Copy(w, resp.Body)
	if err != nil {
		// Error occurred while copying object content
		http.Error(w, "Failed to serve HTML page", http.StatusInternalServerError)
		return
	}
}

func DynamicRouter(w http.ResponseWriter, r *http.Request, source string) {
	subdo := strings.Split(source, ".")[0]
	log.Println("Request to subdomain ", subdo)
	task, err := Master_.GetRecord(subdo)
	fmt.Println(task)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Page not found")
	}
	if task.Type == TYPE_WEBSERVICE {
		Hostip := strings.Split(task.URL.Host, ":")[0]
		deps := fmt.Sprintf("http://%s:%s", Hostip, task.Hostport)
		log.Printf("Redirecting %s to %s", r.Host, deps)
		u, _ := url.ParseRequestURI(deps)
		ProxyRequest(w, r, u)
		// http.Redirect(w, r, deps, http.StatusMovedPermanently)
	} else if task.Type == TYPE_STATIC {
		var s3pth string = ""
		if r.URL.Path == "" {
			s3pth = task.ContainerID + "index.html"
		} else {
			if r.Referer() != "" {
				s3pth = task.ContainerID + r.URL.Path
			}
		}
		ServeStaticFiles(w, r, s3pth)
	}
}

func NewMasterHttpInstance(port int) *http.Server {
	LoadApi(router)
	LoadUI(router)
	return &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: corsHandler(http.HandlerFunc(Handler)),
	}
}
