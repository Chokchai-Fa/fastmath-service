package handlers

import (
	"fastmath/utils/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": data,
	})
}

func HandleError(c *gin.Context, err error) {
	switch e := err.(type) {
	case errors.AppError:
		c.JSON(e.Code, gin.H{
			"code":          e.Code,
			"error_message": e.Message,
		})
	case error:
		errMsg := "unexpected error"
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":          http.StatusInternalServerError,
			"error_message": errMsg,
		})
	}
}
