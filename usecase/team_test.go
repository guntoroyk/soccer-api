package usecase

import (
	"errors"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/guntoroyk/soccer-api/entity"
	"github.com/guntoroyk/soccer-api/mocks"
	"github.com/guntoroyk/soccer-api/repository"
)

func TestNewTeamUsecase(t *testing.T) {
	type args struct {
		teamRepo repository.TeamRepoItf
	}
	tests := []struct {
		name string
		args args
		want TeamUsecaseItf
	}{
		{
			name: "success create new team usecase",
			args: args{
				teamRepo: &mocks.MockTeamRepoItf{},
			},
			want: &teamUsecase{
				teamRepo: &mocks.MockTeamRepoItf{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTeamUsecase(tt.args.teamRepo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTeamUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_teamUsecase_GetTeam(t *testing.T) {
	type fields struct {
		teamRepo func(ctrl *gomock.Controller) repository.TeamRepoItf
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
			name: "success get team",
			fields: fields{
				teamRepo: func(ctrl *gomock.Controller) repository.TeamRepoItf {
					mockTeamRepo := mocks.NewMockTeamRepoItf(ctrl)
					mockTeamRepo.EXPECT().GetTeam(gomock.Any()).Return(&entity.Team{
						ID:   1,
						Name: "Manchester United",
					}, nil)
					return mockTeamRepo
				},
			},
			args: args{
				id: 1,
			},
			want: &entity.Team{
				ID:   1,
				Name: "Manchester United",
			},
			wantErr: false,
		},
		{
			name: "failed get team",
			fields: fields{
				teamRepo: func(ctrl *gomock.Controller) repository.TeamRepoItf {
					mockTeamRepo := mocks.NewMockTeamRepoItf(ctrl)
					mockTeamRepo.EXPECT().GetTeam(gomock.Any()).Return(nil, entity.ErrTeamNotFound)
					return mockTeamRepo
				},
			},
			args: args{
				id: 1,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			tr := &teamUsecase{
				teamRepo: tt.fields.teamRepo(ctrl),
			}
			got, err := tr.GetTeam(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("teamUsecase.GetTeam() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("teamUsecase.GetTeam() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_teamUsecase_GetTeams(t *testing.T) {
	type fields struct {
		teamRepo func(ctrl *gomock.Controller) repository.TeamRepoItf
	}
	tests := []struct {
		name    string
		fields  fields
		want    []*entity.Team
		wantErr bool
	}{
		{
			name: "success get teams",
			fields: fields{
				teamRepo: func(ctrl *gomock.Controller) repository.TeamRepoItf {
					mockTeamRepo := mocks.NewMockTeamRepoItf(ctrl)
					mockTeamRepo.EXPECT().GetTeams().Return([]*entity.Team{
						{
							ID:   1,
							Name: "Manchester United",
						},
						{
							ID:   2,
							Name: "Manchester City",
						},
					}, nil)
					return mockTeamRepo
				},
			},
			want: []*entity.Team{
				{
					ID:   1,
					Name: "Manchester United",
				},
				{
					ID:   2,
					Name: "Manchester City",
				},
			},
			wantErr: false,
		},
		{
			name: "failed get teams",
			fields: fields{
				teamRepo: func(ctrl *gomock.Controller) repository.TeamRepoItf {
					mockTeamRepo := mocks.NewMockTeamRepoItf(ctrl)
					mockTeamRepo.EXPECT().GetTeams().Return(nil, errors.New("some error"))
					return mockTeamRepo
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			tr := &teamUsecase{
				teamRepo: tt.fields.teamRepo(ctrl),
			}
			got, err := tr.GetTeams()
			if (err != nil) != tt.wantErr {
				t.Errorf("teamUsecase.GetTeams() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("teamUsecase.GetTeams() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_teamUsecase_CreateTeam(t *testing.T) {
	type fields struct {
		teamRepo func(ctrl *gomock.Controller) repository.TeamRepoItf
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
				teamRepo: func(ctrl *gomock.Controller) repository.TeamRepoItf {
					mockTeamRepo := mocks.NewMockTeamRepoItf(ctrl)
					mockTeamRepo.EXPECT().CreateTeam(gomock.Any()).Return(&entity.Team{
						ID:   1,
						Name: "Manchester United",
					}, nil)
					return mockTeamRepo
				},
			},
			args: args{
				team: &entity.Team{
					ID:   1,
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
			name: "failed create team with input nil",
			fields: fields{
				teamRepo: func(ctrl *gomock.Controller) repository.TeamRepoItf {
					mockTeamRepo := mocks.NewMockTeamRepoItf(ctrl)
					return mockTeamRepo
				},
			},
			args: args{
				team: nil,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "failed create team with empty name",
			fields: fields{
				teamRepo: func(ctrl *gomock.Controller) repository.TeamRepoItf {
					mockTeamRepo := mocks.NewMockTeamRepoItf(ctrl)
					return mockTeamRepo
				},
			},
			args: args{
				team: &entity.Team{
					Name: "",
				},
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			tr := &teamUsecase{
				teamRepo: tt.fields.teamRepo(ctrl),
			}
			got, err := tr.CreateTeam(tt.args.team)
			if (err != nil) != tt.wantErr {
				t.Errorf("teamUsecase.CreateTeam() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("teamUsecase.CreateTeam() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_teamUsecase_AddPlayerToTeam(t *testing.T) {
	type fields struct {
		teamRepo func(ctrl *gomock.Controller) repository.TeamRepoItf
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
				teamRepo: func(ctrl *gomock.Controller) repository.TeamRepoItf {
					mockTeamRepo := mocks.NewMockTeamRepoItf(ctrl)
					mockTeamRepo.EXPECT().GetTeam(gomock.Any()).Return(&entity.Team{
						ID:   1,
						Name: "Manchester United",
					}, nil)
					mockTeamRepo.EXPECT().AddPlayerToTeam(gomock.Any(), gomock.Any()).Return(nil)
					return mockTeamRepo
				},
			},
			args: args{
				teamID: 1,
				player: &entity.Player{
					ID:   1,
					Name: "Cristiano Ronaldo",
				},
			},
			wantErr: false,
		},
		{
			name: "error add player",
			fields: fields{
				teamRepo: func(ctrl *gomock.Controller) repository.TeamRepoItf {
					mockTeamRepo := mocks.NewMockTeamRepoItf(ctrl)
					mockTeamRepo.EXPECT().GetTeam(gomock.Any()).Return(&entity.Team{
						ID:   1,
						Name: "Manchester United",
					}, nil)
					mockTeamRepo.EXPECT().AddPlayerToTeam(gomock.Any(), gomock.Any()).Return(errors.New("some error"))
					return mockTeamRepo
				},
			},
			args: args{
				teamID: 1,
				player: &entity.Player{
					Name: "Cristiano Ronaldo",
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			tr := &teamUsecase{
				teamRepo: tt.fields.teamRepo(ctrl),
			}
			if err := tr.AddPlayerToTeam(tt.args.teamID, tt.args.player); (err != nil) != tt.wantErr {
				t.Errorf("teamUsecase.AddPlayerToTeam() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_teamUsecase_GetPlayer(t *testing.T) {
	type fields struct {
		teamRepo func(ctrl *gomock.Controller) repository.TeamRepoItf
	}
	type args struct {
		teamID   int
		playerID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entity.Player
		wantErr bool
	}{
		{
			name: "success get player",
			fields: fields{
				teamRepo: func(ctrl *gomock.Controller) repository.TeamRepoItf {
					mockTeamRepo := mocks.NewMockTeamRepoItf(ctrl)
					mockTeamRepo.EXPECT().GetTeam(gomock.Any()).Return(&entity.Team{
						ID:   1,
						Name: "Manchester United",
						Players: []*entity.Player{
							{
								ID:   1,
								Name: "Cristiano Ronaldo",
							},
						},
					}, nil)
					return mockTeamRepo
				},
			},
			args: args{
				teamID:   1,
				playerID: 1,
			},
			want: &entity.Player{
				ID:   1,
				Name: "Cristiano Ronaldo",
			},
			wantErr: false,
		},
		{
			name: "error get team",
			fields: fields{
				teamRepo: func(ctrl *gomock.Controller) repository.TeamRepoItf {
					mockTeamRepo := mocks.NewMockTeamRepoItf(ctrl)
					mockTeamRepo.EXPECT().GetTeam(gomock.Any()).Return(nil, errors.New("some error"))
					return mockTeamRepo
				},
			},
			args: args{
				teamID:   1,
				playerID: 1,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "error player not found",
			fields: fields{
				teamRepo: func(ctrl *gomock.Controller) repository.TeamRepoItf {
					mockTeamRepo := mocks.NewMockTeamRepoItf(ctrl)
					mockTeamRepo.EXPECT().GetTeam(gomock.Any()).Return(&entity.Team{
						ID:   1,
						Name: "Manchester United",
						Players: []*entity.Player{
							{
								ID:   1,
								Name: "Cristiano Ronaldo",
							},
						},
					}, nil)
					return mockTeamRepo
				},
			},
			args: args{
				teamID:   1,
				playerID: 2,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			tr := &teamUsecase{
				teamRepo: tt.fields.teamRepo(ctrl),
			}
			got, err := tr.GetPlayer(tt.args.teamID, tt.args.playerID)
			if (err != nil) != tt.wantErr {
				t.Errorf("teamUsecase.GetPlayer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("teamUsecase.GetPlayer() = %v, want %v", got, tt.want)
			}
		})
	}
}
