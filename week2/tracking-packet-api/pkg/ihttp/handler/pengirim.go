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

type M map[string]interface{}

type PengirimHandler struct {
	Pengirim domain.PengirimInterfance
	DB       *sql.DB
	Validate *validator.Validate
}

type PengirimHandlerInter interface {
	Create(ctx *gin.Context)
}

func NewPengirimHandler(pengirim domain.PengirimInterfance, db *sql.DB, validate *validator.Validate) PengirimHandlerInter {
	return &PengirimHandler{Pengirim: pengirim, DB: db, Validate: validate}
}

func (ph *PengirimHandler) Create(ctx *gin.Context) {
	bodyRequest := domain.PengirimRequest{}
	if err := helper.GetRequestBody(ctx.Request, &bodyRequest); err != nil {
		er := errorhttp.NewHttpError(err.Error(), http.StatusBadRequest)
		ctx.Error(er)
		return
	}

	err := ph.Validate.Struct(bodyRequest)
	if err != nil {
		er := errorhttp.NewHttpError(err.Error(), http.StatusBadRequest)
		ctx.Error(er)
		return
	}

	newUUID := uuid.New()
	id := "pengirim-" + newUUID.String()

	pengirim := domain.Pengirim{
		Id:           id,
		NamaPengirim: bodyRequest.NamaPengirim,
		NoTelepon:    bodyRequest.NoTelepon,
	}

	tx, err := ph.DB.Begin()
	if err != nil {
		ctx.Error(err)
		return
	}
	defer helper.RollBackCommit(tx)

	result := ph.Pengirim.Add(ctx.Request.Context(), tx, pengirim)

	responseData := domain.PengirimResponse{
		Id:           result.Id,
		NamaPengirim: result.NamaPengirim,
		NoTelepon:    result.NoTelepon,
	}

	response := helper.Response{
		Code:   http.StatusCreated,
		Status: "success",
		Data:   M{"pengirim": responseData},
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusCreated, response)
}
