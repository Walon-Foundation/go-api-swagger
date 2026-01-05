package db

import (
	"context"

	"github.com/Walon-Foundation/go-gin-doc/cmd/api/config"
	"github.com/jackc/pgx/v5"
)


func Insert(ctx context.Context, sql string, args ...any) (pgx.Row, error) {
    row := config.Db.QueryRow(ctx, sql, args...)
    return row, nil
}


func SelectOne(ctx context.Context, sql string, args ...any)(pgx.Row){
	return config.Db.QueryRow(ctx,sql,args...)
}

func SelectMany(ctx context.Context, sql string, args ...any)(pgx.Rows, error){
	return config.Db.Query(ctx, sql, args...)
}

// Exec runs any statement (no rows expected)
func Delete(ctx context.Context, sql string, args ...any) error {
    _, err := config.Db.Exec(ctx, sql, args...)
    return err
}

func Update(ctx context.Context, sql string, args ...any)pgx.Row{
	return config.Db.QueryRow(ctx, sql, args...)
}