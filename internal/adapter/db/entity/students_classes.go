package entity

type StudentsClasses struct {
	StudentId uint `json:"studentId" gorm:"column:student_id;primaryKey"`
	ClassId   uint `json:"classId" gorm:"column:class_id;primaryKey"`
	Grade     int
}

func (s StudentsClasses) TableName() string {
	return "students_classes"
}
