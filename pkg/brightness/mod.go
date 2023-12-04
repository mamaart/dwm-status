package brightness

import (
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/mamaart/statusbar/internal/models"
)

func Get() (models.Brightness, error) {
	data, err := os.ReadFile(
		"/sys/class/backlight/amdgpu_bl1/brightness",
	) // TODO make more generic (amdgpu_bl0)
	if err != nil {
		return 0, err
	}

	val, err := strconv.Atoi(strings.TrimSpace(string(data)))
	if err != nil {
		return 0, err
	}

	return models.Brightness((val * 100) / 255), nil
}

func Stream(errch chan<- error) (<-chan models.Brightness, error) {
	ch := make(chan models.Brightness)
	go stream(ch, errch)
	return ch, nil
}

func stream(output chan<- models.Brightness, err chan<- error) {
	for {
		w, e := Get()
		if e != nil {
			if err != nil {
				err <- e
			}
			time.Sleep(time.Second * 10)
		} else {
			output <- w
			time.Sleep(time.Second * 2)
		}
	}
}
