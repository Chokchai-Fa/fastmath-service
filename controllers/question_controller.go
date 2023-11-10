package controllers

import (
	"fastmath/services"
	"fastmath/utils/handlers"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type QuestionController struct {
	questionService services.QuestionService
}

func NewQuestionController(allService services.AllService) (QuestionController, error) {

	return QuestionController{
		questionService: allService.QuestionService,
	}, nil
}

func (h *QuestionController) GetAllQuestion(c *gin.Context) {

	_, err := h.questionService.GetAllQuestion()

	if err != nil {
		handlers.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "OK",
		"version": os.Getenv("API_VERSION"),
		"now":     time.Now().Format("2006-01-02 15:04:05"),
	})

}

func (h *QuestionController) GetQuestion(c *gin.Context) {
	data := []string{"1", "2", "3"}
	handlers.HandleSuccess(c, data)
}
