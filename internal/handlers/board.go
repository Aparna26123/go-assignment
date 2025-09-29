package handlers

import (
	"encoding/json" //for encoding and decoding json data
	"net/http"
	"strconv"                          //converts string to other types
	"task-manager/internal/models"     //models contain board struct def
	"task-manager/internal/repository" //repo handles database operations

	"github.com/gorilla/mux"
)

func CreateBoard(w http.ResponseWriter, r *http.Request) {
	var board models.Board
	json.NewDecoder(r.Body).Decode(&board)

	repository.DB.Create(&board)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(board)
}

func GetBoards(w http.ResponseWriter, r *http.Request) {
	var boards []models.Board         //to hold mul boards
	repository.DB.Find(&boards)       //fetches all board from database
	json.NewEncoder(w).Encode(boards) //sends list of boards
}

func DeleteBoard(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	repository.DB.Delete(&models.Board{}, id) //del board with given ID
	w.WriteHeader(http.StatusNoContent)       //204 no content
}
