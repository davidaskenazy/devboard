package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"time"

	"github.com/davidaskenazy/devboard/internal/domain"
	"github.com/davidaskenazy/devboard/internal/usecase"
	"github.com/davidaskenazy/devboard/internal/validator"
)

// TaskHandler struct
type TaskHandler struct {
	usercase  *usecase.TaskUseCase
	validator *validator.Validator
	logger    *slog.Logger
}

// NewTaskHandler constructor
func NewTaskHandler(uc *usecase.TaskUseCase, v *validator.Validator, logger *slog.Logger) *TaskHandler {
	return &TaskHandler{
		usercase:  uc,
		validator: v,
		logger:    logger,
	}
}

// createTaskRequest estructura
type createTaskRequest struct {
	ProyectID string `json:"proyect_id" validate:"required"`
	Title     string `json:"title" validate:"required,min=3,max=200"`
}

// createTaskResponse estructura
type createTaskResponse struct {
	ProjectID string `json:"proyect_id"`
	Title     string `json:"title"`
}

// taskResponse estructura
type taskResponse struct {
	ID          string            `json:"id"`
	ProyectID   string            `json:"proyect_id"`
	Title       string            `json:"title"`
	Description string            `json:"desctiption,omitempty"`
	Status      domain.TaskStatus `json:"status"`
	AssigneeID  *string           `json:"assignee_id,omitempty"`
	CreatedBy   string            `json:"created_by"`
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
}

// toTaskResponse convierte entidad de dominio a DTO de respuesta
func toTaskResponse(t *domain.Task) taskResponse {
	return taskResponse{
		ID:          t.ID,
		ProyectID:   t.ProyectID,
		Title:       t.Title,
		Description: t.Description,
		Status:      t.Status,
		AssigneeID:  t.AssigneeID,
		CreatedBy:   t.CreatedBy,
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   t.UpdatedAt,
	}
}

// Create task endpoint
//
// @Summary     Crear tarea
// @Description Registra una nueva tarea en el sistema
// @Tags        tasks
// @Accept json
// @Produce     json
// @Param request body createTaskRequest true "Datos de la tarea"
// @Success     201 {object} createTaskResponse
// @Failure 400 {object} ProblemDetails
// @Router      /tasks [post]
func (h *TaskHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req createTaskRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		RespondError(w, r, h.logger, domain.ErrInvalidInput)
		return
	}

	if errs := h.validator.Validate(req); errs != nil {
		if err := RespondValidationError(w, errs); err != nil {
			h.logger.ErrorContext(r.Context(),
				"error al recibir respuesta de validacion",
				"error", err,
				"method", r.Method,
				"path", r.URL.Path,
			)
		}
		return
	}

	task, err := h.usercase.CreateTask(r.Context(), req.ProyectID, req.Title, "system")
	if err != nil {
		RespondError(w, r, h.logger, err)
		return
	}
	if err := RespondJSON(w, http.StatusCreated, createTaskResponse{
		ProjectID: task.ProyectID,
		Title:     task.Title,
	}); err != nil {
		h.logger.ErrorContext(r.Context(),
			"error al escribir respuesta JSON de tarea",
			"error", err,
			"method", r.Method,
			"path", r.URL.Path,
		)
	}
}

// Get metodo para obtener tarea por ID
//
// @Summary     Obtener tarea por ID
// @Description Obtiene una tarea por su ID
// @Tags        tasks
// @Produce     json
// @Param id path string true "ID de la tarea"
// @Success     200 {object} taskResponse
// @Failure 404 {object} ProblemDetails
// @Router      /tasks/{id} [get]
func (h *TaskHandler) Get(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	task, err := h.usercase.GetTask(r.Context(), id)

	if err != nil {
		RespondError(w, r, h.logger, err)
		return
	}

	if err := RespondJSON(w, http.StatusOK, toTaskResponse(task)); err != nil {
		h.logger.ErrorContext(r.Context(),
			"error al escribir respuesta JSON",
			"error", err,
			"method", r.Method,
			"path", r.URL.Path)
	}
}

//ListByProject
//Assign
//UpdateStatus
