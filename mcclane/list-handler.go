package mcclane

import (
	"fmt"
	"github.com/ign/ipl-mcclane/mcclane/models"
	"labix.org/v2/mgo/bson"
	"net/http"
)

func ListHandler(w http.ResponseWriter, r *http.Request) {
	c, session := ConnectToCollection("brackets")
	defer session.Close()
	var result []*models.BracketElement
	err := c.Find(nil).Select(bson.M{"title": 1, "_id": 1, "sulg": 1}).All(&result)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, index := range result {
		fmt.Fprintln(w, index)
	}
}
