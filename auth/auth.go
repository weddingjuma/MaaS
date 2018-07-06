package auth

import (
	"github.com/AndroidStudyOpenSource/mpesa-api-go"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

type Handler struct {
	m mpesa.Service
}

func New(m mpesa.Service) Handler {
	return Handler{m: m}
}

// ---------------------------------------------------------------------------------------------------------------------
// AUTHENTICATION
// ---------------------------------------------------------------------------------------------------------------------

func (h Handler) OAuthRequest(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

}
