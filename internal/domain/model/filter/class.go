package filter

type ClassFilter struct {
	Page         int    `query:"page"`
	PerPage      int    `query:"perPage"`
	StudentId    string `query:"studentId"`
	TeachersId   string `query:"teachersId"`
	SortBy       string `query:"sortBy"`
	SortType     string `query:"sortType"`
	ClassCode    string `query:"classCode"`
	ClassSubject string `query:"classSubject"`
}
