package battery

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/mamaart/statusbar/internal/models"
)

func Get() (b models.Battery, err error) {
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
		Charging: string(stat) == "Charging",
		Capacity: value,
	}, nil
}

func Stream(errch chan<- error) (<-chan models.Battery, error) {
	ch := make(chan models.Battery)
	go stream(ch, errch)
	return ch, nil
}

func stream(output chan<- models.Battery, err chan<- error) {
	for {
		w, e := Get()
		if e != nil {
			if err != nil {
				err <- e
			}
			time.Sleep(time.Second * 10)
		} else {
			output <- w
			time.Sleep(time.Minute * 10)
		}
	}
}
