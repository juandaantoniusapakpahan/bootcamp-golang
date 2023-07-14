package db

import (
	"context"
	"database/sql"
	"trancking-packet/pkg/domain"
)

type LokasiImpl struct {
}

func NewLokasi() domain.LokasiInterface {
	return &LokasiImpl{}
}

func (l *LokasiImpl) Add(ctx context.Context, tx *sql.Tx, lokasi domain.Lokasi) (domain.Lokasi, error) {
	_, err := tx.ExecContext(ctx, "INSERT INTO lokasi(lokasi_id, nama_lokasi, alamat) VALUES(?,?,?)", lokasi.Id, lokasi.NamaLokasi, lokasi.Alamat)
	if err != nil {
		return domain.Lokasi{}, err
	}
	return lokasi, nil
}
