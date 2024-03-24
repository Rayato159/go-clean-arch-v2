package main

import (
	"github.com/Rayato159/go-clean-arch-v2/config"
	"github.com/Rayato159/go-clean-arch-v2/database"
	"github.com/Rayato159/go-clean-arch-v2/server"
)

func main() {
	conf := config.GetConfig()

	db := database.NewPostgresDatabase(conf)

	server.NewEchoServer(conf, db.GetDb()).Start()
}
