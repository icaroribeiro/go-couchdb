package models

type Document struct {
	Id  string `json:"_id,omitempty"`
	Rev   string `json:"_rev,omitempty"`
	Name string `json:"name"`
}
