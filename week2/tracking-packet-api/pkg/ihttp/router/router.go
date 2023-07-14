package router

import (
	"trancking-packet/middleware"
	"trancking-packet/pkg/ihttp/handler"

	"github.com/gin-gonic/gin"
)

func NewRoute(
	penerima handler.PenerimaHandlerInter,
	pengirim handler.PengirimHandlerInter,
	service handler.ServiceHandlerInter,
	lokasi handler.LokasiHandlerInter) *gin.Engine {

	r := gin.Default()
	r.Use(middleware.CORSMiddleware())
	r.Use(middleware.ErrorHandler())
	r.POST("/pengirim", pengirim.Create)
	r.POST("/penerima", penerima.Create)
	r.POST("/service", service.Create)
	r.POST("/lokasi", lokasi.Create)
	return r
}
