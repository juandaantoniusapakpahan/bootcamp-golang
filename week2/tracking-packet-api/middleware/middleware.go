package middleware

import (
	"net/http"
	errorhttp "trancking-packet/error"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
		for _, err := range ctx.Errors {
			switch e := err.Err.(type) {
			case errorhttp.Http:
				ctx.AbortWithStatusJSON(e.StatusCode, e)
			default:
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"message": "maaf, terjadi kesalah pada server", "description": e.Error()})
			}
		}
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, DELETE, GET, PUT")
	}
}
