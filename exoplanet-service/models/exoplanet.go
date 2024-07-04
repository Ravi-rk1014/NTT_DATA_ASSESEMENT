package models

import (
	"errors"
	"math/rand"
)

type Exoplanet struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Distance    int     `json:"distance"`
	Radius      float64 `json:"radius"`
	Mass        float64 `json:"mass,omitempty"`
	Type        string  `json:"type"`
}

var exoplanets = make(map[string]Exoplanet)

func ValidateExoplanet(e Exoplanet) error {
	if e.Name == "" || e.Description == "" {
		return errors.New("name and description are required")
	}
	if e.Distance < 10 || e.Distance > 1000 {
		return errors.New("distance must be between 10 and 1000 light years")
	}
	if e.Radius < 0.1 || e.Radius > 10 {
		return errors.New("radius must be between 0.1 and 10 Earth-radius units")
	}
	if e.Type == "Terrestrial" {
		if e.Mass < 0.1 || e.Mass > 10 {
			return errors.New("mass must be between 0.1 and 10 Earth-mass units for terrestrial planets")
		}
	}
	return nil
}

func AddExoplanet(e Exoplanet) Exoplanet {
	e.ID = generateID()
	exoplanets[e.ID] = e
	return e
}

func ListExoplanets() []Exoplanet {
	list := []Exoplanet{}
	for _, e := range exoplanets {
		list = append(list, e)
	}
	return list
}

func GetExoplanetByID(id string) (Exoplanet, error) {
	e, exists := exoplanets[id]
	if !exists {
		return Exoplanet{}, errors.New("exoplanet not found")
	}
	return e, nil
}

func UpdateExoplanet(id string, updated Exoplanet) (Exoplanet, error) {
	_, err := GetExoplanetByID(id)
	if err != nil {
		return Exoplanet{}, err
	}
	updated.ID = id
	exoplanets[id] = updated
	return updated, nil
}

func DeleteExoplanet(id string) error {
	_, err := GetExoplanetByID(id)
	if err != nil {
		return err
	}
	delete(exoplanets, id)
	return nil
}

func generateID() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 8)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
