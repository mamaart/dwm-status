package brightnessmodule

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type BrightnessModule struct {
	output <-chan []byte
}

func New() *BrightnessModule {
	output := make(chan []byte)
	go func() {
		for {
			w, err := Get()
			if err != nil {
				log.Println(err)
				time.Sleep(time.Second * 10)
			} else {
				output <- []byte(w.String())
				// TODO: get real time updates not time.Sleep
				time.Sleep(time.Second * 2)
			}
		}
	}()
	return &BrightnessModule{output: output}
}

func (b *BrightnessModule) Reader() <-chan []byte {
	return b.output
}

func Get() (Brightness, error) {
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

	return Brightness((val * 100) / 255), nil
}
