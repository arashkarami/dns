package models

import (
	"strconv"
)

// swagger:parameters locationRequest
type LocationRequest struct {
	// type: string
	//	in: body
	X   string `json:"x"`
	Y   string `json:"y"`
	Z   string `json:"z"`
	Vel string `json:"vel"`
}

type Location struct {
	X   float64 `json:"x"`
	Y   float64 `json:"y"`
	Z   float64 `json:"z"`
	Vel float64 `json:"vel"`
}

// swagger:response locationResponse
type LocationResponse struct {
	Loc float64 `json:"loc"`
}

func CreateLocationResponse(loc float64) LocationResponse {
	return LocationResponse{Loc: loc}
}

func CreateLocation(req *LocationRequest) (Location, error) {
	model := Location{}

	X, err := parseStringToFloat(req.X)
	if err != nil {
		return model, err
	}
	Y, err := parseStringToFloat(req.Y)
	if err != nil {
		return model, err
	}
	Z, err := parseStringToFloat(req.Z)
	if err != nil {
		return model, err
	}
	Vel, err := parseStringToFloat(req.Vel)
	if err != nil {
		return model, err
	}

	model.X = X
	model.Y = Y
	model.Z = Z
	model.Vel = Vel

	return model, nil
}

func parseStringToFloat(dimension string) (float64, error) {
	dimensionFloat, err := strconv.ParseFloat(dimension, 64)
	if err != nil {
		return 0, err
	}

	return dimensionFloat, nil
}
