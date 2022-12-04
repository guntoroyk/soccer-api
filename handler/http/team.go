package http

import (
	"errors"
	"net/http"

	"github.com/guntoroyk/soccer-api/entity"
	"github.com/guntoroyk/soccer-api/lib/converter"
	"github.com/guntoroyk/soccer-api/usecase"
	"github.com/labstack/echo/v4"
)

type handler struct {
	teamUseCase usecase.TeamUsecaseItf
}

// NewHandler is a constructor for handler
func NewHandler(teamUseCase usecase.TeamUsecaseItf) *handler {
	return &handler{
		teamUseCase: teamUseCase,
	}
}

// GetTeams is a handler for GET /api/teams
func (h *handler) GetTeams(c echo.Context) error {
	teams, err := h.teamUseCase.GetTeams()

	resp := HttpResponse{
		Code: http.StatusOK,
		Data: teams,
	}

	if err != nil {
		resp.Code = http.StatusInternalServerError
		resp.Error = err.Error()
	}

	return c.JSON(resp.Code, resp)
}

// GetTeam is a handler for GET /api/teams/:id
func (h *handler) GetTeam(c echo.Context) error {
	id := c.Param("id")

	team, err := h.teamUseCase.GetTeam(converter.ToInt(id))

	resp := HttpResponse{}

	if errors.Is(err, entity.ErrTeamNotFound) {
		resp.Code = http.StatusNotFound
		resp.Error = err.Error()
	} else if err != nil {
		resp.Code = http.StatusInternalServerError
		resp.Error = err.Error()
	} else {
		resp.Code = http.StatusOK
		resp.Data = team
	}

	return c.JSON(resp.Code, resp)
}

// CreateTeam is a handler for POST /api/teams
func (h *handler) CreateTeam(c echo.Context) error {
	team := new(entity.Team)
	resp := HttpResponse{}

	if err := c.Bind(team); err != nil {
		resp.Code = http.StatusBadRequest
		resp.Error = err.Error()

		return c.JSON(resp.Code, resp)
	}

	newTeam, err := h.teamUseCase.CreateTeam(team)
	if err != nil {
		resp.Code = http.StatusInternalServerError
		resp.Error = err.Error()
	} else {
		resp.Code = http.StatusCreated
		resp.Data = newTeam
	}

	return c.JSON(resp.Code, resp)
}

// AddPlayerToTeam is a handler for POST /api/teams/:id/players
func (h *handler) AddPlayerToTeam(c echo.Context) error {
	player := new(entity.Player)
	resp := HttpResponse{}

	if err := c.Bind(player); err != nil {
		resp.Code = http.StatusBadRequest
		resp.Error = err

		return c.JSON(resp.Code, resp)
	}

	id := c.Param("id")

	err := h.teamUseCase.AddPlayerToTeam(converter.ToInt(id), player)
	if errors.Is(err, entity.ErrTeamNotFound) {
		resp.Code = http.StatusNotFound
		resp.Error = err.Error()
	} else if errors.Is(err, entity.ErrPlayerAlreadyInTeam) {
		resp.Code = http.StatusBadRequest
		resp.Error = err.Error()
	} else if err != nil {
		resp.Code = http.StatusInternalServerError
		resp.Error = err.Error()
	} else {
		resp.Code = http.StatusCreated
		resp.Data = player
	}

	return c.JSON(resp.Code, resp)
}

// GetPlayer is a handler for GET /api/teams/:id/players/:playerId
func (h *handler) GetPlayer(c echo.Context) error {
	id := c.Param("id")
	playerID := c.Param("playerId")
	resp := HttpResponse{}

	player, err := h.teamUseCase.GetPlayer(converter.ToInt(id), converter.ToInt(playerID))
	if errors.Is(err, entity.ErrPlayerNotFound) || errors.Is(err, entity.ErrTeamNotFound) {
		resp.Code = http.StatusNotFound
		resp.Error = err.Error()
	} else if err != nil {
		resp.Code = http.StatusInternalServerError
		resp.Error = err.Error()
	} else {
		resp.Code = http.StatusOK
		resp.Data = player
	}

	return c.JSON(resp.Code, resp)
}
