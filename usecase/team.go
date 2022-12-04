package usecase

import (
	"github.com/guntoroyk/soccer-api/entity"
	"github.com/guntoroyk/soccer-api/repository"
)

type teamUsecase struct {
	teamRepo repository.TeamRepoItf
}

// NewTeamUsecase will create new an teamUsecase object representation of TeamUsecaseItf interface
func NewTeamUsecase(teamRepo repository.TeamRepoItf) TeamUsecaseItf {
	return &teamUsecase{
		teamRepo: teamRepo,
	}
}

// GetTeams will get all teams
func (t *teamUsecase) GetTeams() ([]*entity.Team, error) {
	return t.teamRepo.GetTeams()
}

// GetTeam will get team by id
func (t *teamUsecase) GetTeam(id int) (*entity.Team, error) {
	return t.teamRepo.GetTeam(id)
}

// CreateTeam will create new team
func (t *teamUsecase) CreateTeam(team *entity.Team) (*entity.Team, error) {
	if team == nil {
		return nil, entity.ErrTeamCannotBeNil
	}

	if team.Name == "" {
		return nil, entity.ErrTeamNameCannotBeEmpty
	}

	return t.teamRepo.CreateTeam(team)
}

// AddPlayerToTeam will add player to team
func (t *teamUsecase) AddPlayerToTeam(teamID int, player *entity.Player) error {
	// // get teams
	// _, err := t.teamRepo.GetTeams()
	// if err != nil {
	// 	return err
	// }

	// // check if player already exist on a team
	// for _, team := range teams {
	// 	for _, p := range team.Players {
	// 		if p.ID == player.ID {
	// 			return entity.ErrPlayerAlreadyInTeam
	// 		}
	// 	}
	// }

	// get team
	team, err := t.teamRepo.GetTeam(teamID)
	if err != nil {
		return err
	}

	// add player to team
	return t.teamRepo.AddPlayerToTeam(team.ID, player)
}

// GetPlayer will get player by id
func (t *teamUsecase) GetPlayer(teamID int, playerID int) (*entity.Player, error) {
	team, err := t.teamRepo.GetTeam(teamID)
	if err != nil {
		return nil, err
	}

	for _, player := range team.Players {
		if player.ID == playerID {
			return player, nil
		}
	}

	return nil, entity.ErrPlayerNotFound
}
