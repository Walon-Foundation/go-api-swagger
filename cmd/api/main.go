package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Walon-Foundation/go-gin-doc/cmd/api/config"
	"github.com/Walon-Foundation/go-gin-doc/cmd/utils"
	"github.com/joho/godotenv"
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
