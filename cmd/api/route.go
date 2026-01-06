package main

import (
	"net/http"

	"github.com/Walon-Foundation/go-gin-doc/cmd/api/routes"
	"github.com/gin-gonic/gin"
	swaggerFiles"github.com/swaggo/files"
	ginSwagger"github.com/swaggo/gin-swagger"
)

func (app *application) route()http.Handler{
	g := gin.Default()
	
	r := routes.Router()
	
	v1 := g.Group("/api/v1")
	{
		v1.POST("/auth/signup", r.RegisterUser)
		v1.POST("/auth/login", r.LoginUser)
	}
	
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	
	return g
}