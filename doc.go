package arangodb

// strings are pointers to allow null values
type Document struct {
	Id   string `json:"_id,omitempty"`
	Key  string `json:"_key,omitempty"`
	Rev  string `json:"_rev,omitempty"`
	To   string `json:"_to,omitempty"`
	From string `json:"_from,omitempty"`
}
