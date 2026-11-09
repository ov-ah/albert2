package api

import (
	"context"
	"net/http"
	"io"
	"encoding/json"
	"log"
	"time"
	"bytes"

	"albert2/db"
	"albert2/cli"
)

var path_root = "https://api.spacetraders.io/v2"
var client = &http.Client{Timeout: 10 * time.Second}
var reset_date = ""

func Pretty(s string) string {
	var buf bytes.Buffer
	if err := json.Indent(&buf, []byte(s), "", " "); err != nil {
		return s
	}
	return buf.String()
}

func Get(ctx context.Context, path string) (string, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", path_root+path, nil)
	if err != nil {
		log.Fatal("GET:", err)
	}

	req.Header.Set("Content-type", "application/json")
	// TODO: add auth header

	resp, err := client.Do(req)
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
	err = db.LogRequest(ctx, conn, reset_date, string(resp.Request.Method), path_root+path, resp.StatusCode, "", string(body))
	if err != nil {
		log.Fatal("GET->db:", err)
		return "", err
	}
	return string(body), nil
}

func Post(ctx context.Context, path string, reqBody map[string]any) (string, error) {
	data, err := json.Marshal(reqBody)
	if err != nil {
		log.Fatal("POST:", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", path_root+path, bytes.NewReader(data))
	if err != nil {
		log.Fatal("POST:", err)
	}

	req.Header.Set("Content-type", "application/json")
	// TODO: add auth header

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("POST:", err)
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	conn, err := db.OpenDB(ctx, "db/temp.db") // TODO: update db path
	if err != nil {
		log.Fatal("POST->db:", err)
		return "", err
	}
	err = db.LogRequest(ctx, conn, reset_date, string(resp.Request.Method), path_root+path, resp.StatusCode, string(data), string(body))
	if err != nil {
		log.Fatal("POST->db:", err)
		return "", err
	}
	return string(body), nil
}

func Init(ctx context.Context) error {
	reset_body, err := Get(ctx, "")
	if err != nil {
		log.Fatal("Space Traders API is donw")
		return nil
	}
	var reset_resp StatusResponse
	json.Unmarshal([]byte(reset_body), &reset_resp)
	reset_date = reset_resp.ResetDate

	cli.RegisterCommand("get", func(args []string) error {
		if len(args) != 1 {
			log.Println("usage: get <path>")
			return nil
		}
		out, err := Get(ctx, args[0])
		if err != nil {
			return err
		}
		log.Println(Pretty(out))
		return nil
	})

	cli.RegisterCommand("post", func(args []string) error {
		if len(args) < 1 {
			log.Println("usage: post <path> <body>(optional)")
			return nil
		}
		m := make(map[string]any)
		if len(args) < 2 {
			for i := 1; i+1 < len(args); i += 2 {
				m[args[i]] = args[i+1]
			}
		}
		out, err := Post(ctx, args[0], m)
		if err != nil {
			return err
		}
		log.Println(Pretty(out))
		return nil
	})

	return nil
}
