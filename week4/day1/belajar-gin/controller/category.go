package controller

import (
	"belajar-gin/controller/controllerinterface"
	"belajar-gin/domain"
	"belajar-gin/helper"
	"belajar-gin/service/serviceinterface"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryControllerImplement struct {
	CategoryService serviceinterface.CategoryServiceInterface
}

func NewCategoryControllerImplement(service serviceinterface.CategoryServiceInterface) controllerinterface.CategoryControllerInterface {
	return &CategoryControllerImplement{CategoryService: service}
}

func (cc *CategoryControllerImplement) Create(ctx *gin.Context) {
	requestBody := domain.AddCategory{}
	// ctx.Bind(&requestBody)
	helper.GetResponseBody(ctx.Request, &requestBody)

	result := cc.CategoryService.Create(requestBody)
	response := helper.ResponseTemplate{
		Code:   http.StatusCreated,
		Status: "success",
		Data:   map[string]interface{}{"category": result},
	}

	ctx.Writer.Header().Add("Content-Type", "application/json")
	ctx.JSON(http.StatusCreated, response)
}

func (cc *CategoryControllerImplement) FindByName(ctx *gin.Context) {
	categoryname := ctx.Query("categoryName")

	result := cc.CategoryService.FindByName(categoryname)
	response := helper.ResponseTemplate{
		Code:   http.StatusOK,
		Status: "success",
		Data:   map[string]interface{}{"category": result},
	}

	ctx.Writer.Header().Add("Content-Type", "application/json")
	ctx.JSON(http.StatusCreated, response)
}

func (cc *CategoryControllerImplement) FindById(ctx *gin.Context) {
	categoryId, _ := strconv.Atoi(ctx.Param("categoryId"))
	result := cc.CategoryService.FindById(categoryId)
	response := helper.ResponseTemplate{
		Code:   http.StatusOK,
		Status: "success",
		Data:   map[string]interface{}{"category": result},
	}

	ctx.Writer.Header().Add("Content-Type", "application/json")
	ctx.JSON(http.StatusCreated, response)
}
