package domain

import (
	"context"
	"database/sql"
)

// Service → IdService string, NamaService, HargaPerKg float
type Service struct {
	Id          string
	NamaService string
	HargaPerKg  float32
}

type ServiceRequest struct {
	NamaService string  `json:"nama_service" validate:"required"`
	HargaPerKg  float32 `json:"harga_per_kg" validate:"required"`
}

type ServiceResponse struct {
	Id          string  `json:"service_id"`
	NamaService string  `json:"nama_service"`
	HargaPerKg  float32 `json:"harga_per_kg"`
}

type ServiceInterface interface {
	Add(ctx context.Context, tx *sql.Tx, service Service) (Service, error)
}
