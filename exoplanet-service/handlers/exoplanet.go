package handlers

import (
	"net/http"
	"strconv"

	"exoplanet-service/models"
	"exoplanet-service/utils"

	"github.com/gin-gonic/gin"
)

func AddExoplanet(c *gin.Context) {
	var e models.Exoplanet
	if err := c.ShouldBindJSON(&e); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.ValidateExoplanet(e); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	e = models.AddExoplanet(e)
	c.JSON(http.StatusOK, e)
}

func ListExoplanets(c *gin.Context) {
	exoplanets := models.ListExoplanets()
	c.JSON(http.StatusOK, exoplanets)
}

func GetExoplanetByID(c *gin.Context) {
	id := c.Param("id")
	exoplanet, err := models.GetExoplanetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, exoplanet)
}

func UpdateExoplanet(c *gin.Context) {
	id := c.Param("id")
	var e models.Exoplanet
	if err := c.ShouldBindJSON(&e); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.ValidateExoplanet(e); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	e, err := models.UpdateExoplanet(id, e)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, e)
}

func DeleteExoplanet(c *gin.Context) {
	id := c.Param("id")
	if err := models.DeleteExoplanet(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func FuelEstimation(c *gin.Context) {
	id := c.Param("id")
	crewCapacityStr := c.Query("crew")
	if crewCapacityStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "crew capacity is required"})
		return
	}

	crewCapacity, err := strconv.Atoi(crewCapacityStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid crew capacity"})
		return
	}

	exoplanet, err := models.GetExoplanetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	fuel, err := utils.CalculateFuel(exoplanet, crewCapacity)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"fuel": fuel})
}
