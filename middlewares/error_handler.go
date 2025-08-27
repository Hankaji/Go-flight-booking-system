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
				"code":    httpErr.StatusCode(),
				"message": httpErr.AppErr(),
				"error":   httpErr.Err(),
			})
			return
		}
		t1, t2 := httperrors.FromErr(err)
		println("DEBUGGING: ", t1, t2)

		// Default generic status
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
	}
}
