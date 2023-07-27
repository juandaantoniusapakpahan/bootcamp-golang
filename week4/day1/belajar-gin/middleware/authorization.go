package middleware

import (
	"belajar-gin/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MiddlewareApp1() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if key := ctx.Request.Header.Get("APP1"); key != "APP1" {
			response := helper.ResponseTemplate{
				Code:   http.StatusUnauthorized,
				Status: "fail",
				Data:   "mohon masukan key dan value header dengan benar (APP1)",
			}
			ctx.AbortWithStatusJSON(response.Code, response)
			return
		}
		ctx.Next()
	}
}

func MiddlewareApp2() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if key := ctx.Request.Header.Get("APP2"); key != "APP2" {
			response := helper.ResponseTemplate{
				Code:   http.StatusUnauthorized,
				Status: "fail",
				Data:   "mohon masukan key dan value header dengan benar (APP2)",
			}
			ctx.AbortWithStatusJSON(response.Code, response)
			return
		}
		ctx.Next()
	}
}
