package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

type NOTE struct {
	Num      int    `json:"num"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Category string `json:"category"`
}

var (
	Notee   = []NOTE{}
	nextNum = 1
	mu      sync.Mutex
)

func getnote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	mu.Lock()
	defer mu.Unlock()
	json.NewEncoder(w).Encode(Notee)
}

func creattenote(w http.ResponseWriter, r *http.Request) {
	var note NOTE
	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	mu.Lock()
	note.Num = nextNum
	nextNum++
	Notee = append(Notee, note)
	mu.Unlock()

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(note)

}

func updattenote(w http.ResponseWriter, r *http.Request) {
	var num int
	if err := json.NewDecoder(r.Body).Decode(&num); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	numStr := r.URL.Path[len("/note/update/")]

	num, err := strconv.Atoi(numStr)
	if err != nil {
		http.Error
	}

}

func main() {
	http.HandleFunc("/note", getnote)
	http.HandleFunc("/note/create", creattenote)

	fmt.Println("server running")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("error loading server:", err)
	}
}
