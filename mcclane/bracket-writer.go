package mcclane

import (
	"github.com/ign/ipl-mcclane/brackets"
	"labix.org/v2/mgo/bson"
	"log"
)

func UpdateBracket(bracket *brackets.Bracket) error {
	c, session := ConnectToCollection("brackets")
	defer session.Close()
	err := c.UpdateId(bracket.Id, bracket)
	if err != nil {
		log.Println("Error with updating ")
		return err
	}
	// log.Println("Updated ")
	return nil
}

func InsertBracket(bracket *brackets.Bracket) error {
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
