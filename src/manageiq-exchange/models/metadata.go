package models

type Metadata struct {
	current_page int `json:"current_page"`
	total_pages  int `json:"total_pages"`
	total_count  int `json:"total_count"`
}
