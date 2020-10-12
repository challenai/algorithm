package solution

func canCompleteCircuit(gas, cost []int) int {
	if len(gas) != len(cost) {
		return -1
	}
	for i := 0; i < len(gas); i++ {
		cost[i] = gas[i] - cost[i]
	}

	ptr := 0
	current := cost[0]
	for ptr < len(gas) {
		if cost[i]
	}

	return 0
}
