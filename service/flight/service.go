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
	return nil, false
}
