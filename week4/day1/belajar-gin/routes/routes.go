package routes

import (
	"belajar-gin/controller/controllerinterface"
	"belajar-gin/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TestServer struct {
	Name string `form:"name"`
}

func NewRoute(categoryController controllerinterface.CategoryControllerInterface) *gin.Engine {
	route := gin.Default()

	route.MaxMultipartMemory = 8 << 20 // 8 MiB

	route.StaticFS("/static", http.Dir("assets"))

	route.Use(middleware.SetCookies(), middleware.ErrorHandling())

	route.GET("/download", func(ctx *gin.Context) {
		// file, err := ioutil.ReadFile("./assets/Pertanyaan.txt")
		// if err != nil {
		// 	panic(err)
		// }
		// ctx.Header("Content-Disposition", "attachment;filename="+"Pertanyaan.txt")
		// ctx.Writer.Write([]byte(file))
		// ctx.JSON(http.StatusOK, gin.H{
		// 	"msg": "Download file successfully",
		// })

		ctx.FileAttachment("./assets/Pertanyaan.txt", "Pertanyaan.txt")
	})

	app1 := route.Group("/app1") // GROUP1
	app1.Use(middleware.MiddlewareApp1())
	{

		app1.GET(
			"/test", func(ctx *gin.Context) {
				ctx.Writer.Header().Add("Content-Type", "application/json")
				ctx.JSON(http.StatusOK, "COMPLETED")
			})

		app1.GET("/redirect", func(ctx *gin.Context) {
			ctx.Writer.Header().Add("Content-Type", "application/json")
			ctx.Redirect(http.StatusPermanentRedirect, "/app1/test")
		})

		app1.POST("/upload", func(ctx *gin.Context) {
			file, err := ctx.FormFile("myfile")
			if err != nil {
				panic(err)
			}

			ctx.SaveUploadedFile(file, "./assets/"+file.Filename)
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "upload file successfully",
			})
		})

		app1.GET("/servefile/:filename", func(ctx *gin.Context) {
			filename := ctx.Param("filename")

			filedir := "./assets/" + filename
			ctx.File(filedir)
		})
		app1.GET("/getcookies", func(ctx *gin.Context) {
			mycookies, err := ctx.Cookie("mycookies-gin")
			if err != nil {
				panic(err.Error())
			}
			ctx.JSON(http.StatusOK, gin.H{"mycookies": mycookies})
		})

		app1.GET("/queryparam", func(ctx *gin.Context) {
			tt := TestServer{}
			ctx.BindQuery(&tt)

			ctx.JSON(200, tt)
		})

	}

	app2 := route.Group("/app2") // GROUP2
	app2.Use(middleware.MiddlewareApp2())
	{
		app2.POST("/category", categoryController.Create)
		app2.GET("/category", categoryController.FindByName)
		app2.GET("/category/:categoryId", categoryController.FindById)
	}

	return route
}
