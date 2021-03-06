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
		if document[key] == nil || document[key] != value {
			return false
		}
	}

	// Foreach interval field
	for key, value := range class.IntervalAttributes {
		if document[key] != nil {
			var interval = convertToRange(document[key])
			if !anonmodel.HasIntersection(value, interval) {
				return false
			}
		}
	}

	return true
}

func convertToRange(object interface{}) anonmodel.NumericRange {
	numericRange, err := object.(map[string]interface{})
	if err {
		return anonmodel.NumericRange{numericRange["Min"].(float64), numericRange["Max"].(float64)}
	} else {
		return anonmodel.NumericRange{object.(float64), object.(float64)}
	}
}
