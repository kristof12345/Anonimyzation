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
func getMatchingClasses(w http.ResponseWriter, r *http.Request) {

	var document anonmodel.Document
	if !tryReadRequestBody(r, &document, w) {
		return
	}

	anonbll.GetMatchingClasses(document)
}

// Create new equlivalence class
func createEqulivalenceClass(w http.ResponseWriter, r *http.Request) {
	var interval = make(map[string]anonmodel.Interval)
	interval["Age"] = anonmodel.Interval{10, 20}

	var categoric = make(map[string]string)
	categoric["City"] = "Bp"

	var class = anonmodel.EqulivalenceClass{1, categoric, interval, 0, true}

	err := anondb.CreateEqulivalenceClass(&class)

	if err != nil {
		logDBError(err)
		respondWithError(w, http.StatusInternalServerError, err.Error())
	} else {
		respondWithJSON(w, http.StatusOK, class)
	}
}

// Get all equlivalence classes
func getAllClasses(w http.ResponseWriter, r *http.Request) {
	classList, err := anondb.ListEqulivalenceClasses()

	anondb.DeleteEqulivalenceClass(1) //TODO: remove

	if err != nil {
		logDBError(err)
		respondWithError(w, http.StatusInternalServerError, err.Error())
	} else {
		respondWithJSON(w, http.StatusOK, classList)
	}
}

// Get equlivalence class by id
func getClassById(w http.ResponseWriter, r *http.Request) {
	var documents anonmodel.Documents
	if !tryReadRequestBody(r, &documents, w) {
		return
	}

	vars := mux.Vars(r)

	if id, err := strconv.Atoi(vars["id"]); err == nil {
		class := anondb.GetEqulivalenceClass(id)
		respondWithJSON(w, http.StatusOK, class)
	}
}
