// Package middlewares
package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler(ctx *gin.Context) {
	ctx.Next()

	if len(ctx.Errors) > 0 {
		err := ctx.Errors.Last().Err

		// Default generic status
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
	}
}
