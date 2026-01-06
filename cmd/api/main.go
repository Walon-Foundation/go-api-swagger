
// @title           Go Gin Doc
// @version         1.0
// @description     This is an api server for making events and attending events.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  MIT
// @license.url   https://mit-license.org/

// @host      localhost:5000
// @BasePath  /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization


package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Walon-Foundation/go-gin-doc/cmd/api/config"
	"github.com/Walon-Foundation/go-gin-doc/cmd/utils"
	"github.com/joho/godotenv"
	_"github.com/Walon-Foundation/go-gin-doc/docs"
)

type application struct{
	Port int
	JwtSecret string
}

func (app *application) serve() error{
	server := &http.Server{
		Addr: fmt.Sprintf(":%d",app.Port),
		Handler: app.route(),
		IdleTimeout: time.Second * 30,
		ReadTimeout: time.Second * 30,
		WriteTimeout: time.Second * 20,
	}
	
	fmt.Printf("server is running on http://localhost:%d\n",app.Port)
	return server.ListenAndServe()
}

func init(){
	if err := godotenv.Load();err != nil {
		log.Fatal("failed to load .env file",err)
	}
	config.LoadDb()
}


func main(){
	portValue := utils.GetEnvInt("PORT", 5000)
	jwtSecretValue := utils.GetEnvString("JWT_SECRET","HELLLL")
	
	app := application{
		Port: portValue,
		JwtSecret: jwtSecretValue,
	}
	
	if err := app.serve(); err != nil {
		log.Fatal(err)
	}
}
