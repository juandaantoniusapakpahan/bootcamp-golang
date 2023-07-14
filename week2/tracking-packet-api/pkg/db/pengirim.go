package db

import (
	"context"
	"database/sql"
	"trancking-packet/pkg/domain"
)

type PengirimImpl struct {
}

func NewPengirim() domain.PengirimInterfance {
	return &PengirimImpl{}
}

func (p *PengirimImpl) Add(ctx context.Context, tx *sql.Tx, pengirim domain.Pengirim) domain.Pengirim {
	_, err := tx.ExecContext(ctx, "INSERT INTO pengirim (pengirim_id, nama_pengirim, no_telepon) VALUES(?,?, ?)", pengirim.Id, pengirim.NamaPengirim, pengirim.NoTelepon)
	if err != nil {
		panic(err)
	}
	return pengirim
}
