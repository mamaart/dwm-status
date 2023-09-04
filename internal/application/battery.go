package application

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/mamaart/statusbar/internal/models"
)

func (a *App) batteryLoop() {
	for {
		bat, err := GetBattery()
		if err != nil {
			fmt.Println(err)
			time.Sleep(time.Second * 10)
		}
		a.battery <- bat
		time.Sleep(time.Second * 2)

	}
}

func GetBattery() (b models.Battery, err error) {
	capacity, err := os.ReadFile("/sys/class/power_supply/BAT0/capacity")
	if err != nil {
		return b, fmt.Errorf("failed to open capacity file: %s", err)
	}
	stat, err := os.ReadFile("/sys/class/power_supply/BAT0/status")
	if err != nil {
		return b, fmt.Errorf("failed to open status file: %s", err)
	}

	value, err := strconv.Atoi(string(strings.TrimSpace(string(capacity))))
	if err != nil {
		return b, fmt.Errorf("failed to parse capacity: %s", err)
	}

	return models.Battery{
		Charging: string(stat) != "Discharging", //Chargin || Not charging
		Capacity: value,
	}, nil
}
