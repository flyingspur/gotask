package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type TaskList []struct {
	TaskDate string `json:"date"`
	//TaskDesc []string `json:"tasks"`
	TaskDesc []map[string]string `json:"tasks"`
}

var taskmap = make(map[string]interface{})
var item = make(map[string]string)

func PostTasks(w http.ResponseWriter, r *http.Request) {

	var msg TaskList
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		panic(err)
	}
	for _, dates := range msg {
		for _, tasks := range dates.TaskDesc {
			for ktaskitem, taskitem := range tasks {
				item[ktaskitem] = taskitem
			}
		}
		taskmap[dates.TaskDate] = item
	}

	j, _ := json.Marshal(msg)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(j)
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	date := vars["date"]
	task := vars["task"]

	log.Println(date, task)
	log.Println(taskmap[date])
	//if date != nil {
	//	if task != nil {
	//		parsestring = taskmap
	j, err := json.Marshal(taskmap[date])
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func DeleteTasks(w http.ResponseWriter, r *http.Request) {
}

func main() {
	t := mux.NewRouter().StrictSlash(false)
	t.HandleFunc("/api/task", PostTasks).Methods("POST")
	t.HandleFunc("/api/task", GetTasks).Methods("GET")
	t.HandleFunc("/api/task/{date}/{task}", GetTasks).Methods("GET")
	t.HandleFunc("/api/task/{date}", GetTasks).Methods("GET")
	t.HandleFunc("/api/task", DeleteTasks).Methods("DELETE")
	server := &http.Server{Addr: ":8080", Handler: t}
	server.ListenAndServe()
}
