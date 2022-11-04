package dto

import "server/internal/domain/model"

type StudentsClasses struct {
	ClassId uint `json:"classId" gorm:"column:class_id;primaryKey"`
	Grade   int  `gorm:"foreignKey:Grade"`
}

func MapToStudentsClassesDto(s *model.StudentsClasses) *StudentsClasses {
	if nil == s {
		return nil
	}

	return &StudentsClasses{
		ClassId: s.ClassId,
		Grade:   s.Grade,
	}
}

//func MapToStudentsClassesDtoArray(s []*model.StudentsClasses) []*StudentsClasses {
//
//	if s == nil {
//		return nil
//	}
//
//	var studentsClasses []*StudentsClasses
//
//	for _, studentClass := range studentsClasses {
//		studentsClasses = append(studentsClasses, MapToStudentsClassesDto(studentClass))
//	}
//
//	return studentsClasses
//}
