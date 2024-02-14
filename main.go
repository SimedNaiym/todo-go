package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type todo struct {
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

var todos []todo = make([]todo, 0)

func addTodoHandler(rw http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		var body todo

		err := json.NewDecoder(r.Body).Decode(&body)

		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		todos = append(todos, body)

		rw.WriteHeader(http.StatusCreated)
	}
}

func getTodosHandler(rw http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		res, err := json.Marshal(todos)

		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		rw.Header().Add("Content-Type", "text/json")
		rw.Write(res)
	}
}

func main() {

	todos = append(todos, todo{
		Title: "learn go",
		Done:  false,
	})

	// Get all Todos : GET /api/todos
	http.HandleFunc("/api/todos", getTodosHandler)

	// Post new todo: POST /api/add {Body}
	http.HandleFunc("/api/add", addTodoHandler)

	err := http.ListenAndServe("0.0.0.0:8080", nil)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

}