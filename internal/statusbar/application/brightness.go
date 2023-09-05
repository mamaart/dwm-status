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
		b, err := getBrightness()
		if err != nil {
			log.Println(err)
			time.Sleep(time.Second * 10)
		} else {
			a.brightness <- b
			time.Sleep(time.Second * 2)
		}
	}
}

func getBrightness() (int, error) {
	data, err := os.ReadFile("/sys/class/backlight/amdgpu_bl0/brightness")
	if err != nil {
		return 0, err
	}

	val, err := strconv.Atoi(strings.TrimSpace(string(data)))
	if err != nil {
		return 0, err
	}

	return (val * 100) / 255, nil
}
