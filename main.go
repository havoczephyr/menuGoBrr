package main

import (
	"log"
	"os"
	"os/exec"
	"time"
)

func main() {
	videoPaths := []string{
		"videos/lazeezBreakfast.mp4",
		"videos/lazeezLunch.mp4",
	}

	changeHour := 12

	for {
		currentHour := time.Now().Hour()

		videoIndex := 0
		if currentHour >= changeHour {
			videoIndex = 1
		}
		// Run the media player with the selected video on loop
		cmd := exec.Command("ffplay", "-loop", "0", "-fs", "-x", "1", videoPaths[videoIndex])
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Start()
		if err != nil {
			log.Println("Video playback error:", err)
			continue
		}

		err = cmd.Wait()
		log.Println("Video playback finished:", err)
	}
}
