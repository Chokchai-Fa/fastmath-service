package controllers

import (
	"fastmath/services"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type QuestionHandler struct {
	questionService services.QuestionService
}

func NewQuestionHandler(allService services.AllService) (QuestionHandler, error) {

	return QuestionHandler{
		questionService: allService.QuestionService,
	}, nil
}

func (h *QuestionHandler) GetAllQuestion(c *gin.Context) {

	h.questionService.GetAllQuestion()

	c.JSON(http.StatusOK, gin.H{
		"status":  "OK",
		"version": os.Getenv("API_VERSION"),
		"now":     time.Now().Format("2006-01-02 15:04:05"),
	})

}
