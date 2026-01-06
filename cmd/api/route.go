package main

import (
	"net/http"

	"github.com/Walon-Foundation/go-gin-doc/cmd/api/middlewares"
	"github.com/Walon-Foundation/go-gin-doc/cmd/api/routes"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (app *application) route()http.Handler{
	g := gin.Default()
	
	r := routes.Router()
	
	v1 := g.Group("/api/v1")
	{
		//Auth router
		v1.POST("/auth/signup", r.RegisterUser)
		v1.POST("/auth/login", r.LoginUser)
		
		//event route
		v1.GET("/events", middlewares.AuthMiddleware(), r.GetEvent)
		v1.POST("/events",middlewares.AuthMiddleware(), r.CreateEvent)
		v1.GET("/events:id", middlewares.AuthMiddleware(), r.GetOneEvent)
		v1.DELETE("/events:id",middlewares.AuthMiddleware(), r.DeleteEvent)
		v1.PUT("/events:id", middlewares.AuthMiddleware(),r.UpdateEvent)
	}
	
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	
	return g
}