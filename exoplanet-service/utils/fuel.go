package utils

import (
	"errors"
	"exoplanet-service/models"
)

func CalculateFuel(e models.Exoplanet, crewCapacity int) (float64, error) {
	if crewCapacity <= 0 {
		return 0, errors.New("crew capacity must be greater than 0")
	}

	var gravity float64
	switch e.Type {
	case "GasGiant":
		gravity = 0.5 / (e.Radius * e.Radius)
	case "Terrestrial":
		gravity = e.Mass / (e.Radius * e.Radius)
	default:
		return 0, errors.New("unknown exoplanet type")
	}

	fuel := float64(e.Distance) / (gravity * gravity) * float64(crewCapacity)
	return fuel, nil
}
