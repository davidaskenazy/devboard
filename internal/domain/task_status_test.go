package domain

import (
	"encoding/json"
	"testing"
)

type CreateTaskRequest struct {
	Title  string     `json:"title"`
	Status TaskStatus `json:"status"`
}

func TestCreateTaskRequestWithValidStatus(t *testing.T) {
	body := []byte(`{"title":"mi tarea", "status":"todo"}`)

	var req CreateTaskRequest

	err := json.Unmarshal(body, &req)
	if err != nil {
		t.Fatalf("no esperaba error, pero recibi: %v", err)
	}

	if req.Status != TaskStatusTodo {
		t.Fatalf("eseperaba status todo, recibi: %s", req.Status)
	}
}

func TestCreateTaskRequestWithInvalidStatus(t *testing.T) {
	body := []byte(`{"title":"mi tarea", "status":"not ok"}`)

	var req CreateTaskRequest

	err := json.Unmarshal(body, &req)
	if err == nil {
		t.Fatalf("esperaba error, pero recibi nil")
	}

	expected := `estado de tarea invalido: "not ok"`

	if err.Error() != expected {
		t.Fatalf("esperaba error %q, recibi %q", expected, err.Error())
	}
}
