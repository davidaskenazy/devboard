package handler

import (
	"encoding/json"
	"net/http"
)

// HealthHandler struct
type HelthHandler struct{}

// NewHealthHandler constructor
func NewHealthHandler() *HelthHandler {
	return &HelthHandler{}
}

// healthResponse struct
type healthResponse struct {
	Status  string `json:"status"`
	Version string `json:"version"`
}

func (handle *HelthHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	resp := healthResponse{
		Status:  "ok",
		Version: "1.0.0",
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK) //optional

	if err := json.NewEncoder(writer).Encode(resp); err != nil {
		return
	}
}
