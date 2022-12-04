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
	if player == nil {
		return entity.ErrPlayerCannotBeNil
	}

	if player.Name == "" {
		return entity.ErrPlayerNameIsRequired
	}

	// check if player already exist on a team
	_, err := t.getPlayerTeam(player.Name)
	if err == nil {
		return entity.ErrPlayerAlreadyInTeam
	}

	// validate teamID
	team, err := t.teamRepo.GetTeam(teamID)
	if err != nil {
		return err
	}

	// add player to team
	return t.teamRepo.AddPlayerToTeam(team.ID, player)
}

func (t *teamUsecase) getPlayerTeam(playerName string) (*entity.Team, error) {
	// get teams
	teams, err := t.teamRepo.GetTeams()
	if err != nil {
		return nil, err
	}

	// check if player already exist on a team
	for _, team := range teams {
		for _, p := range team.Players {
			if p.Name == playerName {
				return team, nil
			}
		}
	}

	return nil, entity.ErrPlayerNotFound
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
