package service

type IHealthService interface {
	Liveness() bool
	Readiness() bool
}

type healthHandler struct {
}

func NewHealthService() IHealthService {
	return &healthHandler{}
}

// Liveness implements IHealthService.
func (h *healthHandler) Liveness() bool {
	panic("unimplemented")
}

// Readiness implements IHealthService.
func (h *healthHandler) Readiness() bool {
	panic("unimplemented")
}
