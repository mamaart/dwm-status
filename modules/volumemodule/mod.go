package volumemodule

import (
	"log"

	"github.com/mafik/pulseaudio"
)

type VolumeModule struct {
	output <-chan []byte
}

func New() *VolumeModule {
	cli, err := pulseaudio.NewClient("/run/user/1000/pulse/native")
	if err != nil {
		log.Fatal(err)
	}

	updates, err := cli.Updates()
	if err != nil {
		log.Fatal(err)
	}

	output := make(chan []byte)

	go func() {
		v, e := Get(cli)
		if e != nil {
			if err != nil {
				log.Println(err)
			}
		}
		output <- []byte(v.String())

		for range updates {
			v, e := Get(cli)
			if e != nil {
				if err != nil {
					log.Println(err)
				}
			} else {
				output <- []byte(v.String())
			}
		}
	}()

	return &VolumeModule{output}
}

func (v *VolumeModule) Reader() <-chan []byte {
	return v.output
}

func Get(cli *pulseaudio.Client) (Volume, error) {
	v, err := cli.Volume()
	if err != nil {
		return Volume(0), err
	}
	return Volume(v * 100), nil
}
