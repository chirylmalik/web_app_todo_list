package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"todo-list-app/database"
)

type Todo struct {
	ID        int    `json:"id"`
	Task      string `json:"task"`
	Completed bool   `json:"completed"`
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	rows, err := database.GetDB().Query("SELECT id, task, completed FROM tasks")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var tasks []Todo
	for rows.Next() {
		var t Todo
		err := rows.Scan(&t.ID, &t.Task, &t.Completed)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tasks = append(tasks, t)
	}

	json.NewEncoder(w).Encode(tasks)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var t Todo
	json.NewDecoder(r.Body).Decode(&t)

	result, err := database.GetDB().Exec("INSERT INTO tasks (task, completed) VALUES (?, ?)", t.Task, t.Completed)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, _ := result.LastInsertId()
	t.ID = int(id)

	json.NewEncoder(w).Encode(t)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var t Todo
	json.NewDecoder(r.Body).Decode(&t)

	_, err := database.GetDB().Exec("UPDATE tasks SET task = ?, completed = ? WHERE id = ?", t.Task, t.Completed, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	t.ID = id
	json.NewEncoder(w).Encode(t)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	_, err := database.GetDB().Exec("DELETE FROM tasks WHERE id = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
