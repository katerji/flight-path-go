package flight

import "testing"

func TestFlightValid(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		f    Flight
		want bool
	}{
		{
			name: "valid flight",
			f:    Flight{"DXB", "AUH"},
			want: true,
		},
		{
			name: "invalid flight with more than 2 destinations",
			f:    Flight{"DXB", "AUH", "BEY"},
			want: false,
		},
		{
			name: "invalid flight with less than 2 destinations",
			f:    Flight{"DXB"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.valid(); got != tt.want {
				t.Errorf("valid() = %v, want %v", got, tt.want)
			}
		})
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
