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

func (s Service) FlightPath(flights []Flight) ([]string, bool) {
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
