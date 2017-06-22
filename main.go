package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type TaskList []struct {
	TaskDate string                   `json:"date"`
	TaskDesc []map[string]interface{} `json:"tasks"`
}

var taskmap = make(map[string]interface{})

func PostTasks(w http.ResponseWriter, r *http.Request) {

	var msg TaskList
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		panic(err)
	}
	for _, dates := range msg {
		taskmap[dates.TaskDate] = dates.TaskDesc
	}
	j, _ := json.Marshal(msg)
	log.Println("Creating tasks:", taskmap)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(j)
}

func GetAllTasks(w http.ResponseWriter, r *http.Request) {

	j, err := json.Marshal(taskmap)
	if err != nil {
		panic(err)
	}
	log.Println("Getting all tasks:", taskmap)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func GetDaysTasks(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	date := vars["date"]
	log.Println("Getting a day's task:", taskmap[date])
	j, err := json.Marshal(taskmap[date])
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func DeleteTasks(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	date := vars["date"]
	log.Println("Deleting task:", date, taskmap[date])
	delete(taskmap, date)

	j, err := json.Marshal(taskmap)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func main() {
	t := mux.NewRouter().StrictSlash(false)
	t.HandleFunc("/api/task", PostTasks).Methods("POST")
	t.HandleFunc("/api/task", GetAllTasks).Methods("GET")
	t.HandleFunc("/api/task/{date}", GetDaysTasks).Methods("GET")
	t.HandleFunc("/api/task/{date}", DeleteTasks).Methods("DELETE")
	server := &http.Server{Addr: ":8080", Handler: t}
	server.ListenAndServe()
}
