package main

import (
	"net/http"

	"github.com/Walon-Foundation/go-gin-doc/cmd/api/routes"
	"github.com/gin-gonic/gin"
)

func (app *application) route()http.Handler{
	g := gin.Default()
	
	v1 := g.Group("/api/v1")
	{
		v1.POST("/auth/signup", routes.Router().RegisterUser)
	}
	
	return g
}