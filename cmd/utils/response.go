package utils

import "github.com/gin-gonic/gin"

func ErrorResponse(c *gin.Context, statusCode int, errorMessage string){
	c.JSON(
		statusCode,
		gin.H{
			"success":false,
			"error":errorMessage,
		},
	)
}


func SuccessResponse(c *gin.Context, statusCode int, message string){
	c.JSON(
		statusCode,
		gin.H{
			"sucess":true,
			"message":message,
		},
	)
}