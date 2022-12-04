package entity

// Team is the struct for a team
type Team struct {
	ID      int       `json:"id"`
	Name    string    `json:"name"`
	Players []*Player `json:"players"`
}
