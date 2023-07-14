package handler

import (
	"database/sql"
	"net/http"
	errorhttp "trancking-packet/error"
	"trancking-packet/pkg/domain"
	"trancking-packet/pkg/helper"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type ServiceHandler struct {
	Service  domain.ServiceInterface
	DB       *sql.DB
	Validate *validator.Validate
}

type ServiceHandlerInter interface {
	Create(ctx *gin.Context)
}

func NewServiceHandler(service domain.ServiceInterface, db *sql.DB, validate *validator.Validate) ServiceHandlerInter {
	return &ServiceHandler{
		Service:  service,
		DB:       db,
		Validate: validate,
	}
}

func (h *ServiceHandler) Create(ctx *gin.Context) {
	bodyRequest := domain.ServiceRequest{}
	if err := helper.GetRequestBody(ctx.Request, &bodyRequest); err != nil {
		er := errorhttp.NewHttpError(err.Error(), http.StatusBadRequest)
		ctx.Error(er)
		return
	}

	err := h.Validate.Struct(bodyRequest)
	if err != nil {
		er := errorhttp.NewHttpError(err.Error(), http.StatusBadRequest)
		ctx.Error(er)
		return
	}

	id := uuid.New()
	serviceId := "service-" + id.String()

	serviceModel := domain.Service{
		Id:          serviceId,
		NamaService: bodyRequest.NamaService,
		HargaPerKg:  bodyRequest.HargaPerKg,
	}

	tx, err := h.DB.Begin()
	if err != nil {
		ctx.Error(err)
		return
	}
	defer helper.RollBackCommit(tx)

	result, err := h.Service.Add(ctx.Request.Context(), tx, serviceModel)
	if err != nil {
		ctx.Error(err)
		return
	}

	responseData := domain.ServiceResponse{
		Id:          result.Id,
		NamaService: result.NamaService,
		HargaPerKg:  result.HargaPerKg,
	}

	response := helper.Response{
		Code:   http.StatusCreated,
		Status: "success",
		Data:   M{"service": responseData},
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusCreated, response)

}

// migrate create -ext sql -dir db/migrations -seq create_songs_table
