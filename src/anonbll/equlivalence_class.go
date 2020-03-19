package anonbll

import (
	"anondb"
	"anonmodel"
)

// Gets the matching equlivalent classes for the given attribute values
func GetMatchingClasses(document anonmodel.Document) ([]anonmodel.EqulivalenceClass, error) {

	var result = []anonmodel.EqulivalenceClass{}

	list, err := anondb.ListActiveEqulivalenceClasses()

	if err == nil {
		// Foreach equlivalence class
		for _, class := range list {
			if fieldsMatchEqulivalenceClass(class, document) {
				result = append(result, class)
			}
		}
	}

	return result, err
}

func fieldsMatchEqulivalenceClass(class anonmodel.EqulivalenceClass, document anonmodel.Document) bool {

	// Foreach categoric field
	for key, value := range class.CategoricAttributes {
		if document[key] != value {
			return false
		}
	}

	// Foreach interval field
	for key, value := range class.IntervalAttributes {
		interval, err := document[key].(anonmodel.NumericRange)
		if err == false && anonmodel.HasIntersection(value, interval) {
			return false
		}
	}

	return true
}

// Registers that a client wants to send a document to the given class, but the class does not contain k elements yet
func RegisterDocumentToClass(id int) {
	// TODO
	var k = 3
	var e1 = 1
	var e2 = 1

	class, err := anondb.GetEqulivalenceClass(id)
	if err == nil {
		class.Count++
		anondb.UpdateEqulivalenceClass(id, &class)
		if class.Count >= k+e1 {
			//TODO: kozponti tablaba kitenni
		}

		if class.Count >= k+e2 {
			//TODO: mar nem aktiv
		}
	}
}

// Inserts the document to the given equlivalence class
func AddDocumentToClass(document anonmodel.Document, id int) {
	//TODO: sava document linked to class
}
