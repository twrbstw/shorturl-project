package service

import "database/sql"

type IHealthService interface {
	CheckReadiness() error
}

type healthHandler struct {
	db *sql.DB
}

func NewHealthService(db *sql.DB) IHealthService {
	return &healthHandler{
		db: db,
	}
}

// Readiness implements IHealthService.
func (h *healthHandler) CheckReadiness() error {
	return h.db.Ping()
}
