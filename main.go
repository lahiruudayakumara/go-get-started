package main

import (
	"log"
	"net/http"

	"github.com/lahiruudayakumara/go-todo-api/handlers"
)

func main() {

	http.HandleFunc("/api/todos", handlers.GetTodos)
	http.HandleFunc("/api/todo", handlers.GetTodoByID)
	http.HandleFunc("/api/todo/create", handlers.CreateTodo)
	http.HandleFunc("/api/todo/update", handlers.UpdateTodo)
	http.HandleFunc("/api/todo/delete", handlers.DeleteTodo)

	log.Println("Server starting on port :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
