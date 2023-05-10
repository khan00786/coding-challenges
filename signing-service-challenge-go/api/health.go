package api

import (
	"net/http"

	"github.com/fiskaly/coding-challenges/signing-service-challenge/domain"
)

// Health evaluates the health of the service and writes a standardized response.
func (s *Server) Health(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		MethodNotAllowedError(response, "Invalid Status Method")
		return
	}

	health := domain.HealthResponse{
		Status:  "pass",
		Version: "v0",
	}

	WriteAPIResponse(response, http.StatusOK, health)
}
