package mongo

import (
	"os"
	"testing"

	"gopkg.in/mgo.v2"
)

func Test_CreateSession(t *testing.T) {
	defer os.Remove(ConfigFile)
	ConnectMongo()
}

func Test_GetEventualSession(t *testing.T) {
	session, database := GetEventualSession()

	if session.Mode() != mgo.Eventual {
		t.Errorf(test.ExpectedNumberButFoundNumber, mgo.Eventual, session.Mode())
	}

	if database.Name != "mongo" {
		t.Error(test.ExpectedStringButFoundString, "mongo", database.Name)
	}
}
