package cmd

import (
	"ewallet-notification/helpers"
	"ewallet-notification/internal/api"
	"ewallet-notification/internal/services"
	"log"

	"github.com/gin-gonic/gin"
)

func ServeHttp() {
	healthCheckSvc := &services.Healthcheck{}
	healtCheckAPI := &api.Healthcheck{
		HealthcheckServices: healthCheckSvc,
	}

	r := gin.Default()

	r.GET("/health", healtCheckAPI.HealthcheckHandlerHTTP)

	err := r.Run(":" + helpers.GetEnv("PORT", ""))
	if err != nil {
		log.Fatal(err)
	}
}
