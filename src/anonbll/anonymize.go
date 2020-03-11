package anonbll

import (
	"anondb"
	"anonmodel"
	"fmt"
	"log"
	"time"
)

type anonymizerAlgorithm interface {
	initialize(*anonmodel.Dataset, string, []anonmodel.FieldAnonymizationInfo)
	anonymize() error
}

func anonymizeDataset(dataset *anonmodel.Dataset, continuous bool) error {
	start := time.Now()
	defer func() { log.Printf("Anonymization took %v", time.Since(start)) }()

	if err := doAnonymization(dataset, continuous); err != nil {
		return err
	}

	if !continuous {
		return nil
	}
	return anondb.MoveTempAnonymizedData(dataset.Name)
}

func doAnonymization(dataset *anonmodel.Dataset, continuous bool) error {
	var anonCollectionName string
	if continuous {
		anonCollectionName = "temp_anon_" + dataset.Name
	} else {
		anonCollectionName = "anon_" + dataset.Name
	}

	if err := anondb.CopyData(dataset.Name, continuous, anonCollectionName); err != nil {
		return err
	}

	fieldsToSuppress := anonmodel.GetSuppressedFields(dataset.Fields)
	if err := anondb.SuppressFields(anonCollectionName, fieldsToSuppress); err != nil {
		return err
	}

	var algorithm anonymizerAlgorithm
	quasiIdentifierFields := anonmodel.GetQuasiIdentifierFields(dataset.Fields)

	if dataset.Settings.Algorithm == "mondrian" {
		algorithm = &mondrian{}
	} else {
		return fmt.Errorf("The only currently supported anonymization algorithm is 'mondrian', got '%v'", dataset.Settings.Algorithm)
	}

	algorithm.initialize(dataset, anonCollectionName, quasiIdentifierFields)
	if err := algorithm.anonymize(); err != nil {
		return err
	}

	return anondb.RenameAnonFields(anonCollectionName, quasiIdentifierFields)
}

// Gets a matching equlivalent class for the given values
func GetMatchingClasses(document anonmodel.Document) []anonmodel.EqulivalenceClass {

	var result []anonmodel.EqulivalenceClass

	list, err := anondb.ListEqulivalenceClasses()

	if err == nil {
		// Foreach equlivalence class
		for _, class := range list {
			if fieldsMatchEqulivalenceClass(class, document) {
				result = append(result, class)
			}
		}
	}

	return result
}

// Inserts dthe document and increases class count
func AddDocumentToClass(document anonmodel.Document, class anonmodel.EqulivalenceClass) {

	class = anondb.GetEqulivalenceClass(class.Id)
	class.Count++
	anondb.UpdateEqulivalenceClass(class.Id, &class)
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
		var interval = document[key].(anonmodel.Interval)
		if interval.BottomLimit > value.UpperLimit || interval.UpperLimit < value.BottomLimit {
			return false
		}
	}

	return true
}
