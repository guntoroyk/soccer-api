package repository

import "github.com/guntoroyk/soccer-api/entity"

// TeamRepoItf is the interface for the TeamRepo struct
type TeamRepoItf interface {
	GetTeams() ([]*entity.Team, error)
	GetTeam(id int) (*entity.Team, error)
	CreateTeam(team *entity.Team) (*entity.Team, error)
	AddPlayerToTeam(teamID int, player *entity.Player) error
}

// PlayerRepoItf is the interface for the PlayerRepo struct
type PlayerRepoItf interface {
	GetPlayer(id int) (*entity.Player, error)
}
