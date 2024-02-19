package diskmodule

import (
	"fmt"
	"log"
	"syscall"
	"time"
)

type DiskModule struct {
	output <-chan []byte
}

func New() *DiskModule {
	output := make(chan []byte)
	go func() {
		for {
			d, err := Get()
			if err != nil {
				log.Println(err)
				time.Sleep(time.Second * 10)
			} else {
				output <- []byte(d.String())
				time.Sleep(time.Minute)
			}
		}
	}()
	return &DiskModule{output: output}
}

func (d *DiskModule) Reader() <-chan []byte {
	return d.output
}

func Get() (Disk, error) {
	var statfs syscall.Statfs_t

	err := syscall.Statfs("/", &statfs)
	if err != nil {
		return "", fmt.Errorf("error getting disk usage: %s", err)
	}

	var (
		total      = statfs.Blocks * uint64(statfs.Bsize)
		available  = statfs.Bavail * uint64(statfs.Bsize)
		used       = total - available
		percentage = float64(used) / float64(total) * 100.0
	)

	return Disk(fmt.Sprintf("%.1f%%", percentage)), nil
}
