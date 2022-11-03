package model

type Teacher struct {
	Id        uint     `json:"id"`
	FirstName string   `json:"firstName"`
	LastName  string   `json:"lastName"`
	Classes   []*Class `json:"classes"`
}
