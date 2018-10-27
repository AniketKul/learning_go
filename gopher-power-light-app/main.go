package main

import "fmt"

func main() {
	var plantCapacities []float64

	plantCapacities = []float64{30, 30, 30, 60, 60, 100}

	var capacity float64 = plantCapacities[0] + plantCapacities[1] +
		plantCapacities[2] + plantCapacities[3] + plantCapacities[4] +
		plantCapacities[5]

	var gridLoad float64 = 75.

	utilization := gridLoad / capacity

	fmt.Println("Capacity: ", capacity)
	fmt.Println("Grid Load: ", gridLoad)
	fmt.Println("Utilization: ", utilization*100)
}
