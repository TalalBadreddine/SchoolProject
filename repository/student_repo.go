package repository

import (
	"gorm.io/gorm"
	"server/entity"
	"server/filter"
	"strings"
)

type StudentRepository struct {
	db *gorm.DB
}

func ProvideStudentRepository(db *gorm.DB) StudentRepository {
	return StudentRepository{db: db}
}

func (s StudentRepository) SearchStudents(filter filter.StudentFilter) []*entity.Student {
	var db = s.db
	var students []*entity.Student

	if filter.PerPage == 0 {
		filter.PerPage = 5
	}

	offset := (filter.Page - 1) * filter.PerPage

	var classArray = strings.Split(filter.ClassesId, ",")
	var studentArray = strings.Split(filter.StudentsId, ",")

	db.Preload("Classes").
		Preload("Classes.Teachers").
		Preload("Classes.Students").
		Joins("LEFT JOIN students_classes on students_classes.student_id = students.id").
		Joins("LEFT JOIN classes on classes.id = students_classes.class_id").
		Scopes(entity.FilterByClassesId(classArray),
			entity.FilterByStudentsId(studentArray)).
		Offset(offset).
		Limit(filter.PerPage).
		Find(&students)

	return students
}
