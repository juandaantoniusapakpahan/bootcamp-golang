package domain

import (
	"context"
	"database/sql"
)

// Pengirim → IdPengirim string, NamaPengirim string, noTelp string

type Pengirim struct {
	Id           string
	NamaPengirim string
	NoTelepon    string
}

type PengirimRequest struct {
	NamaPengirim string `json:"nama_pengirim" validate:"required,alpha"`
	NoTelepon    string `json:"no_telepon"  validate:"required,number"`
}

type PengirimResponse struct {
	Id           string `json:"pengirim_id"`
	NamaPengirim string `json:"nama_pengirim"`
	NoTelepon    string `json:"no_telepon"`
}

type PengirimInterfance interface {
	Add(ctx context.Context, tx *sql.Tx, pengirim Pengirim) Pengirim
}
