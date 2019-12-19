package common

// CalculateRequiredFuel "to find the fuel required for a module, take its mass, divide by three, round down, and subtract 2"
func CalculateRequiredFuel(mass int) int {
	return (mass / 3) - 2
}

// CalculateRequiredFuelIncludingFuelMass "treat the fuel amount you just calculated as the input mass and repeat the process, continuing until a fuel requirement is zero or negative"
func CalculateRequiredFuelIncludingFuelMass(mass int) int {
	fuel := CalculateRequiredFuel(mass)
	sum := 0
	for fuel > 0 {
		sum += fuel
		fuel = CalculateRequiredFuel(fuel)
	}
	return sum
}
