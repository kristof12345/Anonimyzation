package anonbll

import (
	"anondb"
	"anonmodel"
	"time"
)

// UploadDocuments validates the documents and inserts them into the database
func UploadDocuments(sessionID string, documents anonmodel.Documents, last bool) (insertSuccessful bool, finalizeSuccessful bool, err error) {
	err = documents.Validate()
	if err != nil {
		return
	}

	dataset, err := anondb.SetUploadSessionBusy(sessionID)
	if err != nil {
		return
	}

	continuous := dataset.Settings.Mode == "continuous"
	err = uploadDocuments(documents, &dataset, continuous, last)
	if err != nil {
		return
	}
	insertSuccessful = true

	if last {
		err = finalizeUpload(&dataset)
		if err == nil {
			finalizeSuccessful = true
		}
	}
	return
}

// Registers the upload intent, if K intents were registered the EC is added to the central table
func RegisterUploadIntent(datasetName string, classId int) bool {

	var dataset, _ = anondb.GetDataset(datasetName)
	var class, _ = anondb.GetEqulivalenceClass(classId)

	class.IntentCount++
	if dataset.Settings.K+dataset.Settings.E == class.IntentCount { // Waits for K + E intents before puting to central table
		var item = anonmodel.CentralTableItem{classId, time.Now().AddDate(0, 0, 1)} //Add one day
		anondb.CreateCentralTableItem(&item)
		anondb.UpdateEqulivalenceClass(classId, &class)
		return true
	}
	anondb.UpdateEqulivalenceClass(classId, &class)
	return false
}

// Inserts documents into the database, connecting it to the given equlivalence class
func UploadDocumentToEqulivalenceClass(sessionID string, document anonmodel.Document, ecId int) (bool, string) {

	dataset, sessionErr := anondb.SetUploadSessionBusy(sessionID)
	if sessionErr != nil {
		return false, "Dataset not found"
	}

	if dataset.Settings.Algorithm != "client-side" {
		return false, "Algorithm should be client-side"
	}

	class, getErr := anondb.GetEqulivalenceClass(ecId)
	if getErr != nil {
		return false, "Equlivalence class not found"
	}

	class.Count++
	if dataset.Settings.Max <= class.Count {
		class.Active = false
		// Split class
		splitEqulivalenceClass(&class)
	}
	anondb.UpdateEqulivalenceClass(ecId, &class)

	document["classId"] = ecId

	var documents = []anonmodel.Document{document}

	// Insert to DB
	var insertErr = anondb.InsertDocuments(dataset.Name, documents, false)
	if insertErr != nil {
		return false, "Unable to insert documents"
	}

	anondb.SetUploadSessionNotBusy(dataset.Name, sessionID)
	anondb.FinishUploadSession(dataset.Name, sessionID)

	return true, "Success!"
}

func splitEqulivalenceClass(class *anonmodel.EqulivalenceClass) {
	var lowerInterval map[string]anonmodel.NumericRange
	var upperInterval map[string]anonmodel.NumericRange
	// Foreach
	for key, value := range class.IntervalAttributes {
		half := value.Max / 2
		lowerInterval[key] = anonmodel.NumericRange{Min: value.Min, Max: half}
		upperInterval[key] = anonmodel.NumericRange{Min: half, Max: value.Max}
	}
	var lowerClass = anonmodel.EqulivalenceClass{IntervalAttributes: lowerInterval, CategoricAttributes: class.CategoricAttributes}
	var upperClass = anonmodel.EqulivalenceClass{IntervalAttributes: upperInterval, CategoricAttributes: class.CategoricAttributes}

	anondb.CreateEqulivalenceClass(&lowerClass)
	anondb.CreateEqulivalenceClass(&upperClass)
}

func uploadDocuments(documents anonmodel.Documents, dataset *anonmodel.Dataset, continuous bool, last bool) error {
	if !last {
		defer anondb.SetUploadSessionNotBusy(dataset.Name, dataset.UploadSessionData.SessionID)
	}

	return anondb.InsertDocuments(dataset.Name, documents, continuous)
}

func finalizeUpload(dataset *anonmodel.Dataset) error {
	defer anondb.FinishUploadSession(dataset.Name, dataset.UploadSessionData.SessionID)

	continuous := dataset.Settings.Mode == "continuous"
	return anonymizeDataset(dataset, continuous)
}
