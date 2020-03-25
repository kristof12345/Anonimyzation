package anondb

import (
	"anonmodel"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

// Creates a new central table item
func CreateCentralTableItem(item *anonmodel.CentralTableItem) (*anonmodel.CentralTableItem, error) {
	session := globalSession.Copy()
	defer session.Close()

	classes := session.DB("anondb").C("central")
	err := classes.Insert(item)
	if err != nil && mgo.IsDup(err) {
		return item, ErrDuplicate
	}

	return item, err
}

// Get an item by id
func GetCentralTableItem(id int) (item anonmodel.CentralTableItem, err error) {
	session := globalSession.Copy()
	defer session.Close()

	var filter = bson.M{"id": id}

	items := session.DB("anondb").C("central")
	err = items.Find(filter).One(&item)
	return
}
