package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

type Task struct {
	TaskDate int       `json:"date"`
	TaskDesc string    `json:"desc"`
	Created  time.Time `json:"created"`
}

var taskhash = make(map[int]string)

func PostTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		panic(err)
	}
	taskhash[task.TaskDate] = task.TaskDesc
	task.Created = time.Now()
	j, err := json.Marshal(task)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(j)
}

func main() {
	t := mux.NewRouter().StrictSlash(false)
	t.HandleFunc("/api/task", PostTaskHandler).Methods("POST")
	server := &http.Server{Addr:":8080", Handler: t,}
	server.ListenAndServe()
}
