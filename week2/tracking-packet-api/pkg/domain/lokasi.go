package domain

import (
	"context"
	"database/sql"
)

// Lokasi → IdLokasi, NamaLokasi, Alamat → formatIdLokasi paket == <gdng-00001>

type Lokasi struct {
	Id         string
	NamaLokasi string
	Alamat     string
}

type LokasiRequest struct {
	NamaLokasi string `json:"nama_lokasi" validate:"required"`
	Alamat     string `json:"alamat" validate:"required"`
}

type LokasiResponse struct {
	Id         string `json:"lokasi_id"`
	NamaLokasi string `json:"nama_lokasi"`
	Alamat     string `json:"alamat"`
}

type LokasiInterface interface {
	Add(ctx context.Context, tx *sql.Tx, lokasi Lokasi) (Lokasi, error)
}
