package application

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func (a *App) brightnessLoop() {
	for {
		a.brightness <- getBrightness()
		time.Sleep(time.Second * 2)
	}
}

func getBrightness() int {
	fd, err := os.ReadFile("/sys/class/backlight/amdgpu_bl0/brightness")
	if err != nil {
		log.Println(err)
		return 0
	}

	val, err := strconv.Atoi(strings.TrimSpace(string(fd)))
	if err != nil {
		log.Println(err)
		return 0
	}

	return (val / 255) * 100
}
