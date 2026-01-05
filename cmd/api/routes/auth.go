package routes

import (
	"net/http"

	"github.com/Walon-Foundation/go-gin-doc/cmd/utils"
	"github.com/gin-gonic/gin"
)

type user struct {
	Name string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=8"`
}

func(r *route) RegisterUser(c *gin.Context){
	var registerRequest user
	if err := c.ShouldBindJSON(&registerRequest); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request body")
		return
	}
}