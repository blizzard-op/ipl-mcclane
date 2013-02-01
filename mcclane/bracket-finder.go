package mcclane

import (
	"encoding/json"
	"github.com/ign/ipl-mcclane/brackets"
	"labix.org/v2/mgo/bson"
	"log"
)

func FindBracket() {
	var result *brackets.Bracket
	c, session := ConnectToCollection("brackets")
	defer session.Close()

	err := c.Find(bson.M{}).One(&result)
	if err != nil {
		log.Println(err)
	}
	out, _ := json.Marshal(result)
	log.Println(string(out))

}
