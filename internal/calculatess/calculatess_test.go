package calculatess

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCounter_HaveCommonFactor(t *testing.T) {
	t.Run("CalculateSS is true when a common factor other than one", func(t *testing.T) {
		t.Parallel()
		destinations := getDestinations()
		drivers := getDrivers()
		want := make(map[string]Assignment)
		want[drivers[0]] = Assignment{
			Driver:      drivers[0],
			Destination: destinations[0],
			SS:          7.5,
		}

		got := CalculateSS(drivers, destinations)

		if !cmp.Equal(want, got) {
			t.Error("the numbers have no common factor",
				cmp.Diff(want, got))
		}

	})

	t.Run("CalculateSS is true when common factor is one", func(t *testing.T) {
		t.Parallel()
		drivers := getDrivers()
		destinations := []string{
			"Aalto",
		}

		want := make(map[string]Assignment)
		want[drivers[0]] = Assignment{
			Driver:      drivers[0],
			Destination: destinations[0],
			SS:          5,
		}

		got := CalculateSS(drivers, destinations)

		if !cmp.Equal(want, got) {
			t.Error("SS is not correct",
				cmp.Diff(want[drivers[0]].SS, got[drivers[0]].SS))
		}

	})

	t.Run("CalculateSS assigns each shipment destination to a given driver while maximizing the total suitability of all shipments to all drivers",
		func(t *testing.T) {
			t.Parallel()
			destinations := getDestinations()
			drivers := getDrivers()
			drivers = append(drivers, "Monica Hermann")
			destinations = append(destinations, "1234 Fake St., San Diego, CA 92126")

			want := make(map[string]Assignment)
			want[drivers[1]] = Assignment{
				Driver:      drivers[1],
				Destination: destinations[0],
				SS:          8,
			}
			want[drivers[0]] = Assignment{
				Driver:      drivers[0],
				Destination: destinations[1],
				SS:          7.5,
			}

			got := CalculateSS(drivers, destinations)

			if !cmp.Equal(want, got) {
				t.Error("SS is not correct",
					cmp.Diff(want, got))
			}

		})

}

func getDestinations() []string {
	return []string{
		"Aalto Place",
	}
}

func getDrivers() []string {
	return []string{
		"Carlo Abate",
	}
}
