package routes

import (
	"server/controller"

	"github.com/labstack/echo/v4"
)

func InitTeachersRoutes(e *echo.Echo) {
	e.GET("Teacher", controller.GetTeachers)
	e.POST("Teacher", controller.AddTeacher)
	e.PUT("Teacher/Class", controller.AddTeacherToClass)
	e.GET("Teacher/:id/student", controller.GetStudentsByTeacherId)
	// e.GET("Teacher/:id/class", controller.GetTeacherByClassId)
}

// Join between teacher class teacherClasses, where (teacher.id = teacherClasses.teacher_id)
