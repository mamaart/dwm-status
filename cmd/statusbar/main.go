package main

import (
	"time"

	"github.com/mamaart/statusbar/internal/api"
	"github.com/mamaart/statusbar/internal/ui"
	"github.com/mamaart/statusbar/modules/batterymodule"
	"github.com/mamaart/statusbar/modules/brightnessmodule"
	"github.com/mamaart/statusbar/modules/diskmodule"
	"github.com/mamaart/statusbar/modules/netmodule"
	"github.com/mamaart/statusbar/modules/textmodule"
	"github.com/mamaart/statusbar/modules/timemodule"
	"github.com/mamaart/statusbar/modules/volumemodule"
	"github.com/mamaart/statusbar/modules/wttrmodule"
)

func main() {
	var (
		tim = timemodule.New()
		bat = batterymodule.New()
		vol = volumemodule.New()
		bri = brightnessmodule.New()
		wtr = wttrmodule.New()
		net = netmodule.New()
		dsk = diskmodule.New()
		txt = textmodule.New(textmodule.Options{
			WindowWidth: 80,
			Delay:       time.Millisecond * 150,
		})
	)

	go api.Run(tim, txt)
	ui.Run(net, dsk, bri, vol, bat, tim, wtr, txt)
}
