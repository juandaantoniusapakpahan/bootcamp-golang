package db

import (
	"context"
	"database/sql"
	"trancking-packet/pkg/domain"
)

type ServiceImpl struct {
}

func NewService() domain.ServiceInterface {
	return &ServiceImpl{}
}

func (s *ServiceImpl) Add(ctx context.Context, tx *sql.Tx, service domain.Service) (domain.Service, error) {
	_, err := tx.ExecContext(
		ctx,
		"INSERT INTO services (service_id, nama_service, harga_per_kg) VALUES(?, ?, ?)",
		service.Id, service.NamaService, service.HargaPerKg)

	if err != nil {
		return domain.Service{}, err
	}
	return service, nil
}
