package main

import (
	"github.com/gin-gonic/gin"
	"exoplanet-service/handlers"
)

func main() {
	r := gin.Default()

	r.POST("/exoplanets", handlers.AddExoplanet)
	r.GET("/exoplanets", handlers.ListExoplanets)
	r.GET("/exoplanets/:id", handlers.GetExoplanetByID)
	r.PUT("/exoplanets/:id", handlers.UpdateExoplanet)
	r.DELETE("/exoplanets/:id", handlers.DeleteExoplanet)
	r.GET("/exoplanets/:id/fuel", handlers.FuelEstimation)

	r.Run(":8080")
}
