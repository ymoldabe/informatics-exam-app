package service

import (
	configs "Zhantasov/config"
	"Zhantasov/internal/models"
)

type Testing interface {
	GetGrade() []int
	GetQuestions(grades int) ([]models.Question, error)
}

type Service struct {
	Config configs.Config
	Data   []models.Data
	Testing
}

func New(cnf configs.Config, data []models.Data) *Service {
	return &Service{
		Config:  cnf,
		Data:    data,
		Testing: NewTestMethod(data),
	}
}
