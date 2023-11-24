package server

import (
	"fmt"

	"github.com/Rayato159/go-clean-arch-v2/cockroach/handlers"
	cockroachRepositories "github.com/Rayato159/go-clean-arch-v2/cockroach/repositories"
	cockroachUsecases "github.com/Rayato159/go-clean-arch-v2/cockroach/usecases"
	"github.com/Rayato159/go-clean-arch-v2/config"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type echoServer struct {
	app *echo.Echo
	db  *gorm.DB
	cfg *config.Config
}

func NewEchoServer(cfg *config.Config, db *gorm.DB) Server {
	return &echoServer{
		app: echo.New(),
		db:  db,
		cfg: cfg,
	}
}

func (s *echoServer) Start() {
	s.initializeCockroachHttpHandler()

	serverUrl := fmt.Sprintf(":%d", s.cfg.App.Port)
	s.app.Logger.Fatal(s.app.Start(serverUrl))
}

func (s *echoServer) initializeCockroachHttpHandler() {
	// Initialize all layers
	cockroachPostgresRepository := cockroachRepositories.NewCockroachPostgresRepository(s.db)
	cockroachFCMMessaging := cockroachRepositories.NewCockroachFCMMessaging()

	cockroachUsecase := cockroachUsecases.NewCockroachUsecaseImpl(
		cockroachPostgresRepository,
		cockroachFCMMessaging,
	)

	cockroachHttpHandler := handlers.NewCockroachHttpHandler(cockroachUsecase)

	// Routers
	cockroachRouters := s.app.Group("/cockroach")
	cockroachRouters.POST("/", cockroachHttpHandler.DetectCockroach)
}
