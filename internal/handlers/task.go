package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

// Mock database
var tasks = []Task{
	{ID: 1, Title: "Fix Launcher"},
}

func TaskHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// 1. Determine if we are looking for a specific ID
	// Path: /tasks/ or /tasks/1
	idString := strings.TrimPrefix(r.URL.Path, "/tasks/")

	switch r.Method {
	case http.MethodGet:
		if idString == "" {
			getAllTasks(w, r)
		} else {
			getTaskByID(w, r, idString)
		}

	case http.MethodPost:
		createTask(w, r)

	case http.MethodPatch:
		updateTask(w, r, idString)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// --- Helper Functions ---

func getAllTasks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(tasks)
}

func getTaskByID(w http.ResponseWriter, r *http.Request, idStr string) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID format (must be integer)", http.StatusBadRequest)
		return
	}

	for _, t := range tasks {
		if t.ID == id {
			json.NewEncoder(w).Encode(t)
			return
		}
	}

	http.Error(w, "Task not found", http.StatusNotFound)
}

func createTask(w http.ResponseWriter, r *http.Request) {
	var newTask Task
	if err := json.NewDecoder(r.Body).Decode(&newTask); err != nil {
		http.Error(w, "Bad Request Body", http.StatusBadRequest)
		return
	}
	tasks = append(tasks, newTask)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)
}

func updateTask(w http.ResponseWriter, r *http.Request, idStr string) {
	// Logic similar to getTaskByID, but modifying the slice
}
