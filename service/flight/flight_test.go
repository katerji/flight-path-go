package flight

import "testing"

func TestFlightValid(t *testing.T) {
	t.Parallel()

	invalidFlight := Flight{
		"DXB", "AUH", "BEY",
	}
	valid := invalidFlight.valid()
	if valid {
		t.Error("expected flight to be invalid")
	}
}

func TestFlightSource(t *testing.T) {
	t.Parallel()

	const (
		expectedSource      = "DXB"
		expectedDestination = "AUH"
	)
	flight := Flight{
		expectedSource, expectedDestination,
	}
	source := flight.source()
	if source != expectedSource {
		t.Errorf("expected source to be %s, got %s instead", expectedSource, source)
	}
}

func TestFlightDestination(t *testing.T) {
	t.Parallel()

	const (
		expectedSource      = "DXB"
		expectedDestination = "AUH"
	)
	flight := Flight{
		expectedSource, expectedDestination,
	}
	destination := flight.destination()
	if destination != expectedDestination {
		t.Errorf("expected source to be %s, got %s instead", expectedDestination, destination)
	}
}
