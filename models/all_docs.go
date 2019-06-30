package models

type AllDocuments struct {
	Offset  int64 `json:"offset"`
	Rows   []Item `json:"rows"`
	TotalRows int64 `json:"total_rows"`
}

type Item struct {
	Id  string `json:"id"`
	Key string `json:"key"`
	Value Rev `json:"value"`
}

type Rev struct {
	Rev string `json:"rev"`
}
