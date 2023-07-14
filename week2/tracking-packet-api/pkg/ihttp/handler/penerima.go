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

type PenerimaHandler struct {
	Penerima domain.PenerimaInterfance
	DB       *sql.DB
	Validate *validator.Validate
}

type PenerimaHandlerInter interface {
	Create(ctx *gin.Context)
}

func NewPenerimaHandler(penerima domain.PenerimaInterfance, db *sql.DB, validate *validator.Validate) PenerimaHandlerInter {
	return &PenerimaHandler{
		Penerima: penerima,
		DB:       db,
		Validate: validate,
	}
}

func (h *PenerimaHandler) Create(ctx *gin.Context) {
	bodyRequest := domain.PenerimaRequest{}
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

	newUUID := uuid.New()
	id := "penerima-" + newUUID.String()

	penerima := domain.Penerima{
		Id:           id,
		NamaPenerima: bodyRequest.NamaPenerima,
		NoTelepon:    bodyRequest.NoTelepon,
	}

	tx, err := h.DB.Begin()
	if err != nil {
		ctx.Error(err)
		return
	}

	defer helper.RollBackCommit(tx)

	result, err := h.Penerima.Add(ctx.Request.Context(), tx, penerima)

	if err != nil {
		ctx.Error(err)
		return
	}

	responsePenerima := domain.PenerimaReponse{
		Id:           result.Id,
		NamaPenerima: result.NamaPenerima,
		NoTelepon:    result.NoTelepon,
	}

	response := helper.Response{
		Code:   http.StatusCreated,
		Status: "success",
		Data:   M{"penerima": responsePenerima},
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusCreated, response)
}
