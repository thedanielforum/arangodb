package types

import "encoding/json"

type Result struct {
	Result  []json.RawMessage `json:"result"`
	Count   int               `json:"count"`
	Cached  bool              `json:"cached"`
	HasMore bool              `json:"hasMore"`
	Error   bool              `json:"error"`
	Code    int               `json:"code"`
}

type Results struct {
	Result  json.RawMessage `json:"result"`
	Count   int             `json:"count"`
	Cached  bool            `json:"cached"`
	HasMore bool            `json:"hasMore"`
	Error   bool            `json:"error"`
	Code    int             `json:"code"`
}
