package handler

import "github.com/katerji/flight-path-go/service/flight"

type FindFlightPathRequest struct {
	Flights []flight.Flight `json:"flights"`
}

type FindFlightPathResponse struct {
	FlightPath []string `json:"flight_path"`
	Success    bool     `json:"success"`
}

type FindFlightPathHandler struct{}

func (h FindFlightPathHandler) Handle(request FindFlightPathRequest) FindFlightPathResponse {
	flightService := flight.GetService()

	flightPath, ok := flightService.FindFlightPath(request.Flights)

	return FindFlightPathResponse{
		FlightPath: flightPath,
		Success:    ok,
	}
}
