package models

import (
	"context"

	"github.com/Walon-Foundation/go-gin-doc/cmd/api/config"
	"github.com/jackc/pgx/v5"
)

type EventModel struct{
	Db *pgx.Conn
}

func (m *EventModel)GetEvent(ctx context.Context)(pgx.Rows, error){
	sql := "SELECT * FROM events"
	return config.Db.Query(ctx, sql)	
}

func(m *EventModel)GetEventByNameAndCreatorId(ctx context.Context, name,creatorId string)(pgx.Row){
	sql := "SELECT name FROM events WHERE name = $1 AND creator_id = $2"
	
	return config.Db.QueryRow(ctx, sql, name,creatorId)
}


func (m *EventModel)CreateEvent(ctx context.Context, args ...string)(pgx.Row){
	sql := "INSERT INTO events (id, name, creator_id) VALUES ($1,$2,$3) RETURNING id"
	return config.Db.QueryRow(ctx, sql, args)
}