package repository

import (
	"gorm.io/gorm"
	"server/entity"
	"server/filter"
	"strings"
)

type ClassRepository struct {
	db *gorm.DB
}

func ProvideClassRepository(db *gorm.DB) ClassRepository {
	return ClassRepository{db: db}
}

func (c ClassRepository) SearchClasses(filter filter.ClassFilter) []*entity.Class {
	var classes []*entity.Class

	if filter.PerPage == 0 {
		filter.PerPage = 5
	}

	offset := (filter.Page) * filter.PerPage

	var studentArray = strings.Split(filter.StudentId, ",")
	var teacherArray = strings.Split(filter.TeachersId, ",")
	var sortBy = filter.SortBy

	//TODO: SORTING NOT WORKING
	switch sortBy {

	case "teacherName":
		c.db.Order("Teachers.FirstName " + filter.SortType)

	case "className":
		c.db.Order("classes.Subject " + filter.SortType)

	case "studentName":
		c.db.Order("Students.FirstName " + filter.SortType)
	}

	c.db.Preload("Teachers").
		Preload("Students").
		Joins("LEFT JOIN students_classes on students_classes.class_id = classes.id").
		Joins("LEFT JOIN students on students.id = students_classes.student_id").
		Joins("LEFT JOIN teacher_classes on teacher_classes.class_id = classes.id").
		Joins("LEFT JOIN teachers on teachers.id = teacher_classes.teacher_id").
		Scopes(entity.FilterByStudentsId(studentArray),
			entity.FilterByTeachersId(teacherArray)).
		Group("classes.id").
		Offset(offset).
		Limit(filter.PerPage).
		Find(&classes)

	return classes
}
