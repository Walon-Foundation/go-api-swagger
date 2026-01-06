package models

import (
	"context"

	"github.com/Walon-Foundation/go-gin-doc/cmd/api/config"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
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

func (m *EventModel) GetOneEventById(ctx context.Context, id string)(pgx.Row){
	sql := "SELECT id,name,creator_id, created_at FROM events WHERE id = $1"
	
	return config.Db.QueryRow(ctx, sql, id)
}

func (m *EventModel)DeleteEvent(ctx context.Context, id, creator_id string)(pgconn.CommandTag, error){
	sql := "DELETE FROM events WHERE id = $1 AND creator_id = $2"
	return config.Db.Exec(ctx, sql, id,creator_id)	
}

func (m *EventModel) UpdateEvent(ctx context.Context, name,id,creator_id string)(pgx.Row){
	sql := "UPDATE events SET name = $1 WHERE id = $2 AND creator_id = $3 RETURNING id, name, creator_id, created_at"
	return config.Db.QueryRow(ctx,sql, name, id, creator_id)
}