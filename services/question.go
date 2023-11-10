package services

import "fastmath/repositories"

type QuestionService interface {
	GetAllQuestion()
}

type questionService struct {
	questionRepo repositories.QuestionRepository
}

func NewQuestionService(allRepo AllRepository) (QuestionService, error) {
	return &questionService{
		questionRepo: allRepo.QuestionRepository,
	}, nil
}

func (h *questionService) GetAllQuestion() {

	print("call service")

	h.questionRepo.GetAllQuestion()

}
