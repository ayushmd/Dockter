package master

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type NameRequest struct {
	Name string `json:"name"`
}

type KeyPairReq struct {
	KeyName   string `json:"keyName"`
	Algorithm string `json:"algorithm"`
}
type KeyPairResp struct {
	KeyName    string `json:"keyName"`
	PrivateKey string `json:"privateKey"`
	PublicKey  string `json:"publicKey"`
}

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

	router.Post("/api/getStatus", func(w http.ResponseWriter, r *http.Request) {
		var req NameRequest
		defer r.Body.Close()
		body, _ := io.ReadAll(r.Body)
		if err := json.Unmarshal(body, &req); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error")
		}
		row := Master_.GetDnsStatus(req.Name)
		var Status string
		err := row.Scan(&Status)
		fmt.Println(req.Name, " With Status ", Status)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, Status)
		}
	})

	router.Delete("/api/obliterate", func(w http.ResponseWriter, r *http.Request) {
		var req NameRequest
		defer r.Body.Close()
		body, _ := io.ReadAll(r.Body)
		if err := json.Unmarshal(body, &req); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error")
		}
		err := Master_.TerminateTask(req.Name)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Could not remove the task")
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Obliterated")
	})

	router.Post("/api/metrics", func(w http.ResponseWriter, r *http.Request) {
		var req NameRequest
		defer r.Body.Close()
		body, _ := io.ReadAll(r.Body)
		if err := json.Unmarshal(body, &req); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error")
		}
		metrics, err := Master_.TaskMetrics(req.Name)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Could not remove the task")
		}
		w.Header().Set("Content-Type", "application/json")
		jsonResp, _ := json.Marshal(metrics)
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

	router.Post("/api/static", func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseMultipartForm(10 << 20)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		domain := r.MultipartForm.Value["name"][0]
		files := r.MultipartForm.File["files"]
		err = Master_.DeployStatic(domain, files)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, "Files Uploaded Successfully!")
	})

	router.Post("/api/keypair", func(w http.ResponseWriter, r *http.Request) {
		var req KeyPairReq
		defer r.Body.Close()
		body, _ := io.ReadAll(r.Body)
		if err := json.Unmarshal(body, &req); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error")
		}
		pk := Master_.GetSSHKeys(req.Algorithm, req.KeyName)
		w.Header().Set("Content-Type", "application/json")
		jsonResp, _ := json.Marshal(KeyPairResp{
			PrivateKey: pk,
			KeyName:    req.KeyName,
		})
		w.Write(jsonResp)
	})
}
