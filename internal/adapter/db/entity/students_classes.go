package entity

type StudentsClasses struct {
	StudentId uint `json:"studentId"`
	Grade     int  `gorm:"foreignKey:Grade"`
}
