package models

import "github.com/jackc/pgx/v5"

type Models struct {
	User UserModel 
	Event EventModel
	Attendee AttendeeModel
}

func NewModel(db *pgx.Conn)Models{
	return Models{
		User: UserModel{ Db:db},
		Event: EventModel{Db:db},
		Attendee: AttendeeModel{Db:db},
	}
}