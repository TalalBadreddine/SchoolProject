package dto

import "server/internal/domain/model"

type Statistics struct {
	ClassName     string `json:"class_name"`
	TeachersCount int    `json:"teachers_count"`
	StudentsCount int    `json:"students_count"`
}

func MapToStatistics(c *model.Class) *Statistics {
	if nil == c {
		return nil
	}

	return &Statistics{
		ClassName:     c.Subject + " " + c.Code,
		TeachersCount: len(c.Teachers),
		StudentsCount: len(c.StudentsClasses),
	}
}

func MapToStatisticsArray(c []*model.Class) []*Statistics {

	if nil == c {
		return nil
	}

	var statistics []*Statistics

	for _, stat := range c {
		statistics = append(statistics, MapToStatistics(stat))
	}

	return statistics
}
