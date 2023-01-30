package slingr

type EntityReference struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type RecordReference struct {
	Id     string          `json:"id"`
	Label  string          `json:"label"`
	Entity EntityReference `json:"entity"`
}
