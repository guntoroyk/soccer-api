package inmemory

import (
	"github.com/guntoroyk/soccer-api/entity"
	"github.com/guntoroyk/soccer-api/repository"
)

type teamRepo struct {
	teams        map[int]*entity.Team
	lastTeamID   int
	lastPlayerID int
}

// NewTeamRepo will create new an TeamRepo object representation of TeamRepoItf interface
func NewTeamRepo() repository.TeamRepoItf {
	return &teamRepo{
		teams:        map[int]*entity.Team{},
		lastTeamID:   0,
		lastPlayerID: 0,
	}
}

// GetTeams will get all teams
func (t *teamRepo) GetTeams() ([]*entity.Team, error) {
	teams := []*entity.Team{}

	for _, team := range t.teams {
		teams = append(teams, team)
	}

	return teams, nil
}

// GetTeam will get team by id
func (t *teamRepo) GetTeam(id int) (*entity.Team, error) {
	team := t.teams[id]
	if team == nil {
		return nil, entity.ErrTeamNotFound
	}

	return team, nil
}

// SetTeam will set a team
func (t *teamRepo) CreateTeam(team *entity.Team) (*entity.Team, error) {
	if team == nil {
		return nil, entity.ErrTeamCannotBeNil
	}

	team.ID = t.lastTeamID + 1
	t.teams[team.ID] = team

	t.lastTeamID++

	return team, nil
}

// AddPlayerToTeam will add player to team
func (t *teamRepo) AddPlayerToTeam(teamID int, player *entity.Player) error {
	team := t.teams[teamID]

	if team == nil {
		return entity.ErrTeamNotFound
	}

	player.ID = t.lastPlayerID + 1

	team.Players = append(team.Players, player)

	t.lastPlayerID++

	return nil
}
