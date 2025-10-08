package main

import (
	"log"
	"net/http"
	"task-manager/internal/handlers"
	"task-manager/internal/middleware"
	"task-manager/internal/repository"

	"github.com/gorilla/mux"
)

func main() {
	repository.InitDB()

	r := mux.NewRouter()

	// Public endpoints (no JWT needed)

	r.HandleFunc("/register", handlers.Register).Methods("POST")
	r.HandleFunc("/login", handlers.Login).Methods("POST")

	// Protected endpoints (JWT needed)
	api := r.PathPrefix("/api").Subrouter()
	api.Use(middleware.JWTMiddleware)

	api.HandleFunc("/me", handlers.Me).Methods("GET")

	api.HandleFunc("/boards", handlers.CreateBoard).Methods("POST")
	api.HandleFunc("/boards", handlers.GetBoards).Methods("GET")
	api.HandleFunc("/boards/{id}", handlers.DeleteBoard).Methods("DELETE")

	api.HandleFunc("/boards/{id}/lists", handlers.CreateList).Methods("POST")
	api.HandleFunc("/boards/{id}/lists", handlers.GetLists).Methods("GET")
	api.HandleFunc("/lists/{id}", handlers.DeleteList).Methods("DELETE")
	api.HandleFunc("/lists/{id}", handlers.UpdateList).Methods("PUT")

	api.HandleFunc("/lists/{id}/tasks", handlers.CreateTask).Methods("POST")
	api.HandleFunc("/lists/{id}/tasks", handlers.GetTasks).Methods("GET")
	api.HandleFunc("/tasks/{id}", handlers.UpdateTask).Methods("PUT")
	api.HandleFunc("/tasks/{id}", handlers.DeleteTask).Methods("DELETE")
	api.HandleFunc("/tasks/{id}/complete", handlers.MarkTaskComplete).Methods("PATCH")

	log.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
