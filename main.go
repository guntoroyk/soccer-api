package main

import (
	"github.com/labstack/echo/v4"

	httpHandler "github.com/guntoroyk/soccer-api/handler/http"
	"github.com/guntoroyk/soccer-api/repository/inmemory"
	"github.com/guntoroyk/soccer-api/usecase"
)

func main() {
	e := echo.New()

	teamRepo := inmemory.NewTeamRepo()
	teamUsecase := usecase.NewTeamUsecase(teamRepo)
	handler := httpHandler.NewHandler(teamUsecase)

	e.GET("/api/teams", handler.GetTeams)
	e.POST("/api/teams", handler.CreateTeam)
	e.GET("/api/teams/:id", handler.GetTeam)
	e.POST("/api/teams/:id/players", handler.AddPlayerToTeam)
	e.GET("/api/teams/:id/players/:playerId", handler.GetPlayer)

	e.Logger.Fatal(e.Start(":8000"))
}
