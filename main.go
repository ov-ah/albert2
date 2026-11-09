package main

import (
	"context"
	"os/signal"
	"os"
	"log"
	
	"albert2/db"
	"albert2/cli"
	"albert2/api"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	conn, err := db.OpenDB(ctx, "db/logs.db")
	if err != nil {
		log.Fatal("POST->db:", err)
	}
	token, err := db.GetAuth(ctx, conn)
	if err != nil { panic(err) }
	api.SetAuth(token)
	api.Init(ctx)

	rl, err := cli.Init()
	if err != nil { panic(err) }
	defer rl.Close()
	cli.Run(rl)
}
