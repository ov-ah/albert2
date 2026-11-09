package main

import (
	"context"
	"os/signal"
	"os"
	
	// "albert2/db"
	"albert2/cli"
	"albert2/api"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	api.Init(ctx)

	rl, err := cli.Init()
	if err != nil { panic(err) }
	defer rl.Close()
	cli.Run(rl)
}
