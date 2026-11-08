package db

import (
	"database/sql"
	"context"
	"time"

	_ "modernc.org/sqlite"
)

const schema = `
CREATE TABLE IF NOT EXISTS api_log (
	id INTEGER PRIMARY KEY,
	reset_date TEXT,
	method TEXT,
	path TEXT,
	status INTEGER,
	request_body TEXT,
	response_body TEXT,
	timestamp TEXT DEFAULT CURRENT_TIMESTAMP
);`

func OpenDB(ctx context.Context, path string) (*sql.DB, error) {
	dbCtx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()
	db, err := sql.Open("sqlite", path)
	if err != nil {
		db.Close()
		return nil, err
	}

	_, err = db.ExecContext(dbCtx, schema)
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, err
}

func LogRequest(ctx context.Context, db *sql.DB, resetDate, method string, path string, status int, reqBody string, respBody string) error {
	dbCtx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()
	_, err := db.ExecContext(
		dbCtx, 
		`INSERT INTO api_log (reset_date, method, path, status, request_body, response_body)
		VALUES (?, ?, ?, ?, ?, ?)`,
		resetDate, method, path, status, reqBody, respBody,
	)
	return err
}
