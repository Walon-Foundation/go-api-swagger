package models

import "github.com/jackc/pgx/v5"

type AttendeeModel struct {
	Db *pgx.Conn
}

type Attendee struct {
	ID string `json:"id"`
}

