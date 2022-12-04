package inmemory

import (
	"reflect"
	"testing"

	"github.com/guntoroyk/soccer-api/entity"
	"github.com/guntoroyk/soccer-api/repository"
)

func TestNewTeamRepo(t *testing.T) {
	tests := []struct {
		name string
		want repository.TeamRepoItf
	}{
		{
			name: "create new team repo",
			want: &teamRepo{
				teams:        map[int]*entity.Team{},
				lastTeamID:   0,
				lastPlayerID: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTeamRepo(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTeamRepo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_teamRepo_GetTeams(t *testing.T) {
	type fields struct {
		teams        map[int]*entity.Team
		lastTeamID   int
		lastPlayerID int
	}
	tests := []struct {
		name    string
		fields  fields
		want    []*entity.Team
		wantErr bool
	}{
		{
			name: "get all teams",
			fields: fields{
				teams: map[int]*entity.Team{
					1: {
						ID:   1,
						Name: "Barcelona",
					},
					2: {
						ID:   2,
						Name: "Real Madrid",
					},
				},
				lastTeamID:   2,
				lastPlayerID: 0,
			},
			want: []*entity.Team{
				{
					ID:   1,
					Name: "Barcelona",
				},
				{
					ID:   2,
					Name: "Real Madrid",
				},
			},
			wantErr: false,
		},
		{
			name: "get all teams with no teams",
			fields: fields{
				teams:        map[int]*entity.Team{},
				lastTeamID:   0,
				lastPlayerID: 0,
			},
			want:    []*entity.Team{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &teamRepo{
				teams:        tt.fields.teams,
				lastTeamID:   tt.fields.lastTeamID,
				lastPlayerID: tt.fields.lastPlayerID,
			}
			got, err := tr.GetTeams()
			if (err != nil) != tt.wantErr {
				t.Errorf("teamRepo.GetTeams() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("teamRepo.GetTeams() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_teamRepo_GetTeam(t *testing.T) {
	type fields struct {
		teams        map[int]*entity.Team
		lastTeamID   int
		lastPlayerID int
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entity.Team
		wantErr bool
	}{
		{
			name: "success get team by id",
			fields: fields{
				teams: map[int]*entity.Team{
					1: {
						ID:   1,
						Name: "Barcelona",
					},
					2: {
						ID:   2,
						Name: "Real Madrid",
					},
				},
				lastTeamID:   2,
				lastPlayerID: 0,
			},
			args: args{
				id: 1,
			},
			want: &entity.Team{
				ID:   1,
				Name: "Barcelona",
			},
			wantErr: false,
		},
		{
			name: "error get team by id given id not found",
			fields: fields{
				teams: map[int]*entity.Team{
					1: {
						ID:   1,
						Name: "Barcelona",
					},
					2: {
						ID:   2,
						Name: "Real Madrid",
					},
				},
				lastTeamID:   2,
				lastPlayerID: 0,
			},
			args: args{
				id: 3,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &teamRepo{
				teams:        tt.fields.teams,
				lastTeamID:   tt.fields.lastTeamID,
				lastPlayerID: tt.fields.lastPlayerID,
			}
			got, err := tr.GetTeam(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("teamRepo.GetTeam() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("teamRepo.GetTeam() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_teamRepo_CreateTeam(t *testing.T) {
	type fields struct {
		teams        map[int]*entity.Team
		lastTeamID   int
		lastPlayerID int
	}
	type args struct {
		team *entity.Team
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entity.Team
		wantErr bool
	}{
		{
			name: "success create team",
			fields: fields{
				teams:        map[int]*entity.Team{},
				lastTeamID:   0,
				lastPlayerID: 0,
			},
			args: args{
				team: &entity.Team{
					Name: "Manchester United",
				},
			},
			want: &entity.Team{
				ID:   1,
				Name: "Manchester United",
			},
			wantErr: false,
		},
		{
			name: "error create team with input nil",
			fields: fields{
				teams:        map[int]*entity.Team{},
				lastTeamID:   0,
				lastPlayerID: 0,
			},
			args: args{
				team: nil,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &teamRepo{
				teams:        tt.fields.teams,
				lastTeamID:   tt.fields.lastTeamID,
				lastPlayerID: tt.fields.lastPlayerID,
			}
			got, err := tr.CreateTeam(tt.args.team)
			if (err != nil) != tt.wantErr {
				t.Errorf("teamRepo.CreateTeam() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("teamRepo.CreateTeam() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_teamRepo_AddPlayerToTeam(t *testing.T) {
	type fields struct {
		teams        map[int]*entity.Team
		lastTeamID   int
		lastPlayerID int
	}
	type args struct {
		teamID int
		player *entity.Player
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "success add player to team",
			fields: fields{
				teams: map[int]*entity.Team{
					1: {
						ID:      1,
						Name:    "Barcelona",
						Players: []*entity.Player{},
					},
				},
				lastTeamID:   1,
				lastPlayerID: 0,
			},
			args: args{
				teamID: 1,
				player: &entity.Player{
					Name: "Lionel Messi",
				},
			},
			wantErr: false,
		},
		{
			name: "error add player to team given team id not found",
			fields: fields{
				teams: map[int]*entity.Team{
					1: {
						ID:      1,
						Name:    "Barcelona",
						Players: []*entity.Player{},
					},
				},
				lastTeamID:   1,
				lastPlayerID: 0,
			},
			args: args{
				teamID: 2,
				player: &entity.Player{
					Name: "Lionel Messi",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &teamRepo{
				teams:        tt.fields.teams,
				lastTeamID:   tt.fields.lastTeamID,
				lastPlayerID: tt.fields.lastPlayerID,
			}
			if err := tr.AddPlayerToTeam(tt.args.teamID, tt.args.player); (err != nil) != tt.wantErr {
				t.Errorf("teamRepo.AddPlayerToTeam() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
