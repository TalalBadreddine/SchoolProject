package repository

type Filter struct {
	Class    string `json:"class"`
	Student  string `json:"student"`
	Teacher  string `json:"teacher"`
	Page     int    `json:"page"`
	PageSize int    `json:"pageSize"`
}
