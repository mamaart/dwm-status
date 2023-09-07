package volume

import (
	"github.com/mafik/pulseaudio"
	"github.com/mamaart/statusbar/internal/models"
)

func Stream() (chan models.Volume, error) {
	cli, err := pulseaudio.NewClient()
	if err != nil {
		return nil, err
	}

	updates, err := cli.Updates()
	if err != nil {
		return nil, err
	}

	v, err := Get(cli)
	if err != nil {
		return nil, err
	}

	ch := make(chan models.Volume)
	ch <- v

	go func(
		cli *pulseaudio.Client,
		input <-chan struct{},
		output chan<- models.Volume,
	) {
		for range input {
			v, err := Get(cli)
			if err != nil {
				continue
			}
			output <- v
		}
	}(cli, updates, ch)

	return ch, nil
}

func Get(cli *pulseaudio.Client) (models.Volume, error) {
	v, err := cli.Volume()
	if err != nil {
		return models.Volume(0), err
	}
	return models.Volume(v * 100), nil
}
