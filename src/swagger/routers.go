/*
 * Data Anonymization Server
 *
 * This is a data anonymization server. You can set the anonymization requirements for the different datasets individually, and upload data to them. The uploaded data is anonymized on the server and can be then downloaded.
 *
 * API version: 0.1-alpha
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// NewRouter creates a new gorilla mux router
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		handler := logger(safetyNet(route.HandlerFunc), route.Name)
		router.Methods(route.Method).Path(route.Pattern).Handler(handler).Name(route.Name)
	}

	return router
}

var routes = []route{

	route{
		Name:        "Ping",
		Method:      strings.ToUpper("Get"),
		Pattern:     "/v1/ping",
		HandlerFunc: ping,
	},

	route{
		Name:        "DataNameGet",
		Method:      strings.ToUpper("Get"),
		Pattern:     "/v1/data/{name}",
		HandlerFunc: dataNameGet,
	},

	route{
		Name:        "DataNameDocumentIdGet",
		Method:      strings.ToUpper("Get"),
		Pattern:     "/v1/data/{name}/{documentId}",
		HandlerFunc: dataNameDocumentIDGet,
	},

	route{
		Name:        "AnonNameGet",
		Method:      strings.ToUpper("Get"),
		Pattern:     "/v1/anon/{name}",
		HandlerFunc: anonNameGet,
	},

	route{
		Name:        "AnonNameDocumentIdGet",
		Method:      strings.ToUpper("Get"),
		Pattern:     "/v1/anon/{name}/{documentId}",
		HandlerFunc: anonNameDocumentIDGet,
	},

	route{
		Name:        "DatasetsGet",
		Method:      strings.ToUpper("Get"),
		Pattern:     "/v1/datasets",
		HandlerFunc: datasetsGet,
	},

	route{
		Name:        "DatasetsNameDelete",
		Method:      strings.ToUpper("Delete"),
		Pattern:     "/v1/datasets/{name}",
		HandlerFunc: datasetsNameDelete,
	},

	route{
		Name:        "DatasetsNameGet",
		Method:      strings.ToUpper("Get"),
		Pattern:     "/v1/datasets/{name}",
		HandlerFunc: datasetsNameGet,
	},

	route{
		Name:        "DatasetsNamePut",
		Method:      strings.ToUpper("Put"),
		Pattern:     "/v1/datasets/{name}",
		HandlerFunc: datasetsNamePut,
	},

	route{
		Name:        "UploadPost",
		Method:      strings.ToUpper("Post"),
		Pattern:     "/v1/upload",
		HandlerFunc: uploadPost,
	},

	route{
		Name:        "UploadSessionIdPost",
		Method:      strings.ToUpper("Post"),
		Pattern:     "/v1/upload/{sessionId}",
		HandlerFunc: uploadSessionIDPost,
	},

	route{
		Name:        "UploadToEqulivalenceClassPost",
		Method:      strings.ToUpper("Post"),
		Pattern:     "/v1/upload/{sessionId}/{classId}",
		HandlerFunc: uploadDocumentToEqulivalenceClass,
	},

	route{
		Name:        "CreateEqulivalenceClass",
		Method:      strings.ToUpper("Post"),
		Pattern:     "/v1/classes",
		HandlerFunc: createEqulivalenceClass,
	},

	route{
		Name:        "ListClasses",
		Method:      strings.ToUpper("Get"),
		Pattern:     "/v1/classes",
		HandlerFunc: getAllEqulivalenceClasses,
	},

	route{
		Name:        "GetMatchingClasses",
		Method:      strings.ToUpper("Get"),
		Pattern:     "/v1/classes/matching",
		HandlerFunc: getMatchingEqulivalenceClasses,
	},

	route{
		Name:        "GetClass",
		Method:      strings.ToUpper("Get"),
		Pattern:     "/v1/classes/{id}",
		HandlerFunc: getEqulivalenceClassById,
	},

	route{
		Name:        "DeleteClass",
		Method:      strings.ToUpper("Delete"),
		Pattern:     "/v1/classes/{id}",
		HandlerFunc: deleteEqulivalenceClassById,
	},
}
