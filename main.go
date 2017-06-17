package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

type PostTask struct {
	TaskDate int64     `json:"date"`
	TaskDesc []string  `json:"desc"`
	Created  time.Time `json:"created"`
}

var taskmap = make(map[int64][]string)

func PostTaskHandler(w http.ResponseWriter, r *http.Request) {

	var msg PostTask
	err := json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		panic(err)
	}
	msg.Created = time.Now()
	for _, v := range msg.TaskDesc {
		taskmap[msg.TaskDate] = append(taskmap[msg.TaskDate], v)
	}
	j, err := json.Marshal(msg)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(j)
}

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	j, err := json.Marshal(taskmap)
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
	t.HandleFunc("/api/task", GetTasksHandler).Methods("GET")
	server := &http.Server{Addr: ":8080", Handler: t}
	server.ListenAndServe()
}
