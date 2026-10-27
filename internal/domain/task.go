package domain

import (
	"errors"
	"time"
)

// Task representa una tarea dentro del tablero
type Task struct {
	ID          string
	ProyectID   string
	Title       string
	Description string
	Status      TaskStatus
	AssigneeID  *string
	CreatedBy   string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// NewTask construye una tarea nueva
func NewTask(proyectID, title, createdBy string) *Task {
	now := time.Now()
	return &Task{
		ProyectID: proyectID,
		Title:     title,
		Status:    TaskStatusTodo,
		CreatedBy: createdBy,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

var (
	//ErrTaskAlreadyDone tarea ya hecha
	ErrTaskAlreadyDone = errors.New("la tarea ya esta completada")
	//ErrTaskNotAssignable tarea no asignable en un estado done
	ErrTaskNotAssignable = errors.New("la tarea no puede asignarse en este estado")
)

// AssignTo asigna la tarea a un usuario
func (t *Task) AssignTo(userID string) error {
	if t.Status == TaskStatusDone {
		return ErrTaskNotAssignable
	}
	t.AssigneeID = &userID
	t.UpdatedAt = time.Now()
	return nil
}

// MarkAsDone marca la tarea como completada
func (t *Task) MarkAsDone() error {
	if t.Status == TaskStatusDone {
		return ErrTaskAlreadyDone
	}
	t.Status = TaskStatusDone
	t.UpdatedAt = time.Now()
	return nil
}

// MoveToStatus mueve la tarea a un nuevo estado del tablero
func (t *Task) MoveToStatus(newStatus TaskStatus) error {
	if !newStatus.IsValid() {
		return errors.New("estado destino invalido")
	}
	t.Status = newStatus
	t.UpdatedAt = time.Now()
	return nil
}
