package repository

import (
	"server/entity"
	"server/filter"
	"server/storage"
	"strings"
)

func SearchClasses(filter filter.ClassFilter) []*entity.Class {
	var db = storage.GetDBInstance()
	var classes []*entity.Class

	if filter.PerPage == 0 {
		filter.PerPage = 5
	}

	offset := (filter.Page) * filter.PerPage

	var studentArray = strings.Split(filter.StudentId, ",")
	//var teacherArray = strings.Split(filter.Teacher, ",")

	db.Preload("Teachers").
		Preload("Students").
		Joins("LEFT JOIN students_classes on students_classes.class_id = classes.id").
		Joins("JOIN students on students.id = students_classes.student_id").
		Joins("LEFT JOIN teacher_classes on teacher_classes.class_id = classes.id").
		Joins("JOIN teachers on teachers.id = teacher_classes.teacher_id").
		Scopes(FilterByStudentsId(studentArray)).
		Offset(offset).
		Limit(filter.PerPage).
		Find(&classes)

	return classes
}
