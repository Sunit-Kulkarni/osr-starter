package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

type Todo struct {
	DueInSeconds int `json:"due_in_seconds"`
	Comment string `json:"comment"`
}

var TheList []Todo

func main() {
	fmt.Println("howdy")

	r := mux.NewRouter()
	r.HandleFunc("/set", SetHandler)
	r.HandleFunc("/list", ListsHandler)
	http.Handle("/", r)

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func ListsHandler(w http.ResponseWriter, r *http.Request) {
	// an example API handler
	json.NewEncoder(w).Encode(map[string][]Todo{"todos": TheList})
}

func SetHandler(w http.ResponseWriter, r *http.Request)  {
	var todo Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	TheList = append(TheList, todo)
	json.NewEncoder(w).Encode(map[string]bool{"status": true})
}