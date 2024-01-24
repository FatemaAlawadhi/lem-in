package AntColony

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func Moves(combination [][]string, AntsInPath map[int]int) {
	ReachedEnd := 0
	AntEntered := 0
	AntsInRoom := make(map[string]string)
	for ReachedEnd != AntsNum {
		MovesLine := ""
		for AntName, Room := range AntsInRoom {
			if ReachedEnd != AntsNum {
				NextRoom := GetNext(combination, Room)
				AntsInRoom[AntName] = NextRoom
				MovesLine = MovesLine + " " + AntName + "-" + NextRoom
				if NextRoom == colony.EndRoom {
					delete(AntsInRoom, AntName)
					ReachedEnd++
				}
			} else {
				continue
			}
		}

		for i, path := range combination {
			var Antname string
			if AntsInPath[i] > 0 {
				AntEntered++
				Antname = "L" + strconv.Itoa(AntEntered)
				AntsInPath[i] -= 1
				MovesLine = MovesLine + " " + Antname + "-" + path[0]
				if path[0] == colony.EndRoom {
					ReachedEnd++
				}else {
					AntsInRoom[Antname] = path[0]
				}
			}
		}
		MovesLine = sortAntLine(MovesLine)
		fmt.Println(MovesLine)
	}

}

func GetNext(combination [][]string, Room string) string {
	for _, path := range combination {
		for a, RoomName := range path {
			if RoomName == Room {
				return path[a+1]
			}
		}
	}
	return ""
}


func sortAntLine(line string) string {
	splitLine := strings.Fields(line)

	sort.Slice(splitLine, func(i, j int) bool {
		antNum1 := getAntNumber(splitLine[i])
		antNum2 := getAntNumber(splitLine[j])
		return antNum1 < antNum2
	})

	sortedLine := strings.Join(splitLine, " ")
	return sortedLine
}

func getAntNumber(ant string) int {
	antNum := strings.Split(ant, "-")[0]
	num, err := strconv.Atoi(antNum[1:])
	if err != nil {
		HandleError(err.Error())
	}
	return num
}
