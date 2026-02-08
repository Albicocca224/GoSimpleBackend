package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

var tasks = []Task{
	{ID: 1, Title: "Write unit tests", Done: false},
	{ID: 2, Title: "Deploy service", Done: true},
}

func TaskHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // [cite: 111]

	switch r.Method {
	case http.MethodGet:
		handleGet(w, r)
	case http.MethodPost:
		handlePost(w, r)
	case http.MethodPatch:
		handlePatch(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")

	if idStr != "" {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "invalid id"}) // [cite: 43]
			return
		}

		for _, task := range tasks {
			if task.ID == id {
				json.NewEncoder(w).Encode(task) // [cite: 34-39]
				return
			}
		}

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "task not found"}) // [cite: 47]
		return
	}

	json.NewEncoder(w).Encode(tasks)
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	var reqBody struct {
		Title string `json:"title"`
	}

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if reqBody.Title == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid title"}) // [cite: 75]
		return
	}

	newID := 1
	if len(tasks) > 0 {
		newID = tasks[len(tasks)-1].ID + 1
	}

	newTask := Task{
		ID:    newID,
		Title: reqBody.Title,
		Done:  false,
	}

	tasks = append(tasks, newTask)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)
}

func handlePatch(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		w.WriteHeader(http.StatusNotFound) // [cite: 89] (Implied by requirement to find task)
		json.NewEncoder(w).Encode(map[string]string{"error": "task not found"})
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid id"})
		return
	}

	var reqBody struct {
		Done bool `json:"done"`
	}
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "Invalid Body", http.StatusBadRequest)
		return
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Done = reqBody.Done

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]bool{"updated": true})
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "task not found"})
}
