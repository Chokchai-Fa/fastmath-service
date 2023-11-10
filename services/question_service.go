package services

import (
	"fastmath/repositories"
	"fastmath/utils/errors"
	"fastmath/utils/logs"
)

type QuestionService interface {
	GetAllQuestion() (*string, error)
}

type questionService struct {
	questionRepo repositories.QuestionRepository
}

func NewQuestionService(allRepo AllRepository) (QuestionService, error) {
	return &questionService{
		questionRepo: allRepo.QuestionRepository,
	}, nil
}

func (h *questionService) GetAllQuestion() (*string, error) {

	res := "1"
	print("call service")

	_, err := h.questionRepo.GetAllQuestion()
	if err != nil {
		logs.Error(err)
		return nil, errors.NewUnexpectedError()
	}

	return &res, nil

}
