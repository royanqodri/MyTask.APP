package helpers

import (
	"github.com/gin-gonic/gin"
)

type MapResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// WebResponse returns a MapResponse
func WebResponse(c *gin.Context, code int, message string, data interface{}) {
	response := MapResponse{
		Code:    code,
		Message: message,
		Data:    data,
	}

	c.JSON(code, response)
	c.Header("Content-Type", "application/json")
}
