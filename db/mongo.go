package db
import (
	"gopkg.in/mgo.v2"
	"log"
)

var (
	mongoSession *mgo.Session
)

func InitMongoDB() {
	session, err := mgo.Dial("mongodb://localhost/gopher")

	if err != nil {
		log.Fatal("Cannot connect to mongodb")
		return
	}
	session.SetMode(mgo.Monotonic, true)

	mongoSession = session
}

func Session() *mgo.Session {
	return mongoSession.Copy()
}

func WithCollection(collection string, s func(*mgo.Collection) error) error {
	session := Session()
	defer session.Close()
	c := session.DB("gopher").C(collection)
	return s(c)
}