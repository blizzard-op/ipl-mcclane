package mcclane

import (
	"encoding/json"
	"github.com/ign/ipl-mcclane/brackets"
	"labix.org/v2/mgo/bson"
	"log"
)

func FindBracket(slug string) ([]byte, error) {
	var result *brackets.Bracket
	c, session := ConnectToCollection("brackets")
	defer session.Close()

	err := c.Find(bson.M{"slug": slug}).One(&result)
	if err != nil {
		log.Println(err)
		return make([]byte, 0), err
	}
	out, _ := json.Marshal(result)
	return out, nil
}
