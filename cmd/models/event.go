package models

import "github.com/jackc/pgx/v5"

type EventModel struct{
	Db *pgx.Conn
}