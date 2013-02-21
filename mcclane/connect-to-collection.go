package mcclane

import (
	"labix.org/v2/mgo"
)

func ConnectToCollection(col string) (*mgo.Collection, *mgo.Session) {
	session, err := mgo.Dial(config.DB_HOST)
	if err != nil {
		panic(err)
	}
	return session.DB(config.MONGO_DB).C(col), session
}
