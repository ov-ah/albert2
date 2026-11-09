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

func LogRequest(ctx context.Context, conn *sql.DB, resetDate, method string, path string, status int, reqBody string, respBody string) error {
	dbCtx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()
	_, err := conn.ExecContext(
		dbCtx, 
		`INSERT INTO api_log (reset_date, method, path, status, request_body, response_body)
		VALUES (?, ?, ?, ?, ?, ?)`,
		resetDate, method, path, status, reqBody, respBody,
	)
	return err
}

func GetAuth(ctx context.Context, conn *sql.DB, resetDate string) (string, error) {
	var auth string
	err := conn.QueryRowContext(ctx, `
	SELECT json_extract(response_body, '$.data.token')
	FROM api_log
	WHERE path LIKE '%/register' 
		AND status = 201
		AND reset_date IS NOT ?
	ORDER BY id DESC LIMIT 1`, resetDate).Scan(&auth)

	return auth, err
}

func ClearAuth(ctx context.Context, conn *sql.DB, resetDate string) (int64, error) {
	res, err := conn.ExecContext(ctx, `
		UPDATE api_log
		SET response_body = json_remove(response_body, '$.data.token')
		WHERE path LIKE '%/register'
			AND reset_date IS NOT ?
		  	AND json_valid(response_body)
		  	AND json_extract(response_body, '$.data.token') IS NOT NULL`,
		resetDate)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
