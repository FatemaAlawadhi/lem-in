package AntColony

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"fmt"
)

//Step one: Parse the file

func ParseFile(filename string) {
	colony = Colony{
		RoomTunnels: make(map[string][]string),
		StartRoom:   "",
		EndRoom:     "",
	}
	NumOfStartRoom := 0
	NumOfEndroom := 0

	file, err := os.Open(filename)
	if err != nil {
		HandleError(err.Error())
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	firstLine := true

	for scanner.Scan() {
		line := scanner.Text()
		var Link []string
		var RoomName []string

		if firstLine {
			// Parse the first line as the number of ants
			// Using strconve.ParseInt to handle the case of large numbers
			numAnts, err := strconv.ParseInt(line, 10, 64)
			if err != nil || numAnts<=0 {
				HandleDataError("invalid number of ants")
			}
			AntsNum = int(numAnts)
			firstLine = false
		} else if strings.HasPrefix(line, "##") {
			word := line[2:]
			var NextLine string
			var NextLineArray []string

			//Save the room value of the next line after ##start or ##end
			if scanner.Scan() {
				NextLine = scanner.Text()
				NextLineArray = strings.Split(NextLine, " ")
			}
			if word == "start" {
				colony.StartRoom = NextLineArray[0]
				colony.RoomTunnels[colony.StartRoom] = make([]string, 0)
				NumOfStartRoom++

			} else if word == "end" {
				colony.EndRoom = NextLineArray[0]
				colony.RoomTunnels[colony.EndRoom] = make([]string, 0)
				NumOfEndroom++
			}
		} else if strings.Contains(line, "-") {
			Link = strings.Split(line, "-")
			//Check if room is linked to unknown room 
			if _, ok := colony.RoomTunnels[Link[0]]; !ok {
				HandleDataError(  "room " +  Link[1] + " is linked to unknow room " + Link[0])
			} else if _, ok := colony.RoomTunnels[Link[1]]; !ok {
				HandleDataError( "room " +  Link[0] + " is linked to unknow room " + Link[1] )
			} else {
				//Check if room is linked to itself
				if Link[0] == Link[1] {
					HandleDataError("room " + Link[0] +" is linked to itself")
				}
				//Check if two tunnels joining same rooms 
				DuplicatedTunnels(Link[0],Link[1])
				//Add tunnel
				colony.RoomTunnels[Link[0]] = append(colony.RoomTunnels[Link[0]], Link[1])
				colony.RoomTunnels[Link[1]] = append(colony.RoomTunnels[Link[1]], Link[0])
			}
		} else if line == "" || strings.HasPrefix(line, "#") {
			// Ignore empty lines and comments
			continue
		} else {
			RoomName = strings.Split(line, " ")
			//Check if room is already exists
			if _, ok := colony.RoomTunnels[RoomName[0]]; ok {
				HandleDataError("room "+ RoomName[0]+ " is duplicated")
			} else {
				colony.RoomTunnels[RoomName[0]] = make([]string, 0)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		HandleError(err.Error())
	}
	
	ValidateData(NumOfStartRoom,NumOfEndroom)
	PrintFileContent(filename)
}


//Validate the number of start and end rooms
func ValidateData(NumOfStartRoom int,NumOfEndroom int) {
	//Check number of start room
	if NumOfStartRoom != 1 {
		if NumOfStartRoom > 1 {
			HandleDataError("more than one start room")
		} else if NumOfStartRoom < 1{
			HandleDataError("no start room found")
		}
	}

	//Check number of end room
	if NumOfEndroom != 1 {
		if NumOfEndroom > 1 {
			HandleDataError("more than one end room")
		} else if NumOfEndroom < 1{
			HandleDataError("no end room found")
		}
	}
}


//Check duplicated tunnels
func DuplicatedTunnels(room1 string,room2 string) {
	for _, room := range colony.RoomTunnels[room1] {
		if room == room2 {
			HandleDataError("two tunnels joining room " + room1 + " and room " + room2)
		}
	}
}

//Print the file after validating the data 
func PrintFileContent(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		HandleError(err.Error())
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "##") {
			fmt.Println(line)
		}else if line == "" || strings.HasPrefix(line, "#") {
			// Skip empty lines and comments
			continue
		}else {
			fmt.Println(line)
		}
	}

	if err := scanner.Err(); err != nil {
		HandleError(err.Error())
	}
}