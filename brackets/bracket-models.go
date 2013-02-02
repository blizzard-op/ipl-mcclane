package brackets

import (
	"labix.org/v2/mgo/bson"
)

type Bracket struct {
	Id      bson.ObjectId `bson:"_id" json:"id"`
	UserId  *string       `json:"userId"`
	Title   string        `json:"title"`
	Slug    string        `json:"slug"`
	Kind    string        `json:"kind"`
	Matches []*Match      `json:"matches"`
	Teams   []*Team       `json:"teams"`
	Groups  []*GroupStage `json:"groups"`
}

type Match struct {
	Id           int       `json:"id"`
	Event        Event     `json:"event"`
	Children     []*int    `json:"children"`
	Parent       *int      `json:"parent"`
	LoserDropsTo *int      `json:"loserDropsTo"`
	Transform2d  Transform `json:"transform2d"`
	HasLoserSlot bool      `json:"hasLoserSlot"`
}

type Transform struct {
	X        int `json:"x"`
	Y        int `json:"y"`
	PaddingX int `json:"paddingX"`
	PaddingY int `json:"paddingY"`
}

type Event struct {
	Title       string        `json:"title"`
	Starts_at   string        `json:"starts_at"`
	Ends_at     string        `json:"ends_at"`
	Stream      Stream        `json:"stream"`
	Rebroadcast bool          `json:"rebroadcast"`
	Matchup     *Matchup      `json:"matchup"`
	Groups      []*MediaGroup `json:"groups"`
}

type MediaGroup struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Stream struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Matchup struct {
	Games []*Game      `json:"games"`
	Teams []*MatchTeam `json:"teams"`
}

type Game struct {
	Number int    `json:"number"`
	Status string `json:"status"`
	Winner *Team  `json:"winner,omitempty"`
}

type Team struct {
	Id   *string `json:"id,omitempty"`
	Name string  `json:"name"`
	Seed int     `json:"seed"`
}

type MatchTeam struct {
	Id     *string `json:"id,omitempty"`
	Name   string  `json:"name"`
	Seed   int     `json:"seed"`
	Points int     `json:"points"`
}

type GroupStage struct {
	Title       string       `json:"title"`
	Teams       []*MatchTeam `json:"teams"`
	Finished    bool         `json:"finished"`
	Matches     []*Match     `json:"matches"`
	AdvanceTo   []int        `json:"advanceTo"`
	Transform2d Transform    `json:"transform2d"`
}
