package mcclane

import (
	"labix.org/v2/mgo"
)

const (
	//__________dev
	DB_HOST  = "localhost"
	MONGO_DB = "ipl-mcclane-dev"
	//__________prod
	// DB_HOST     = "media-cms-prd-mongodb-01.las1.colo.ignops.com:27017"
	// MONGO_DB    = "ipl-fantasy-prd"
	// CONTENT_API = "http://esports.ign.com/content/v1/events.json?startDate=2012-11-28T00:00:00-07:00"
)

func ConnectToCollection(col string) (*mgo.Collection, *mgo.Session) {
	session, err := mgo.Dial(DB_HOST)
	if err != nil {
		panic(err)
	}
	return session.DB(MONGO_DB).C(col), session
}
