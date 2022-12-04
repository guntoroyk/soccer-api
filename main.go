package main

import (
	"github.com/guntoroyk/soccer-api/config"
	httpHandler "github.com/guntoroyk/soccer-api/handler/http"
	"github.com/guntoroyk/soccer-api/repository/inmemory"
	"github.com/guntoroyk/soccer-api/usecase"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	config := config.GetConfig()
	e := echo.New()

	teamRepo := inmemory.NewTeamRepo()
	teamUsecase := usecase.NewTeamUsecase(teamRepo)
	handler := httpHandler.NewHandler(teamUsecase)

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/api/teams", handler.GetTeams)
	e.POST("/api/teams", handler.CreateTeam)
	e.GET("/api/teams/:id", handler.GetTeam)
	e.POST("/api/teams/:id/players", handler.AddPlayerToTeam)
	e.GET("/api/teams/:id/players/:playerId", handler.GetPlayer)

	e.Logger.Fatal(e.Start(config.Host + ":" + config.PORT))
}
