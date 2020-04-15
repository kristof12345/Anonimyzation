package anondb

import (
	"anonmodel"
	"github.com/globalsign/mgo/bson"
)

// Creates a new equlivalence class
func CreateEqulivalenceClass(class *anonmodel.EqulivalenceClass) (*anonmodel.EqulivalenceClass, error) {
	session := globalSession.Copy()
	defer session.Close()

	//Check if id is already used
	var filter = bson.M{"id": class.Id}
	eq := session.DB("anondb").C("classes")
	count, err := eq.Find(filter).Count()
	if count > 0 {
		return nil, ErrDuplicate
	}

	// Insert document
	classes := session.DB("anondb").C("classes")
	err = classes.Insert(class)
	return class, err
}

// Updates an equlivalence class
func UpdateEqulivalenceClass(id int, class *anonmodel.EqulivalenceClass) error {
	session := globalSession.Copy()
	defer session.Close()

	var filter = bson.M{"id": id}

	classes := session.DB("anondb").C("classes")
	err := classes.Update(filter, class)
	return err
}

// Deletes an equlivalence class
func DeleteEqulivalenceClass(id int) error {
	session := globalSession.Copy()
	defer session.Close()

	var filter = bson.M{"id": id}

	classes := session.DB("anondb").C("classes")
	_, err := classes.RemoveAll(filter)
	return err
}

// Get an equlivalence class by id
func GetEqulivalenceClass(id int) (class anonmodel.EqulivalenceClass, err error) {
	session := globalSession.Copy()
	defer session.Close()

	var filter = bson.M{"id": id}

	classes := session.DB("anondb").C("classes")
	err = classes.Find(filter).One(&class)
	return
}

// Lists all the equlivalence classes in the database
func ListEqulivalenceClasses() (classList []anonmodel.EqulivalenceClass, err error) {
	session := globalSession.Copy()
	defer session.Close()

	classes := session.DB("anondb").C("classes")
	if err = classes.Find(nil).All(&classList); err != nil {
		return
	}

	if classList == nil {
		classList = []anonmodel.EqulivalenceClass{}
	}
	return
}

// Lists the active equlivalence classes in the database
func ListActiveEqulivalenceClasses() (classList []anonmodel.EqulivalenceClass, err error) {
	session := globalSession.Copy()
	defer session.Close()

	var filter = bson.M{"active": true}

	classes := session.DB("anondb").C("classes")
	if err = classes.Find(filter).All(&classList); err != nil {
		return
	}

	if classList == nil {
		classList = []anonmodel.EqulivalenceClass{}
	}
	return
}
