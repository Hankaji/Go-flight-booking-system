// Package middlewares
package middlewares

import (
	httperrors "flight-booking-server/http-errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler(ctx *gin.Context) {
	ctx.Next()

	if len(ctx.Errors) > 0 {
		err := ctx.Errors.Last().Err

		if httpErr, ok := httperrors.FromErr(err); ok {
			ctx.JSON(httpErr.StatusCode(), gin.H{
				"success": false,
				"message": httpErr.Error(),
			})
			return
		}

		// Default generic status
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
	}
}
