package service

import "Zhantasov/internal/models"

type Tests struct {
	data []models.Data
}

func NewTestMethod(data []models.Data) *Tests {
	return &Tests{data: data}
}

func (t *Tests) GetGrade() []int {
	var grades []int

	for _, grade := range t.data {
		grades = append(grades, grade.Grades)
	}
	return grades
}

func (t *Tests) GetQuestions(grade int) ([]models.Question, error) {
	var questions []models.Question
	for _, data := range t.data {
		if data.Grades == grade {
			questions = append(questions, data.Questions...)
		}
	}
	return questions, nil
}
