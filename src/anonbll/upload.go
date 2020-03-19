package anonbll

import (
	"anondb"
	"anonmodel"
)

// UploadDocuments validates the documents and inserts them into the database
func UploadDocuments(sessionID string, documents anonmodel.Documents, last bool, ecId int) (insertSuccessful bool, finalizeSuccessful bool, err error) {
	err = documents.Validate()
	if err != nil {
		return
	}

	dataset, err := anondb.SetUploadSessionBusy(sessionID)
	if err != nil {
		return
	}

	continuous := dataset.Settings.Mode == "continuous"
	err = uploadDocuments(documents, &dataset, continuous, last, ecId)
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

func uploadDocuments(documents anonmodel.Documents, dataset *anonmodel.Dataset, continuous bool, last bool, ecId int) error {
	if !last {
		defer anondb.SetUploadSessionNotBusy(dataset.Name, dataset.UploadSessionData.SessionID)
	}

	if dataset.Settings.Mode == "client-side" {
		for i := 1; i < len(documents); i++ {
			class, err := anondb.GetEqulivalenceClass(ecId)
			if err == nil {
				class.Count++
				anondb.UpdateEqulivalenceClass(ecId, &class)
				if dataset.Settings.E <= class.Count {
					class.Active = false
					anondb.UpdateEqulivalenceClass(ecId, &class)
				}
			}
		}
		return nil
	}

	return anondb.InsertDocuments(dataset.Name, documents, continuous)
}

func finalizeUpload(dataset *anonmodel.Dataset) error {
	defer anondb.FinishUploadSession(dataset.Name, dataset.UploadSessionData.SessionID)

	if dataset.Settings.Mode == "client-side" {
		return nil
	}

	continuous := dataset.Settings.Mode == "continuous"
	return anonymizeDataset(dataset, continuous)
}
