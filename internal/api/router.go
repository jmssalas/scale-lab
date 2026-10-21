package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Response struct {
	Success bool       `json:"success"`
	Data    any        `json:"data,omitempty"`
	Error   *ErrorInfo `json:"error,omitempty"`
}

type ErrorInfo struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func OK(c *gin.Context, data any) {
	c.JSON(http.StatusOK, Response{
		Success: true,
		Data:    data,
	})
}

func Fail(c *gin.Context, status int, code, message string) {
	c.JSON(status, Response{
		Success: false,
		Error:   &ErrorInfo{Code: code, Message: message},
	})
}

func Router(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	r.GET("/api/health", func(c *gin.Context) { OK(c, true) })

	eventHandler := EventHandler{db}
	v1 := r.Group("/api/v1")
	{
		v1.POST("/events", eventHandler.ProcessEvent)
	}

	return r
}
