package entity

import "errors"

var (
	// ErrTeamCannotBeNil is the error message when team cannot be nil
	ErrTeamCannotBeNil = errors.New("team cannot be nil")

	// ErrTeamNameCannotBeEmpty is the error message when team name is required
	ErrTeamNameCannotBeEmpty = errors.New("team name is required")

	// ErrTeamNotFound is the error message when team is not found
	ErrTeamNotFound = errors.New("team not found")

	// ErrTeamIDIsRequired is the error message when team id is required
	ErrTeamIDIsRequired = errors.New("team id is required")

	// ErrPlayerCannotBeNil is the error message when player cannot be nil
	ErrPlayerCannotBeNil = errors.New("player cannot be nil")

	// ErrPlayerIDIsRequired is the error message when player id is required
	ErrPlayerIDIsRequired = errors.New("player id is required")

	// ErrPlayerNotFound is the error message when player is not found
	ErrPlayerNotFound = errors.New("player not found")

	// ErrPlayerAlreadyInTeam is the error message when player is already in team
	ErrPlayerAlreadyInTeam = errors.New("player already in team")

	// ErrPlayerNotInTeam is the error message when player is not in team
	ErrPlayerNotInTeam = errors.New("player not in team")

	// ErrPlayerNameIsRequired is the error message when player name is required
	ErrPlayerNameIsRequired = errors.New("player name is required")
)
