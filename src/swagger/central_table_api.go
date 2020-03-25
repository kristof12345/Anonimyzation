package swagger

import (
	"anondb"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// Returns Ok if the central table contains the given Ec
func centralTableContainsClass(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	// Parsing id
	id, parseErr := strconv.Atoi(vars["id"])
	if parseErr != nil {
		respondWithError(w, http.StatusBadRequest, "Unable to parse id.")
		return
	}

	item, err := anondb.GetCentralTableItem(id)

	if err != nil {
		respondWithJSON(w, http.StatusNotFound, "Not found.")
	} else {
		respondWithJSON(w, http.StatusOK, item)
	}
}
