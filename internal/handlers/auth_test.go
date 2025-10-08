package handlers

import (
	"net/http"          //provides http client and server implementations
	"net/http/httptest" //for testing http servers & create mock requests
	"task-manager/internal/repository"
	"testing"
)

func TestRegister(t *testing.T) {

	repository.InitDB()
	req := httptest.NewRequest("POST", "/register", nil) //creates new http post request
	w := httptest.NewRecorder()                          //creates response recorder
	Register(w, req)                                     //calling reg handler fun with mock req & response writer
	if w.Code != http.StatusCreated {                    //checks response ststus @201
		t.Errorf("Expected 201, got %d", w.Code)
	}
}
