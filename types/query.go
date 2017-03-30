package types

type Query struct {
	Aql       string `json:"query"`
	Count     bool   `json:"count"`
	BatchSize int    `json:"batch_size"`
}
