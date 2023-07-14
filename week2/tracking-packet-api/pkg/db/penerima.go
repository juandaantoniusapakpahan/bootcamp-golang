package db

import (
	"context"
	"database/sql"
	"trancking-packet/pkg/domain"
)

type PenerimaImpl struct {
}

func NewPenerima() domain.PenerimaInterfance {
	return &PenerimaImpl{}
}

func (p *PenerimaImpl) Add(ctx context.Context, tx *sql.Tx, penerima domain.Penerima) (domain.Penerima, error) {
	_, err := tx.ExecContext(ctx, "INSERT INTO penerima (penerima_id, nama_penerima, no_telepon) VALUES(?, ?, ?)", penerima.Id, penerima.NamaPenerima, penerima.NoTelepon)
	if err != nil {
		return domain.Penerima{}, err
	}

	return penerima, nil
}
