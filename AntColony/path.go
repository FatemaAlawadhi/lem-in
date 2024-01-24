package AntColony

import (
	"sort"
	"strings"
)

//Step two: Find all possible paths

// Function to find all possible paths
func FindAllPaths() {
	visited := make(map[string]bool)
	findPaths(colony.RoomTunnels, colony.StartRoom, colony.EndRoom, []string{}, visited, PossiblePaths)
	if len(PossiblePaths) == 0 {
		HandleDataError("no path between start and end")
	}
}

// Recursive function to find all paths
func findPaths(graph map[string][]string, currentNode string, destination string, path []string, visited map[string]bool, paths [][]string) {
	path = append(path, currentNode)
	visited[currentNode] = true

	if currentNode == destination {
		// Append the completed path to the paths slice
		PossiblePaths = append(PossiblePaths, append([]string(nil), path...))
	} else {
		for _, neighbor := range graph[currentNode] {
			if !visited[neighbor] {
				findPaths(graph, neighbor, destination, path, visited, paths)
			}
		}
	}

	// Remove the current node from the path and mark it as unvisited
	path = path[:len(path)-1]
	visited[currentNode] = false
}


//Step three: Find all possible combinations of the paths

func FindCombinations() {
	for _, Path := range PossiblePaths {
		combinedPath := ComparePath(Path[1:])
		Combinations = append(Combinations, combinedPath)
	}

	//Filter identical combinations
	FilterSimilarCombinations()
	
	//Arrange combinations from shortest to longest
	Combinations = SortSubpathsByLength(Combinations)
}

func ComparePath(Path []string) CombinedPath {
	combinedPath := CombinedPath{
		Path: make([][]string, 0),
	}
	combinedPath.Path = append(combinedPath.Path, Path)

	var UsedNodes = Path
	for _, Path2 := range PossiblePaths {
		Path2 := Path2[1:]
		if !equalSlices(Path, Path2) {
			Used := false
			for _, Path2Nodes := range Path2 {
				for _, Nodes := range UsedNodes {
					if Path2Nodes == Nodes && Nodes != colony.EndRoom {
						Used = true
						break
					}
				}
			}
			if !Used {
				combinedPath.Path = append(combinedPath.Path, Path2)
				UsedNodes = append(UsedNodes, Path2...)
			}
		}
	}

	return combinedPath
}

func equalSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// Remove Similar Combinations
func FilterSimilarCombinations() {
	filteredCombinations := make([]CombinedPath, 0)
	visitedPaths := make(map[string]bool)

	for _, combination := range Combinations {
		sortedPath := getSortedPath(combination.Path)
		if !visitedPaths[sortedPath] {
			visitedPaths[sortedPath] = true
			filteredCombinations = append(filteredCombinations, combination)
		}
	}
	Combinations = filteredCombinations
}

func getSortedPath(path [][]string) string {
	sortedPath := make([]string, len(path))
	for i, subpath := range path {
		sortedSubpath := make([]string, len(subpath))
		copy(sortedSubpath, subpath)
		sort.Strings(sortedSubpath)
		sortedPath[i] = strings.Join(sortedSubpath, "-")
	}
	sort.Strings(sortedPath)
	return strings.Join(sortedPath, "#")
}

// Sort each combination from the shortest path to the longest
func SortSubpathsByLength(combinations []CombinedPath) []CombinedPath {
	for i := range combinations {
		sort.SliceStable(combinations[i].Path, func(j, k int) bool {
			return len(combinations[i].Path[j]) < len(combinations[i].Path[k])
		})
	}

	return combinations
}
