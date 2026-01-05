package models

import "github.com/jackc/pgx/v5"

type UserModel struct {
	Db *pgx.Conn
}