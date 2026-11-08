package api

import (
	"context"
	"net/http"
	"io"
	// "encoding/json"
	"log"

	"albert2/db"
)

var path_root = "https://api.spacetraders.io/v2"

func Get(ctx context.Context, path string) (string, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", path_root+path, nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("GET:", err)
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	conn, err := db.OpenDB(ctx, "db/temp.db") // TODO: update db path
	if err != nil {
		log.Fatal("GET->db:", err)
		return "", err
	}
	db.LogRequest(ctx, conn, "", string(resp.Request.Method), path_root+path, resp.StatusCode, "", string(body))
	if err != nil {
		log.Fatal("GET->db:", err)
		return "", err
	}
	return string(body), nil
}
