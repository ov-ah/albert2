package main

import (
	"context"
	"time"
	
	// "albert2/db"
	"albert2/cli"
	"albert2/api"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	api.Get(ctx, "/error-codes")

	rl, err := cli.Init()
	if err != nil { panic(err) }
	defer rl.Close()
	cli.Run(rl)
}
