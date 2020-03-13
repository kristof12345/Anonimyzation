package swagger

import (
	"anonbll"
	"anondb"
	"anonmodel"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// Get equlivalence classes matching the given document
func getMatchingEqulivalenceClasses(w http.ResponseWriter, r *http.Request) {

	var document anonmodel.Document
	if !tryReadRequestBody(r, &document, w) {
		return
	}

	anonbll.GetMatchingClasses(document)
}

// Create new equlivalence class
func createEqulivalenceClass(w http.ResponseWriter, r *http.Request) {
	var request anonmodel.EqulivalenceClass
	if !tryReadRequestBody(r, &request, w) {
		return
	}

	result, err := anondb.CreateEqulivalenceClass(&request)

	if err != nil {
		logDBError(err)
		respondWithError(w, http.StatusInternalServerError, err.Error())
	} else {
		respondWithJSON(w, http.StatusOK, result)
	}
}

// Get all equlivalence classes
func getAllEqulivalenceClasses(w http.ResponseWriter, r *http.Request) {

	result, err := anondb.ListEqulivalenceClasses()

	if err != nil {
		logDBError(err)
		respondWithError(w, http.StatusInternalServerError, err.Error())
	} else {
		respondWithJSON(w, http.StatusOK, result)
	}
}

// Get equlivalence class by id
func getEqulivalenceClassById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	// Parsing id
	id, parseErr := strconv.Atoi(vars["id"])
	if parseErr != nil {
		respondWithError(w, http.StatusBadRequest, "Unable to parse id.")
		return
	}

	result, err := anondb.GetEqulivalenceClass(id)

	if err != nil {
		logDBError(err)
		respondWithError(w, http.StatusInternalServerError, err.Error())
	} else {
		respondWithJSON(w, http.StatusOK, result)
	}
}

// Delete equlivalence class by id
func deleteEqulivalenceClassById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	// Parsing id
	id, parseErr := strconv.Atoi(vars["id"])
	if parseErr != nil {
		respondWithError(w, http.StatusBadRequest, "Unable to parse id.")
		return
	}

	err := anondb.DeleteEqulivalenceClass(id)

	if err != nil {
		logDBError(err)
		respondWithError(w, http.StatusInternalServerError, err.Error())
	} else {
		respondWithJSON(w, http.StatusOK, "Deleted.")
	}
}

// Register document to equlivalence class
func registerDocumentToEqulivalenceClass(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	if id, err := strconv.Atoi(vars["id"]); err == nil {
		anonbll.RegisterDocumentToClass(id)
		respondWithJSON(w, http.StatusOK, "Registered.")
	}
}
