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

type StudentRepository struct {
	db *gorm.DB
}

func ProvideStudentRepository(db *gorm.DB) repository.StudentRepository {
	return StudentRepository{db: db}
}

func (s StudentRepository) SearchStudents(filter filter.StudentFilter) []*model.Student {
	var db = s.db
	var students []*entity.Student

	if filter.PerPage == 0 {
		filter.PerPage = 5
	}

	offset := (filter.Page - 1) * filter.PerPage

	var classArray = strings.Split(filter.ClassesId, ",")
	var studentArray = strings.Split(filter.StudentsId, ",")
	var sortBy = filter.SortBy

	orderByMap := make(map[string]string)

	orderByMap["studentName"] = "students.first_name "
	orderByMap["teacherName"] = "teachers.first_name "
	orderByMap["className"] = "classes.subject "

	db.Preload("Classes").
		Preload("Classes.Teachers").
		Preload("Classes.Students").
		Joins("FULL OUTER JOIN students_classes on students_classes.student_id = students.id").
		Joins("FULL OUTER JOIN classes on classes.id = students_classes.class_id").
		Scopes(entity.FilterByClassesId(classArray),
			entity.FilterByStudentsId(studentArray)).
		Offset(offset).
		Limit(filter.PerPage).
		Order(orderByMap[sortBy] + filter.SortType).
		Find(&students)

	return mapper.MapToStudentModelArray(students)
}
