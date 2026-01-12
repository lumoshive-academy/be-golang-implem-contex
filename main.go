package main

import (
	"session-26/cmd"
	"session-26/internal/data/repository"
	"session-26/internal/wire"
	"session-26/pkg/database"
	"session-26/pkg/utils"
)

func main() {
	config, err := utils.ReadConfiguration()
	if err != nil {
		// print error
	}

	db, err := database.InitDB(config.DB)
	if err != nil {
		// print err
	}

	repo := repository.NewRepository(db)
	route := wire.Wiring(*repo)
	cmd.APiserver(route)
}
