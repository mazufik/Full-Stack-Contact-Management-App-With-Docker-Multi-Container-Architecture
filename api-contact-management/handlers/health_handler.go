package handlers

import (
	"api-contact-management/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthHandler handles HTTP requests related to health checks.
type HealthHandler struct{}

// NewHealthHandler creates a new instance of HealthHandler.
func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

// HealthCheck responds with a simple message indicating that the API is running.
//
// It returns a JSON response with a 200 OK status code and a message
// confirming the operational status of the API.
//
// Example Response:
//
//	{
//	    "code": "SUCCESS",
//	    "message": "API is running."
//	}
func (h *HealthHandler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, responses.APIResponse{
		Code:    "SUCCESS",
		Message: "API is running.",
	})
}
