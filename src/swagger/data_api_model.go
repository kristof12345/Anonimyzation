package swagger

import "anonmodel"

// ListDataResponse represents the JSON object sent by the server when listing data
type ListDataResponse struct {
	Result anonmodel.Documents `json:"result"`
	Next   string              `json:"next,omitempty"`
}
