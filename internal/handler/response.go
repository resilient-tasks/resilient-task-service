package handler

import (
	"net/http"

	"github.com/fmarsico03/resilient-task-service/internal/httperror"
	"github.com/gin-gonic/gin"
)

func respondWithError(c *gin.Context, err error) {
	if httpErr, ok := err.(*httperror.HttpError); ok {
		c.AbortWithStatusJSON(httpErr.StatusCode, gin.H{"error": httpErr.Message})
	} else {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
	}
}
