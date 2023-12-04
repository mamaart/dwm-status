package volume

import (
	"github.com/mafik/pulseaudio"
	"github.com/mamaart/statusbar/internal/models"
)

func Get(cli *pulseaudio.Client) (models.Volume, error) {
	v, err := cli.Volume()
	if err != nil {
		return models.Volume(0), err
	}
	return models.Volume(v * 100), nil
}

func Stream(errch chan<- error) (<-chan models.Volume, error) {
	cli, err := pulseaudio.NewClient("/run/user/1000/pulse/native")
	if err != nil {
		return nil, err
	}

	updates, err := cli.Updates()
	if err != nil {
		return nil, err
	}

	ch := make(chan models.Volume)
	go stream(cli, updates, ch, errch)
	return ch, nil
}

func stream(
	cli *pulseaudio.Client,
	input <-chan struct{},
	output chan<- models.Volume,
	err chan<- error,
) {
	v, e := Get(cli)
	if e != nil {
		if err != nil {
			err <- e
		}
	}
	output <- v

	for range input {
		v, e := Get(cli)
		if e != nil {
			if err != nil {
				err <- e
			}
		} else {
			output <- v
		}
	}
}
