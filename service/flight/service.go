package flight

type Service struct{}

var instance *Service

func GetService() *Service {
	if instance != nil {
		return instance
	}

	instance = &Service{}

	return instance
}

// FindFlightPath	 reconstructs the travel itinerary from a given list of flight tickets.
//
// The function takes an array of Flight objects, where each Flight consists of a source and a destination airport code.
// It returns the ordered sequence of airport codes representing the travel path and a boolean indicating if a valid itinerary was found.
//
// How It Works:
// - It first validates the input flights using the validFlights function.
// - It builds a mapping of destinations and a graph linking sources to destinations.
// - It identifies the starting airport, which is the source that is not a destination in any ticket.
// - It reconstructs the flight path by following the mapped connections in order.
//
// Parameters:
// - flights ([]Flight): A list of Flight objects representing individual flight tickets.
//
// Returns:
// - ([]string): The reconstructed travel itinerary in order of departure to final destination.
// - (bool): A boolean indicating whether a valid itinerary was found.
//
// Example:
// Given the input:
//   flights = [
//     Flight{source: "LAX", destination: "DXB"},
//     Flight{source: "JFK", destination: "LAX"},
//     Flight{source: "SFO", destination: "SJC"},
//     Flight{source: "DXB", destination: "SFO"}
//   ]
//
// The function returns:
//   (["JFK", "LAX", "DXB", "SFO", "SJC"], true)
func (s Service) FindFlightPath(flights []Flight) ([]string, bool) {
	if !validFlights(flights) {
		return nil, false
	}

	destinations := map[string]bool{}
	flightsGraph := map[string]string{}
	for _, flight := range flights {
		destinations[flight.destination()] = true
		flightsGraph[flight.source()] = flight.destination()
	}

	var origin string
	for _, flight := range flights {
		if _, found := destinations[flight.source()]; !found {
			origin = flight.source()
			break
		}
	}

	var flightsInOrder []string

	currentFlight := origin
	flightsInOrder = append(flightsInOrder, currentFlight)
	for i := 0; i < len(flights); i++ {
		nextDestination := flightsGraph[currentFlight]
		flightsInOrder = append(flightsInOrder, nextDestination)
		currentFlight = nextDestination
	}

	return flightsInOrder, true
}

func validFlights(flights []Flight) bool {
	for _, f := range flights {
		if !f.valid() {
			return false
		}
	}

	return true
}
