package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64 ( float64 (productionRate) * (successRate / 100.0))
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int (CalculateWorkingCarsPerHour(productionRate, successRate) / 60)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	var setsOfTenCars = carsCount / 10
    var remainderCars = carsCount % 10
    var tensOfCarsTotal =  setsOfTenCars * 95000
    var remainderCarsTotal = remainderCars * 10000
    return uint (tensOfCarsTotal + remainderCarsTotal)
}
