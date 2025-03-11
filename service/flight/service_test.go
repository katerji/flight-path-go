package flight

import (
	"reflect"
	"testing"
)

func TestService_FlightPath(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		flights []Flight
		want    []string
		want1   bool
	}{
		{
			name: "valid test",
			flights: []Flight{
				{"DXB", "AUH"},
				{"AUH", "JFK"},
				{"BEY", "DXB"},
			},
			want: []string{
				"BEY", "DXB", "AUH", "JFK",
			},
			want1: true,
		},
		{
			name:    "no flights",
			flights: []Flight{},
			want:    nil,
			want1:   false,
		},
		{
			name: "single flight",
			flights: []Flight{
				{"DXB", "AUH"},
			},
			want: []string{
				"DXB", "AUH",
			},
			want1: true,
		},
		{
			name: "malformed flights",
			flights: []Flight{
				{"DXB", "AUH"},
				{"DXB", "BEY", "LAO"},
			},
			want:  nil,
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Service{}
			got, got1 := s.FindFlightPath(tt.flights)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FlightPath() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("FlightPath() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
