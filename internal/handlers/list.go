package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"task-manager/internal/models"
	"task-manager/internal/repository"

	"github.com/gorilla/mux"
)

func CreateList(w http.ResponseWriter, r *http.Request) {
	boardID, _ := strconv.Atoi(mux.Vars(r)["id"])
	var list models.List
	json.NewDecoder(r.Body).Decode(&list)
	list.BoardID = uint(boardID)

	repository.DB.Create(&list)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(list)
}

func GetLists(w http.ResponseWriter, r *http.Request) {
	boardID, _ := strconv.Atoi(mux.Vars(r)["id"])
	var lists []models.List
	repository.DB.Where("board_id = ?", boardID).Find(&lists)
	json.NewEncoder(w).Encode(lists)
}

func UpdateList(w http.ResponseWriter, r *http.Request) {
	listID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid list ID", http.StatusBadRequest)
		return
	}

	var updateData models.List
	if err := json.NewDecoder(r.Body).Decode(&updateData); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	var list models.List
	if err := repository.DB.First(&list, listID).Error; err != nil {
		http.Error(w, "List not found", http.StatusNotFound)
		return
	}

	list.Name = updateData.Name
	list.Position = updateData.Position
	if err := repository.DB.Save(&list).Error; err != nil {
		http.Error(w, "Failed to update list", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(list)
}

func DeleteList(w http.ResponseWriter, r *http.Request) {
	listID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid list ID", http.StatusBadRequest)
		return
	}

	if err := repository.DB.Delete(&models.List{}, listID).Error; err != nil {
		http.Error(w, "Failed to delete list", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
