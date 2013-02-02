package models

import (
	"labix.org/v2/mgo/bson"
)

type BracketElement struct {
	Id    bson.ObjectId `bson:"_id" json:"id"`
	Title string        `json:"title"`
	Slug  string        `json:"slug"`
}
