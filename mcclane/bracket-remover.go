package mcclane

import (
	"github.com/ign/ipl-mcclane/brackets"
	"labix.org/v2/mgo/bson"
	"log"
	"net/http"
)

func RemoveBracket(slug string) error {
	res, _ := FindBracket(slug)
	result, err := brackets.Parse(res)
	if err != nil {
		return err
	}
	client := new(http.Client)
	for _, match := range result.Matches {
		if match.Event.Id != nil {
			deleteRequest, err := http.NewRequest("DELETE", "http://esports.ign.com/content/v2/events/"+string(*match.Event.Id), nil)
			if err != nil {
				log.Println(err)
			}
			resp, err := client.Do(deleteRequest)
			if err != nil {
				log.Println(err)
			}
			defer resp.Body.Close()
			log.Printf("deleted %s", *match.Event.Id)
		}
	}

	c, session := ConnectToCollection("brackets")
	defer session.Close()

	err = c.Remove(bson.M{"slug": slug})
	if err != nil {
		log.Println(err)
	}
	return nil
}
