package main

var (
	gas  = []int{2, 3, 4}
	cost = []int{3, 4, 3}
)

func canCompleteCircuit(gas []int, cost []int) int {
	var totalTank, currentTank, startStation int

	for i := range gas {
		totalTank += gas[i] - cost[i]
		currentTank += gas[i] - cost[i]

		if currentTank < 0 {
			startStation = i + 1
			currentTank = 0
		}
	}

	if totalTank < 0 {
		return -1
	}
	return startStation
}
