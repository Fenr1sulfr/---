package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func Connect(host string, port int, user, password, dbname string) (*pgx.Conn, error) {
	url := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", user, password, host, port, dbname)
	conn, err := pgx.Connect(context.Background(), url)

	if err != nil {
		return nil, err
	}

	return conn, nil
}
