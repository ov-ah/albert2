package main

import (
	"context"
	"time"
	
	"albert2/db"
	"albert2/cli"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := db.OpenDB(ctx, "db/temp.db")
	if err != nil { panic(err) }
	err = db.LogRequest(ctx, conn, "10-10-2000", "GET", "/test/test", 201, "test request", "test response")
	if err != nil { panic(err) }


	rl, err := cli.Init()
	if err != nil { panic(err) }
	defer rl.Close()
	cli.Run(rl)
}
