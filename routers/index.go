package routers

import (
	"database/sql"
	"fastmath/utils/dbpostgres"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine, dbPostgres *dbpostgres.DBPG, dbPostgresCli *sql.DB) {
	createRouteCommon(router)
	group := router.Group("/fastmath")
	{
		initQuestionRoute(group, dbPostgres, dbPostgresCli)
	}
}

func createRouteCommon(router *gin.Engine) {
	// Health is Handle for health check service
	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  "OK",
			"version": os.Getenv("API_VERSION"),
			"now":     time.Now().Format("2006-01-02 15:04:05"),
		})
	})
}
