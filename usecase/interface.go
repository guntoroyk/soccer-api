package usecase

import "github.com/guntoroyk/soccer-api/entity"

// TeamUsecaseItf is the interface for the TeamUsecase struct
type TeamUsecaseItf interface {
	GetTeams() ([]*entity.Team, error)
	GetTeam(id int) (*entity.Team, error)
	CreateTeam(team *entity.Team) (*entity.Team, error)
	AddPlayerToTeam(teamID int, player *entity.Player) error
	GetPlayer(teamID int, playerID int) (*entity.Player, error)
}
