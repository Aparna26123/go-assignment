package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"task-manager/internal/models"
	"task-manager/internal/repository"

	"github.com/gorilla/mux"
)

// POST /lists/{id}/tasks
func CreateTask(w http.ResponseWriter, r *http.Request) {
	listID, _ := strconv.Atoi(mux.Vars(r)["id"])
	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)
	task.ListID = uint(listID)

	repository.DB.Create(&task)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

// GET /lists/{id}/tasks
func GetTasks(w http.ResponseWriter, r *http.Request) {
	listID, _ := strconv.Atoi(mux.Vars(r)["id"])
	var tasks []models.Task
	repository.DB.Where("list_id = ?", listID).Find(&tasks)
	json.NewEncoder(w).Encode(tasks)
}

// PUT /tasks/{id}
func UpdateTask(w http.ResponseWriter, r *http.Request) {
	taskID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	var updatedData models.Task
	if err := json.NewDecoder(r.Body).Decode(&updatedData); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	var task models.Task
	if err := repository.DB.First(&task, taskID).Error; err != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	task.Title = updatedData.Title
	task.Done = updatedData.Done
	if err := repository.DB.Save(&task).Error; err != nil {
		http.Error(w, "Failed to update task", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(task)
}

// DELETE /tasks/{id}
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	taskID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	if err := repository.DB.Delete(&models.Task{}, taskID).Error; err != nil {
		http.Error(w, "Failed to delete task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// PATCH /tasks/{id}/complete
func MarkTaskComplete(w http.ResponseWriter, r *http.Request) {
	taskID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	var task models.Task
	if err := repository.DB.First(&task, taskID).Error; err != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	task.Done = true

	if err := repository.DB.Save(&task).Error; err != nil {
		http.Error(w, "Failed to mark task as complete", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(task)
}
