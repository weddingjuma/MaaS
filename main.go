package main

import (
	"github.com/AndroidStudyOpenSource/mpesa-api-go"
	"github.com/julienschmidt/httprouter"
	h "MaaS/health"
	a "MaaS/auth"
)

func main() {

	//Initialize the Database Connection First
	//storer, err := postgres.New()
	//if err != nil {
	//	panic(err)
	//}

	m, err := mpesa.New("", "", mpesa.SANDBOX)
	if err != nil {
		panic(err)
	}

	router := httprouter.New()
	health := h.New(m)
	auth := a.New(m)

	// -----------------------------------------------------------------------------------------------------------------
	// HEALTH CHECK API
	// -----------------------------------------------------------------------------------------------------------------
	router.GET("/health", health.HealthCheck)

	// -----------------------------------------------------------------------------------------------------------------
	// HEALTH CHECK API
	// -----------------------------------------------------------------------------------------------------------------
	router.POST("/auth", auth.OAuthRequest)

}
