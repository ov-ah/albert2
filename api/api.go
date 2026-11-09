package api

import (
	"context"
	"net/http"
	"io"
	// "encoding/json"
	"log"
	"time"

	"albert2/db"
	"albert2/cli"
)

var path_root = "https://api.spacetraders.io/v2"

func Get(ctx context.Context, path string) (string, error) {
	reqCtx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(reqCtx, "GET", path_root+path, nil)
	if err != nil {
		log.Fatal("Get:", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("GET:", err)
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	conn, err := db.OpenDB(reqCtx, "db/temp.db") // TODO: update db path
	if err != nil {
		log.Fatal("GET->db:", err)
		return "", err
	}
	db.LogRequest(reqCtx, conn, "", string(resp.Request.Method), path_root+path, resp.StatusCode, "", string(body))
	if err != nil {
		log.Fatal("GET->db:", err)
		return "", err
	}
	return string(body), nil
}

func Init(ctx context.Context) error {
	cli.RegisterCommand("get", func(args []string) error {
		if len(args) != 1 {
			log.Println("usage: get <path>")
			return nil
		}
		out, err := Get(ctx, args[0])
		if err != nil {
			return err
		}
		log.Println(out)
		return nil
	})
	return nil
}
