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

type LokasiHandler struct {
	Lokasi   domain.LokasiInterface
	DB       *sql.DB
	Validate *validator.Validate
}

type LokasiHandlerInter interface {
	Create(ctx *gin.Context)
}

func NewLokasiHandler(lokasi domain.LokasiInterface, db *sql.DB, validate *validator.Validate) LokasiHandlerInter {
	return &LokasiHandler{
		Lokasi:   lokasi,
		DB:       db,
		Validate: validate,
	}
}

func (l *LokasiHandler) Create(ctx *gin.Context) {
	bodyRequest := domain.LokasiRequest{}
	err := helper.GetRequestBody(ctx.Request, &bodyRequest)
	if err != nil {
		er := errorhttp.NewHttpError(err.Error(), http.StatusBadRequest)
		ctx.Error(er)
		return
	}

	err = l.Validate.Struct(bodyRequest)
	if err != nil {
		er := errorhttp.NewHttpError(err.Error(), http.StatusBadRequest)
		ctx.Error(er)
		return
	}

	id := uuid.New()
	newId := "lokasi-" + id.String()

	dataLokasi := domain.Lokasi{
		Id:         newId,
		NamaLokasi: bodyRequest.NamaLokasi,
		Alamat:     bodyRequest.Alamat,
	}

	tx, err := l.DB.Begin()
	if err != nil {
		ctx.Error(err)
		return
	}
	defer helper.RollBackCommit(tx)

	result, err := l.Lokasi.Add(ctx.Request.Context(), tx, dataLokasi)
	if err != nil {
		ctx.Error(err)
		return
	}

	dataResponse := domain.LokasiResponse{
		Id:         result.Id,
		NamaLokasi: result.NamaLokasi,
		Alamat:     result.Alamat,
	}

	response := helper.Response{
		Code:   http.StatusCreated,
		Status: "success",
		Data:   M{"lokasi": dataResponse},
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusCreated, response)
}
