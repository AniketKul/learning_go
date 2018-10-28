package main

import "fmt"

func main() {
	plantCapacities := []float64{30, 30, 30, 60, 60, 100}
	activePlants := []int{0, 1}
	var gridLoad float64 = 75.
	var option string

	fmt.Println("1) Generate Power Plant Report  ")
	fmt.Println("2) Generate Power Grid Report  ")
	fmt.Println("Please choose an option: ")
	fmt.Scanln(&option)
	println(option)
	askUserForOption(option, plantCapacities, gridLoad, activePlants)
}

func askUserForOption(option string, plantCapacities []float64, gridLoad float64, activePlants []int) {
	switch option {
	case "1":
		generatePlantCapacityReport(plantCapacities...)

	case "2":
		generatePowerGridReport(gridLoad, plantCapacities, activePlants)

	default:
		fmt.Println("Unknown option, no action taken")
	}
}

func generatePlantCapacityReport(plantCapacities ...float64) {
	for idx, cap := range plantCapacities {
		fmt.Printf("Plant %d Capacity: %.0f\n", idx, cap)
	}
}

func generatePowerGridReport(gridLoad float64, plantCapacities []float64, activePlants []int) {
	capacity := 0.
	for _, plantId := range activePlants {
		capacity += plantCapacities[plantId]
	}
	fmt.Printf("%-20s%.0f\n", "Capacity: ", capacity)
	fmt.Printf("%-20s%.0f\n", "Grid Load: ", gridLoad)
	fmt.Printf("%-20s%.1f%%\n", "Utilization: ", (gridLoad/capacity)*100)
}
