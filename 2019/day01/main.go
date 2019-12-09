package main

import (
	"fmt"
	"log"
	"math"
	"strconv"

	"github.com/blackskad/adventofcode/io"
)

func fuelForModule(mass float64) float64 {
	fuel := math.Floor(float64(mass)/3) - 2

	total := 0.
	for fuel > 0 {
		total += fuel

		fuel = (math.Floor(float64(fuel)/3) - 2)
		fmt.Printf("%f\n", fuel)
	}
	return total
}

func main() {
	modules, err := io.ReadLinesFromFile("data.txt")
	if err != nil {
		log.Fatal(err.Error())
	}

	//modules = []string{"14"}

	drysum := 0.
	sum := 0.
	for _, module := range modules {
		mass, err := strconv.Atoi(module)
		if err != nil {
			log.Fatal(err.Error())
		}

		drysum += (math.Floor(float64(mass)/3) - 2)
		sum += fuelForModule(float64(mass))
	}
	fmt.Printf("%.f\n", drysum)
	fmt.Printf("%.f\n", sum)
}
