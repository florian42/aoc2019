package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("./Day1/input")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	lines := strings.Split(string(input), "\n")
	fuel := 0
	for _, l := range lines {
		if len(l) == 0 { continue }
		mass, err := strconv.Atoi(l)
		if err != nil {
			panic("Could not convert line to int!")
		}
		fuel += calculateFuelNeededForFuel(calculateFuelRequiredToLaunch(mass))
	}
	fmt.Printf("Fuel required: %d\n", fuel)

}

func calculateFuelRequiredToLaunch(mass int) int {
	return int(math.Floor(float64(mass)/3) - 2)
}

func calculateFuelNeededForFuel(mass int) int {
	if mass <= 0 {return 0}
	return mass + calculateFuelNeededForFuel(calculateFuelRequiredToLaunch(mass))
}