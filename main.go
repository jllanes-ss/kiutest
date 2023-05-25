package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

var names sync.Map

type postHelloReq struct {
	Name string `json:"name"`
}

type postHelloRes struct {
	Message string `json:"message"`
	Exists  bool   `json:"exists"`
}

type getHelloRes struct {
	Names []string `json:"names"`
}

// helloHandler general handler of route hello
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/json" {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	r.Header.Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		getHello(w, r)
	case "POST":
		postHello(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// getHello get handler of route hello
func getHello(w http.ResponseWriter, r *http.Request) {

	var res getHelloRes

	names.Range(func(key, value interface{}) bool {
		res.Names = append(res.Names, fmt.Sprint(key))
		return true
	})
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

// postHello post handler of route hello
func postHello(w http.ResponseWriter, r *http.Request) {

	var req postHelloReq
	var res postHelloRes
	var statusCode int
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "missing payload"}`))
		return
	}

	if _, ok := names.Load(req.Name); ok {
		statusCode = http.StatusOK
		res.Message = fmt.Sprintf("Hello, %s! Welcome back!", req.Name)
		res.Exists = true
	} else {
		statusCode = http.StatusCreated
		names.Store(req.Name, nil)
		res.Message = fmt.Sprintf("Hello, %s!", req.Name)
		res.Exists = false
	}

	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(res)
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
