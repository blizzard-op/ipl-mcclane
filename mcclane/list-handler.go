package mcclane

import (
	"encoding/json"
	"fmt"
	"github.com/ign/ipl-mcclane/mcclane/models"
	"labix.org/v2/mgo/bson"
	"net/http"
)

func ListHandler(w http.ResponseWriter, r *http.Request) {
	c, session := ConnectToCollection("brackets")
	defer session.Close()
	SetCORHeaders(w, r)
	var result []*models.BracketElement
	err := c.Find(nil).Select(bson.M{"title": 1, "_id": 1, "slug": 1, "date": 1}).All(&result)
	if err != nil {
		fmt.Println(err)
		return
	}
	out, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintln(w, string(out))
}
