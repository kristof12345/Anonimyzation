package anonbll

import (
	"anondb"
	"anonmodel"
	"errors"
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
	if getErr == nil {
		class.Count++
		anondb.UpdateEqulivalenceClass(ecId, &class)
		if dataset.Settings.E <= class.Count {
			class.Active = false
			anondb.UpdateEqulivalenceClass(ecId, &class)
		}
	}

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
