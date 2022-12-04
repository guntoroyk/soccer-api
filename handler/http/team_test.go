package http

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/guntoroyk/soccer-api/entity"
	"github.com/guntoroyk/soccer-api/mocks"
	"github.com/guntoroyk/soccer-api/usecase"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestNewHandler(t *testing.T) {
	type args struct {
		teamUseCase usecase.TeamUsecaseItf
	}
	tests := []struct {
		name string
		args args
		want *handler
	}{
		{
			name: "success create new handler",
			args: args{
				teamUseCase: &mocks.MockTeamUsecaseItf{},
			},
			want: &handler{
				teamUseCase: &mocks.MockTeamUsecaseItf{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHandler(tt.args.teamUseCase); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

var (
	successGetTeamsJSON         = "{\"code\":200,\"data\":[{\"id\":1,\"name\":\"Manchester United\",\"players\":null}]}\n"
	successGetTeamJSON          = "{\"code\":200,\"data\":{\"id\":1,\"name\":\"Manchester United\",\"players\":null}}\n"
	failedGetTeamJSON           = "{\"code\":500,\"error\":\"failed get team\"}\n"
	failedGetTeamNotfoundJSON   = "{\"code\":404,\"error\":\"team not found\"}\n"
	failedCreateTeamJSON        = "{\"code\":500,\"error\":\"failed create team\"}\n"
	successCreateTeamJSON       = "{\"code\":201,\"data\":{\"id\":1,\"name\":\"Manchester United\",\"players\":null}}\n"
	failedAddPlayerToTeamJSON   = "{\"code\":500,\"error\":\"failed add player to team\"}\n"
	successGetPlayerJSON        = "{\"code\":200,\"data\":{\"id\":1,\"name\":\"Cristiano Ronaldo\"}}\n"
	failedGetPlayerJSON         = "{\"code\":500,\"error\":\"failed get player\"}\n"
	failedGetPlayerNotFoundJSON = "{\"code\":404,\"error\":\"player not found\"}\n"
)

func Test_handler_GetTeams(t *testing.T) {
	// Setup
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/teams")

	t.Run("success get teams", func(t *testing.T) {
		mockTeamUC := mocks.NewMockTeamUsecaseItf(ctrl)
		mockTeamUC.EXPECT().GetTeams().Return([]*entity.Team{
			{
				ID:   1,
				Name: "Manchester United",
			},
		}, nil)
		h := NewHandler(mockTeamUC)

		// Assertions
		if assert.NoError(t, h.GetTeams(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, successGetTeamsJSON, rec.Body.String())
		}
	})
}

func Test_handler_GetTeam(t *testing.T) {
	// Setup
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	e := echo.New()

	t.Run("success get team", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/api/teams/:id")

		mockTeamUC := mocks.NewMockTeamUsecaseItf(ctrl)
		mockTeamUC.EXPECT().GetTeam(gomock.Any()).Return(&entity.Team{
			ID:   1,
			Name: "Manchester United",
		}, nil)
		h := NewHandler(mockTeamUC)

		// Assertions
		if assert.NoError(t, h.GetTeam(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, successGetTeamJSON, rec.Body.String())
		}
	})

	t.Run("failed get team", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/api/teams/:id")

		mockTeamUC := mocks.NewMockTeamUsecaseItf(ctrl)
		mockTeamUC.EXPECT().GetTeam(gomock.Any()).Return(nil, errors.New("failed get team"))
		h := NewHandler(mockTeamUC)

		// Assertions
		if assert.NoError(t, h.GetTeam(c)) {
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
			assert.Equal(t, failedGetTeamJSON, rec.Body.String())
		}
	})

	t.Run("failed get team not found", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/api/teams/:id")

		mockTeamUC := mocks.NewMockTeamUsecaseItf(ctrl)
		mockTeamUC.EXPECT().GetTeam(gomock.Any()).Return(nil, entity.ErrTeamNotFound)
		h := NewHandler(mockTeamUC)

		// Assertions
		if assert.NoError(t, h.GetTeam(c)) {
			assert.Equal(t, http.StatusNotFound, rec.Code)
			assert.Equal(t, failedGetTeamNotfoundJSON, rec.Body.String())
		}
	})
}

func Test_handler_CreateTeam(t *testing.T) {
	// Setup
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	e := echo.New()

	t.Run("success create team", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"name":"Manchester United"}`))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/api/teams")

		mockTeamUC := mocks.NewMockTeamUsecaseItf(ctrl)
		mockTeamUC.EXPECT().CreateTeam(gomock.Any()).Return(&entity.Team{
			ID:   1,
			Name: "Manchester United",
		}, nil)
		h := NewHandler(mockTeamUC)

		// Assertions
		if assert.NoError(t, h.CreateTeam(c)) {
			assert.Equal(t, http.StatusCreated, rec.Code)
			assert.Equal(t, successCreateTeamJSON, rec.Body.String())
		}
	})

	t.Run("failed create team", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"name":"Manchester United"}`))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/api/teams")

		mockTeamUC := mocks.NewMockTeamUsecaseItf(ctrl)
		mockTeamUC.EXPECT().CreateTeam(gomock.Any()).Return(nil, errors.New("failed create team"))
		h := NewHandler(mockTeamUC)

		// Assertions
		if assert.NoError(t, h.CreateTeam(c)) {
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
			assert.Equal(t, failedCreateTeamJSON, rec.Body.String())
		}
	})

	t.Run("failed create team bad request", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"bad request}`))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/api/teams")

		mockTeamUC := mocks.NewMockTeamUsecaseItf(ctrl)
		h := NewHandler(mockTeamUC)

		// Assertions
		if assert.NoError(t, h.CreateTeam(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})
}

