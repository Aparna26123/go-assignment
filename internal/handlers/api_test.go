package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"task-manager/internal/repository"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

func generateTestToken(userID uint) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour).Unix(),
	})
	tokenString, _ := token.SignedString(jwtKey) // uses jwtKey from auth.go
	return tokenString
}

func TestMe(t *testing.T) {
	repository.InitDB()
	token := generateTestToken(1)
	req := httptest.NewRequest("GET", "/api/me", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()
	Me(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Expected 200 OK, got %d", w.Code)
	}
}

func TestCreateBoard(t *testing.T) {
	repository.InitDB()
	token := generateTestToken(1)
	body := `{"name":"Test Board","user_id":1}`
	req := httptest.NewRequest("POST", "/api/boards", strings.NewReader(body))
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	CreateBoard(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("Expected 201 Created, got %d", w.Code)
	}
}

func TestGetBoards(t *testing.T) {
	repository.InitDB()
	token := generateTestToken(1)
	req := httptest.NewRequest("GET", "/api/boards", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()
	GetBoards(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Expected 200 OK, got %d", w.Code)
	}
}

func TestDeleteBoard(t *testing.T) {
	repository.InitDB()
	token := generateTestToken(1)
	req := httptest.NewRequest("DELETE", "/api/boards/1", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	req.Header.Set("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()
	DeleteBoard(w, req)
	if w.Code != http.StatusNoContent {
		t.Errorf("Expected 204 No Content, got %d", w.Code)
	}
}

func TestCreateList(t *testing.T) {
	repository.InitDB()
	token := generateTestToken(1)
	body := `{"name":"Test List","position":1}`
	req := httptest.NewRequest("POST", "/api/boards/1/lists", strings.NewReader(body))
	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	CreateList(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("Expected 201 Created, got %d", w.Code)
	}
}

func TestGetLists(t *testing.T) {
	repository.InitDB()
	token := generateTestToken(1)
	req := httptest.NewRequest("GET", "/api/boards/1/lists", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	req.Header.Set("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()
	GetLists(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Expected 200 OK, got %d", w.Code)
	}
}

func TestUpdateList(t *testing.T) {
	repository.InitDB()
	token := generateTestToken(1)
	body := `{"name":"Updated List","position":2}`
	req := httptest.NewRequest("PUT", "/api/lists/1", strings.NewReader(body))
	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	UpdateList(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Expected 200 OK, got %d", w.Code)
	}
}

func TestDeleteList(t *testing.T) {
	repository.InitDB()
	token := generateTestToken(1)
	req := httptest.NewRequest("DELETE", "/api/lists/1", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	req.Header.Set("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()
	DeleteList(w, req)
	if w.Code != http.StatusNoContent {
		t.Errorf("Expected 204 No Content, got %d", w.Code)
	}
}

func TestCreateTask(t *testing.T) {
	repository.InitDB()
	token := generateTestToken(1)
	body := `{"title":"Test Task"}`
	req := httptest.NewRequest("POST", "/api/lists/1/tasks", strings.NewReader(body))
	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	CreateTask(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("Expected 201 Created, got %d", w.Code)
	}
}

func TestGetTasks(t *testing.T) {
	repository.InitDB()
	token := generateTestToken(1)
	req := httptest.NewRequest("GET", "/api/lists/1/tasks", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	req.Header.Set("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()
	GetTasks(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Expected 200 OK, got %d", w.Code)
	}
}

func TestUpdateTask(t *testing.T) {
	repository.InitDB()
	token := generateTestToken(1)
	body := `{"title":"Updated Task","done":true}`
	req := httptest.NewRequest("PUT", "/api/tasks/1", strings.NewReader(body))
	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	UpdateTask(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Expected 200 OK, got %d", w.Code)
	}
}

func TestDeleteTask(t *testing.T) {
	repository.InitDB()
	token := generateTestToken(1)
	req := httptest.NewRequest("DELETE", "/api/tasks/1", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	req.Header.Set("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()
	DeleteTask(w, req)
	if w.Code != http.StatusNoContent {
		t.Errorf("Expected 204 No Content, got %d", w.Code)
	}
}

func TestMarkTaskComplete(t *testing.T) {
	repository.InitDB()
	token := generateTestToken(1)
	req := httptest.NewRequest("PATCH", "/api/tasks/1/complete", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	req.Header.Set("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()
	MarkTaskComplete(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Expected 200 OK, got %d", w.Code)
	}
}
