package main

import (
	"context"
	"github.com/Purnachndar/exnomy-scaffold-dex/backend/api"
	"github.com/Purnachndar/exnomy-scaffold-dex/backend/cli"
	"github.com/Purnachndar/exnomy-sdk-backend/utils"
	_ "github.com/joho/godotenv/autoload"
	"os"
)

func run() int {
	ctx, stop := context.WithCancel(context.Background())

	go cli.WaitExitSignal(stop)
	api.StartServer(ctx, utils.StartMetrics)

	return 0
}

func main() {
	os.Exit(run())
}
