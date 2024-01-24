package main

import ("os"
		"lem-in/AntColony"
	    "fmt")

func main() {
	Arg := os.Args 
	if len(Arg) != 2 {
		fmt.Println("Usage Example: go run . example00.txt")
	}else {
		//1. To Parse the file
		AntColony.ParseFile("Audit/" + Arg[1])

		//2. Find all possible paths
		AntColony.FindAllPaths()

		//3. Find all possible combinations of the paths
		AntColony.FindCombinations()

		//4. Find the best combination
		combinationh, AntsInPath:= AntColony.CheapestCombination()

		//5. Regulate the ants movement + 6. Handle output
		AntColony.Moves(combinationh, AntsInPath) 
	}

}
