package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Todo struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

var todos []Todo

func main() {
	//	define routes
	http.HandleFunc("/todos", getTodos)

	//	start the server and listen
	log.Fatal(http.ListenAndServe(":9000", nil))
}

// define getTodos controller
func getTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(todos)
	if err != nil {
		return
	}
}
