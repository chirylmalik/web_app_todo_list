package main

import (
	"log"
	"net/http"

	"todo-list-app/database"
	"todo-list-app/handlers"

	"github.com/gorilla/mux"
)

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	database.ConnectDB()

	r := mux.NewRouter()
	r.HandleFunc("/tasks", handlers.GetTasks).Methods("GET")
	r.HandleFunc("/tasks", handlers.CreateTask).Methods("POST")
	r.HandleFunc("/tasks/{id:[0-9]+}", handlers.UpdateTask).Methods("PUT")
	r.HandleFunc("/tasks/{id:[0-9]+}", handlers.DeleteTask).Methods("DELETE")

	log.Println("Server berjalan di port 8080")
	log.Fatal(http.ListenAndServe(":8080", enableCORS(r)))
}
