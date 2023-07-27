package middleware

import "github.com/gin-gonic/gin"

func SetCookies() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.SetCookie("mycookies-gin", "inihanyauntukyangjomblo", 3600, "/", "localhost", false, true)
	}
}
