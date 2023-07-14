package domain

import (
	"context"
	"database/sql"
)

// Penerima → IdPenerima string, NamaPenerima string, noTelp string

type Penerima struct {
	Id           string
	NamaPenerima string
	NoTelepon    string
}

type PenerimaRequest struct {
	NamaPenerima string `json:"nama_penerima" validate:"required"`
	NoTelepon    string `json:"no_telepon" validate:"required,number"`
}

type PenerimaReponse struct {
	Id           string `json:"penerima_id"`
	NamaPenerima string `json:"nama_penerima"`
	NoTelepon    string `json:"no_telepon"`
}

type PenerimaInterfance interface {
	Add(ctx context.Context, tx *sql.Tx, penerima Penerima) (Penerima, error)
}
