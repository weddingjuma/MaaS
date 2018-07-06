package health

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/AndroidStudyOpenSource/mpesa-api-go"
	"encoding/json"
)

type Handler struct {
	m mpesa.Service
}

func New(m mpesa.Service) Handler {
	return Handler{m: m}
}

// ---------------------------------------------------------------------------------------------------------------------
// HEALTH CHECK
// ---------------------------------------------------------------------------------------------------------------------

func (h Handler) HealthCheck(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	type Health struct {
		Status string `json:"Status"`
	}

	health := new(Health)
	health.Status = "MaaS Service is Healthy"

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(health)
}
