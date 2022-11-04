package repository

import (
	"gorm.io/gorm"
	"server/internal/adapter/db/entity"
	"server/internal/adapter/db/mapper"
	"server/internal/domain/model"
	"server/internal/domain/model/filter"
	"server/internal/domain/repository"
	"strings"
)

type ClassRepository struct {
	db *gorm.DB
}

func ProvideClassRepository(db *gorm.DB) repository.ClassRepository {
	return ClassRepository{db: db}
}

func (c ClassRepository) SearchClasses(filter filter.ClassFilter) []*model.Class {
	var classes []*entity.Class

	if filter.PerPage == 0 {
		filter.PerPage = 5
	}

	offset := (filter.Page) * filter.PerPage

	var studentArray = strings.Split(filter.StudentId, ",")
	var teacherArray = strings.Split(filter.TeachersId, ",")
	var sortBy = filter.SortBy

	switch sortBy {

	case "teacherName":
		c.db.Order("Teachers.FirstName " + filter.SortType)

	case "className":
		c.db = c.db.Order("classes.subject " + filter.SortType)

	case "studentName":
		c.db.Order("Students.FirstName " + filter.SortType)
	}

	c.db.Preload("Teachers").
		Preload("Students").
		Preload("StudentsClasses").
		Joins("LEFT JOIN students_classes on students_classes.class_id = classes.id").
		Joins("LEFT JOIN students on students.id = students_classes.student_id ").
		Joins("LEFT JOIN teacher_classes on teacher_classes.class_id = classes.id").
		Joins("LEFT JOIN teachers on teachers.id = teacher_classes.teacher_id").
		Scopes(entity.FilterByStudentsId(studentArray),
			entity.FilterByTeachersId(teacherArray)).
		Distinct().
		Order("students_classes.grade desc").
		Offset(offset).
		Limit(filter.PerPage).
		Find(&classes)

	return mapper.MapToClassModelArray(classes)
}
