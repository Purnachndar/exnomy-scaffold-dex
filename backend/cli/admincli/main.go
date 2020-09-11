package main

import (
	"github.com/Purnachndar/hydro-scaffold-dex/backend/admin/cli"
	"github.com/Purnachndar/hydro-sdk-backend/utils"
	_ "github.com/joho/godotenv/autoload"
	"os"
)

func main() {
	cli := admincli.NewDexCli()
	err := cli.Run(os.Args)
	if err != nil {
		utils.Errorf(err.Error())
	}
}
