package utils

import "github.com/gin-gonic/gin"

type SuccessResponseSchema struct {
	Success bool        `json:"success" example:"true"`
	Message string      `json:"message" example:"User created"`
	Data    any `json:"data"`
}

// ErrorResponseSchema represents an error API response
type ErrorResponseSchema struct {
	Success bool   `json:"success" example:"false"`
	Error   string `json:"error" example:"Invalid request"`
}


func ErrorResponse(c *gin.Context, statusCode int, errorMessage string){
	c.JSON(
		statusCode,
		gin.H{
			"success":false,
			"error":errorMessage,
		},
	)
}


func SuccessResponse(c *gin.Context, statusCode int, message string, data ...any){
	c.JSON(
		statusCode,
		gin.H{
			"sucess":true,
			"message":message,
			"data": data,
		},
	)
}