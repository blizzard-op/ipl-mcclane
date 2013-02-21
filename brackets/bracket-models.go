package brackets

import (
	"labix.org/v2/mgo/bson"
)

type Bracket struct {
	Id        bson.ObjectId `bson:"_id" json:"id"`
	Users     []*User       `json:"userId"`
	Date      string        `json:"date"`
	Title     string        `json:"title"`
	Franchise string        `json:"franchise"`
	Slug      string        `json:"slug"`
	Kind      string        `json:"kind"`
	Matches   []*Match      `json:"matches"`
	Teams     []*Team       `json:"teams"`
	Groups    []*GroupStage `json:"groups"`
}

type Match struct {
	Id           string     `json:"id,omitempty"`
	Event        Event      `json:"event"`
	Children     []*int     `json:"children"`
	Parent       *int       `json:"parent"`
	LoserDropsTo *LoserDrop `json:"loserDropsTo"`
	Transform2d  Transform  `json:"transform2d"`
	HasLoserSlot bool       `json:"hasLoserSlot"`
}

type Transform struct {
	X        float32 `json:"x"`
	Y        float32 `json:"y"`
	PaddingX float32 `json:"paddingX"`
	PaddingY float32 `json:"paddingY"`
}

type Event struct {
	Id          *string       `json:"id,omitempty"`
	Title       string        `json:"title"`
	Starts_at   string        `json:"starts_at"`
	Ends_at     string        `json:"ends_at"`
	Stream      *Stream       `json:"stream"`
	Rebroadcast bool          `json:"rebroadcast"`
	Matchup     *Matchup      `json:"matchup"`
	Groups      []*MediaGroup `json:"groups"`
}

type MediaGroup struct {
	Id        string `json:"id"`
	Image_url string `json:"image_url"`
	Slug      string `json:"slug"`
	Name      string `json:"name"`
}

type Stream struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Matchup struct {
	Id      *string      `json:"id,omitempty"`
	Games   []*Game      `json:"games"`
	Teams   []*MatchTeam `json:"teams"`
	Best_of int          `json:"best_of"`
}

type Game struct {
	Id        *string `json:"id,omitempty"`
	Number    int     `json:"number"`
	Status    string  `json:"status"`
	Winner    *Team   `json:"winner,omitempty"`
	Starts_at *string `bson:"starts_at,omitempty" json:"starts_at,omitempty"`
	Ends_at   *string `bson:"ends_at,omitempty" json:"ends_at,omitempty"`
}

type Team struct {
	Id        *string `json:"id,omitempty"`
	Name      string  `json:"name"`
	Image_url string  `bson:"image_url,omitempty" json:"image_url,omitempty"`
	Seed      int     `bson:"seed,omitempty" json:"seed,omitempty"`
}

type LoserDrop struct {
	Match int `json:"match"`
	Slot  int `json:"slot"`
}

type MatchTeam struct {
	Id        *string `json:"id,omitempty"`
	Name      string  `json:"name"`
	Image_url string  `json:"image_url,omitempty"`
	Seed      int     `json:"seed"`
	Points    int     `json:"points"`
}

type GroupStage struct {
	Title       string       `json:"title"`
	Teams       []*MatchTeam `json:"teams"`
	Finished    bool         `json:"finished"`
	Matches     []*Match     `json:"matches"`
	AdvanceTo   []int        `json:"advanceTo"`
	Transform2d Transform    `json:"transform2d"`
}

type User struct {
	UserId string `json:"userId"`
}
