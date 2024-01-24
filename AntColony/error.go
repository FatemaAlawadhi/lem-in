package AntColony

import (
	"fmt"
	"log"
)

//Using log fetal to print the error and quit the program
func HandleDataError(err string) {
	errorMessage := fmt.Sprintf("\033[31mError: invalid data format, %s\033[0m\n", err)
	log.Fatal(errorMessage)
}

func HandleError(err string) {
	errorMessage := fmt.Sprintf("\033[31mError: %s\033[0m\n", err)
	log.Fatal(errorMessage)
}