package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"

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
	Title     string `json:"title" validate:"required, min=3, max=200"`
}

// Create metodo para crear tarea
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
	if err := RespondJSON(w, http.StatusCreated, task); err != nil {
		h.logger.ErrorContext(r.Context(),
			"error al escribir respuesta JSON de tarea",
			"error", err,
			"method", r.Method,
			"path", r.URL.Path,
		)
	}
}
