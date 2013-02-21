package mcclane

import (
	"labix.org/v2/mgo"
)

const (
	//__________dev
	// DB_HOST  = "localhost"
	// MONGO_DB = "ipl-mcclane-dev"
	//__________prod
	DB_HOST  = "media-cms-prd-mongodb-01.las1.colo.ignops.com:27017"
	MONGO_DB = "ipl-mcclane-prd"
)

func ConnectToCollection(col string) (*mgo.Collection, *mgo.Session) {
	session, err := mgo.Dial(DB_HOST)
	if err != nil {
		panic(err)
	}
	return session.DB(MONGO_DB).C(col), session
}
