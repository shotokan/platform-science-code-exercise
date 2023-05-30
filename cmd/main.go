package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/shotokan/platform-science-code-exercise/internal/calculatess"
	"github.com/shotokan/platform-science-code-exercise/pkg/input"
)

func main() {
	driversPath, err := input.ReadPath("Path to drivers file")
	if err != nil {
		log.Fatalf("error reading file input: %s", err)
	}

	destnationsPath, err := input.ReadPath("Path to destinations file")
	if err != nil {
		log.Fatalf("error reading file input: %s", err)
	}

	// Read the shipment destinations from the first file
	destinationsFile, err := ioutil.ReadFile(destnationsPath)
	if err != nil {
		log.Fatal("Error reading destinations file:", err)
	}

	destinations := strings.Split(strings.TrimSpace(string(destinationsFile)), "\n")

	// Read the driver names from the second file
	driversFile, err := ioutil.ReadFile(driversPath)
	if err != nil {
		log.Fatal("Error reading drivers file:", err)
	}

	drivers := strings.Split(strings.TrimSpace(string(driversFile)), "\n")

	matching := calculatess.CalculateSS(drivers, destinations)
	totalSS := 0.0

	// Print the matching and total SS
	fmt.Println("Matching:")
	for driver, assigment := range matching {
		fmt.Printf("%s -> %s\n", driver, assigment.Destination)
		totalSS += assigment.SS
	}

	fmt.Printf("Total SS: %.2f\n", totalSS)
}
