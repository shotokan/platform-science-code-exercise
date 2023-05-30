package calculatess

import (
	"strings"

	"github.com/shotokan/platform-science-code-exercise/pkg/commonfactor"
	"github.com/shotokan/platform-science-code-exercise/pkg/counter"
)

type Assignment struct {
	Driver      string
	Destination string
	SS          float64
}

// CalculateSSassigns shipment destinations to drivers in a way that
// maximizes the total SS over the set of drivers
//
// CalculateSS uses a mathematical model for determining which
// drivers are best suited to deliver each shipment based on:
//
// If the length of the shipment's destination street name is even,
// the base suitability score (SS) is the number of vowels in the driver’s
// name multiplied by 1.5.
//
// If the length of the shipment's destination street name is odd,
// the base SS is the number of consonants in the driver’s name multiplied by 1.
//
// If the length of the shipment's destination street name shares any
// common factors (besides 1) with the length of the driver’s name, the
// SS is increased by 50% above the base SS.
//
// Args:
//     drivers: slice of strings with the driver names.
//     destinations: slice of strings with the destination names.
//
// Returns:
//     A map with the destination assignment data of each driver.
//
func CalculateSS(drivers, destinations []string) map[string]Assignment {
	matching := make(map[string]Assignment)

	total := 0.0

	for _, destination := range destinations {
		bestSS := 0.0
		bestDriver := ""

		for _, driver := range drivers {
			ss := getSSForDriver(driver, destination)

			_, ok := matching[driver]

			if ss > bestSS && !ok {
				bestSS = ss
				bestDriver = driver
			}
		}

		if bestDriver != "" {
			total += bestSS
			matching[bestDriver] = Assignment{
				Destination: destination,
				Driver:      bestDriver,
				SS:          bestSS,
			}
		}
	}
	return matching
}

func getSSForDriver(driverName, destination string) float64 {
	driverName = strings.ToLower(driverName)
	destination = strings.ToLower(destination)

	baseSS := 0.0

	if len(destination)%2 == 0 {
		vowelCount := counter.Vowels(driverName)
		baseSS = float64(vowelCount) * 1.5
	} else {
		consonantCount := counter.Consonants(driverName)
		baseSS = float64(consonantCount)
	}

	commonFactor := commonfactor.HaveCommonFactors(len(driverName), len(destination))

	if commonFactor {
		baseSS *= 1.5
	}

	return baseSS
}
