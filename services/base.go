package services

import (
	"fastmath/repositories"
)

type AllRepository struct {
	QuestionRepository repositories.QuestionRepository
}

type AllService struct {
	QuestionService QuestionService
}

type AllConfigFromDB struct {
}
