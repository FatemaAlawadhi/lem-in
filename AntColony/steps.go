package AntColony


//Step Four: Find the best combination
func CheapestCombination() ([][]string, map[int]int) {
	var CheapestPath [][]string
	var MinSteps int
	AntsInCheapestPath := make(map[int]int)
	for i, combination := range Combinations {
		TotalSteps, AntsInPath := StepsCalculator(combination.Path)
		if TotalSteps < MinSteps || i == 0 {
			MinSteps = TotalSteps
			CheapestPath = combination.Path
			AntsInCheapestPath = AntsInPath
		}
	}

	return CheapestPath, AntsInCheapestPath
}

// Calculate the steps needed for each Combination
func StepsCalculator(combination [][]string) (int, map[int]int) {
	AntsInPath := make(map[int]int)
	LeftAnts := AntsNum
	var next int
	for LeftAnts > 0 {
		for i, path := range combination {
			if i == len(combination)-1 {
				next = 0
			} else {
				next = i + 1
			}
			for len(path)+AntsInPath[i] <= len(combination[next])+AntsInPath[next] && LeftAnts > 0 {
				AntsInPath[i] = AntsInPath[i] + 1
				LeftAnts -= 1
			}
		}
	}

	Steps := len(combination[0]) + AntsInPath[0]
	for i, path := range combination {
		if len(path)+AntsInPath[i] > Steps {
			Steps = len(path) + AntsInPath[i]
		}
	}

	return Steps, AntsInPath
}
