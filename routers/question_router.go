package routers

import (
	"database/sql"
	"fastmath/controllers"
	"fastmath/repositories"
	"fastmath/services"
	"fastmath/utils/dbpostgres"
	"log"

	"github.com/gin-gonic/gin"
)

func initQuestionRoute(group *gin.RouterGroup, dbPostgres *dbpostgres.DBPG, dbPostgresCli *sql.DB) {
	allRepo := services.AllRepository{
		QuestionRepository: repositories.NewQuestionRepositoryDB(dbPostgres, dbPostgresCli),
	}

	questionService, err := services.NewQuestionService(allRepo)
	if err != nil {
		log.Fatal("service error!:", err.Error())
	}

	allSerivce := services.AllService{
		QuestionService: questionService,
	}

	questionHandler, err := controllers.NewQuestionController(allSerivce)
	if err != nil {
		log.Fatal("Handler error!:", err.Error())
	}

	group.GET("/", questionHandler.GetAllQuestion)

	group.GET("/:id", questionHandler.GetQuestion)

}
