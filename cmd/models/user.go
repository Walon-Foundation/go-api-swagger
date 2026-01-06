package models

import (
	"context"
	"time"

	"github.com/Walon-Foundation/go-gin-doc/cmd/api/config"
	"github.com/jackc/pgx/v5"
)

type UserModel struct {
	Db *pgx.Conn
}

type User struct {
	Id string `json:"id"`
	Name string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (m *UserModel) GetUser(ctx context.Context,name,username string)(pgx.Row){

	sql := "SELECT username FROM users WHERE name = $1 AND username = $2"
	
	row := config.Db.QueryRow(ctx,sql, name, username)
	return row
}

func (m *UserModel) InsertUser(ctx context.Context, args ...string)(pgx.Row){
	sql := "INSERT INTO users (id,name,username,password) VALUES ($1,$2,$3,$4) RETURNING id"
	
	row := config.Db.QueryRow(ctx,sql, args)
	
	return row
}

func (m *UserModel) GetUserByUsername(ctx context.Context,username string)(pgx.Row){

	sql := "SELECT id,username,password FROM users WHERE username = $1"
	
	row := config.Db.QueryRow(ctx,sql, username)
	return row
}