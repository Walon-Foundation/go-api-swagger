package config

import (
	"context"
	"fmt"
	"log"

	"github.com/Walon-Foundation/go-gin-doc/cmd/utils"
	"github.com/jackc/pgx/v5"
	_"github.com/joho/godotenv/autoload"
)


var Db *pgx.Conn

func LoadDb(){
	database_url := utils.GetEnvString("DATABASE_URL", "postgres://walon:password@localhost:5432/app_db?sslmode=disable")
	ctx := context.Background()
	
	db,err := pgx.Connect(ctx,database_url)
	if err != nil {
		log.Fatal("Database connection error",err)
	}
	
	defer db.Close(ctx)
	
	Db = db
	
	fmt.Println("Database connected")
}