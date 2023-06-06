package main

import (
	"os"
	"os/exec"
	"time"
)

// player will play the 2 provided videos on loop.
// params will adjust what time to play lazeezBreakfast or lazeezLunch.
// play will return an error if an error occurs.
func player(timeChoice int) error {

	videoPaths := []string{
		"videos/lazeezBreakfast.mp4",
		"videos/lazeezLunch.mp4",
	}

	changeHour := timeChoice

	for {
		currentHour := time.Now().Hour()

		videoIndex := 0
		if currentHour >= changeHour {
			videoIndex = 1
		}
		cmd := exec.Command("ffplay", "-loop", "0", "-fs", "-x", "1", videoPaths[videoIndex])
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Start()
		if err != nil {
			return err
		}

		err = cmd.Wait()
		if err != nil {
			return err
		}
	}
}
