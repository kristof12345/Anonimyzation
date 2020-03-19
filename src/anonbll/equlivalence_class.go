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
