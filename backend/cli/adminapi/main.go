package main

import (
	"context"
	"github.com/Purnachndar/exnomy-scaffold-dex/backend/admin/api"
	"github.com/Purnachndar/exnomy-scaffold-dex/backend/cli"
	_ "github.com/joho/godotenv/autoload"
	"os"
)

func run() int {
	ctx, stop := context.WithCancel(context.Background())

	go cli.WaitExitSignal(stop)
	adminapi.StartServer(ctx)

	return 0
}

func main() {
	os.Exit(run())
}