func Test_handler_AddPlayerToTeam(t *testing.T) {
	// Setup
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	e := echo.New()

	t.Run("success add player to team", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"name":"Cristiano Ronaldo"}`))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/api/teams/:id/players")

		mockTeamUC := mocks.NewMockTeamUsecaseItf(ctrl)
		mockTeamUC.EXPECT().AddPlayerToTeam(gomock.Any(), gomock.Any()).Return(nil)
		h := NewHandler(mockTeamUC)

		// Assertions
		if assert.NoError(t, h.AddPlayerToTeam(c)) {
			assert.Equal(t, http.StatusCreated, rec.Code)
		}
	})

	t.Run("failed add player to team", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"name":"Cristiano Ronaldo"}`))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/api/teams/:id/players")

		mockTeamUC := mocks.NewMockTeamUsecaseItf(ctrl)
		mockTeamUC.EXPECT().AddPlayerToTeam(gomock.Any(), gomock.Any()).Return(errors.New("failed add player to team"))
		h := NewHandler(mockTeamUC)

		// Assertions
		if assert.NoError(t, h.AddPlayerToTeam(c)) {
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
			assert.Equal(t, failedAddPlayerToTeamJSON, rec.Body.String())
		}
	})

	t.Run("failed add player to team bad request", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"bad request}`))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/api/teams/:id/players")

		mockTeamUC := mocks.NewMockTeamUsecaseItf(ctrl)
		h := NewHandler(mockTeamUC)

		// Assertions
		if assert.NoError(t, h.AddPlayerToTeam(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})
}

func Test_handler_GetPlayer(t *testing.T) {
	// Setup
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	e := echo.New()

	t.Run("success get player", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/api/players/:id")

		mockTeamUC := mocks.NewMockTeamUsecaseItf(ctrl)
		mockTeamUC.EXPECT().GetPlayer(gomock.Any(), gomock.Any()).Return(&entity.Player{
			ID:   1,
			Name: "Cristiano Ronaldo",
		}, nil)
		h := NewHandler(mockTeamUC)

		// Assertions
		if assert.NoError(t, h.GetPlayer(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, successGetPlayerJSON, rec.Body.String())
		}
	})

	t.Run("failed get player", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/api/players/:id")

		mockTeamUC := mocks.NewMockTeamUsecaseItf(ctrl)
		mockTeamUC.EXPECT().GetPlayer(gomock.Any(), gomock.Any()).Return(nil, errors.New("failed get player"))
		h := NewHandler(mockTeamUC)

		// Assertions
		if assert.NoError(t, h.GetPlayer(c)) {
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
			assert.Equal(t, failedGetPlayerJSON, rec.Body.String())
		}
	})

	t.Run("failed get player not found", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/api/players/:id")

		mockTeamUC := mocks.NewMockTeamUsecaseItf(ctrl)
		mockTeamUC.EXPECT().GetPlayer(gomock.Any(), gomock.Any()).Return(nil, entity.ErrPlayerNotFound)
		h := NewHandler(mockTeamUC)

		// Assertions
		if assert.NoError(t, h.GetPlayer(c)) {
			assert.Equal(t, http.StatusNotFound, rec.Code)
			assert.Equal(t, failedGetPlayerNotFoundJSON, rec.Body.String())
		}
	})
}
