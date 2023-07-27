package middleware

import (
	"belajar-gin/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandling() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			err := recover()
			if err != nil {
				response := helper.ResponseTemplate{
					Code:   http.StatusInternalServerError,
					Status: "error",
					Data:   err,
				}
				ctx.Writer.Header().Add("Content-Type", "application/json")
				ctx.AbortWithStatusJSON(response.Code, response)
				return
			}

		}()
		ctx.Next()
	}
}
