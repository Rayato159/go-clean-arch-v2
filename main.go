package main

import (
	"github.com/Rayato159/go-clean-arch-v2/config"
	"github.com/Rayato159/go-clean-arch-v2/database"
	"github.com/Rayato159/go-clean-arch-v2/server"
)

func main() {
	cfg := config.GetConfig()

	db := database.NewPostgresDatabase(&cfg)

	server.NewEchoServer(&cfg, db.GetDb()).Start()
}
