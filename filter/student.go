package filter

type StudentFilter struct {
	Page       int    `query:"page"`
	PerPage    int    `query:"perPage"`
	ClassesId  string `query:"classesId"`
	StudentsId string `query:"studentsId"`
	SortBy     string `query:"sortBy"`
	SortType   string `query:"sortType"`
}
