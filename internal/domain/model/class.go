package model

type Class struct {
	Id              uint               `json:"id"`
	Subject         string             `json:"subject"`
	Code            string             `json:"code"`
	Teachers        []*Teacher         `json:"teachers"`
	Students        []*Student         `json:"students"`
	StudentsClasses []*StudentsClasses `json:"studentsClasses"`
}
