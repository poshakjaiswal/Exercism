package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	//panic("CalculateWorkingCarsPerHour not implemented")
	return (float64(productionRate) * successRate) / 100
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	//panic("CalculateWorkingCarsPerMinute not implemented")
	return int((float64(productionRate) * successRate) / 6000)
}

// CalculateCost works out the cost of producing the given number of cars
func CalculateCost(carsCount int) uint {
	//panic("CalculateCost not implemented")
	batches := carsCount % 10
	costOneCarBatchOfTen := 95000
	totalCost := ((carsCount-batches)*costOneCarBatchOfTen)/10 + batches*10000
	return uint(totalCost)
}
