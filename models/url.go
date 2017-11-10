package models

// URL - Model for the short_url table
type URL struct {
	ID      int    `json:"id"`
	Hash    string `json:"hash"`
	Address string `json:"address"`
	Clicks  int    `json:"clicks"`
}
