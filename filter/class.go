package filter

type ClassFilter struct {
	Page       int    `query:"page"`
	PerPage    int    `query:"perPage"`
	StudentId  string `query:"studentId"`
	TeachersId string `query:"teachersId""`
}
