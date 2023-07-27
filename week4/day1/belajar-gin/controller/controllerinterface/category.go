package controllerinterface

import "github.com/gin-gonic/gin"

type CategoryControllerInterface interface {
	Create(ctx *gin.Context)
	FindByName(ctx *gin.Context)
	FindById(ctx *gin.Context)
}
