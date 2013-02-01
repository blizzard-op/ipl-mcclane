package mcclane

import (
	"github.com/ign/ipl-mcclane/brackets"
	"labix.org/v2/mgo/bson"
	"log"
)

func CreateBracket(bracket *brackets.Bracket) error {
	c, session := ConnectToCollection("brackets")
	defer session.Close()

	var newId bson.ObjectId
	if !bracket.Id.Valid() {
		newId = bson.NewObjectId()
		bracket.Id = newId
	} else {
		newId = bracket.Id
	}

	err := c.Insert(bracket)
	if err != nil {
		log.Println(err)
	}

	return nil
}
