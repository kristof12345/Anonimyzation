package anonbll

import (
	"anondb"
	"anonmodel"
	"errors"
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

// Inserts them into the database, connecting it to the given equlivalence class
func UploadDocumentToEqulivalenceClass(sessionID string, document anonmodel.Document, ecId int) (bool, error) {

	dataset, sessionErr := anondb.SetUploadSessionBusy(sessionID)
	if sessionErr != nil {
		return false, sessionErr
	}

	defer anondb.SetUploadSessionNotBusy(dataset.Name, dataset.UploadSessionData.SessionID)

	if dataset.Settings.Algorithm != "client-side" {
		return false, errors.New("Algorithm should be client-side.")
	}

	class, getErr := anondb.GetEqulivalenceClass(ecId)
	if getErr != nil {
		return false, sessionErr
	}

	class.Count++
	if dataset.Settings.Max <= class.Count {
		class.Active = false
		// TODO: split class
	}
	anondb.UpdateEqulivalenceClass(ecId, &class)

	document["classId"] = ecId
	var documents = []anonmodel.Document{document}
	var insertErr = anondb.InsertDocuments(dataset.Name, documents, false)
	if insertErr != nil {
		return false, insertErr
	}

	return true, nil
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
