package master

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func LoadApi(router *Router) {
	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "My router pongged")
	})

	router.Get("/api/getdata", func(w http.ResponseWriter, r *http.Request) {
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

	router.Post("/api/buildraw", func(w http.ResponseWriter, r *http.Request) {
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

	router.Post("/api/buildfile", func(w http.ResponseWriter, r *http.Request) {
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

	router.Post("/api/deploy", func(w http.ResponseWriter, r *http.Request) {
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
}
