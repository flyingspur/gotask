package main

import ("fmt"
		"net/http"
		"github.com/gorilla/mux"
		"time"
		"encoding/json"
)

type Task struct {
	TaskDate 	string 	`json:"date"`
	TaskDesc	string		`json:"desc"`
	Created		time.Time	`json:"created"`
}

var task = make(map[string]string)

func PostTask(w http.ResponseWriter, r *http.Request) {

}