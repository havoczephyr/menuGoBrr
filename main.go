package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	timeOption, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Failed to convert string", err)
		return
	}
	isTimeInvalid := (timeOption < 0 || timeOption > 23)
	if isTimeInvalid {
		fmt.Println("Your time value is incorrect. please choose an hour from 0 to 23.")
		return
	} else {
		err := player(timeOption)
		if err != nil {
			fmt.Println("playback error", err)
			return
		}
	}

}
