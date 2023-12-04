package disk

import (
	"fmt"
	"syscall"
	"time"

	"github.com/mamaart/statusbar/internal/models"
)

func Get() (models.Disk, error) {
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

	return models.Disk(fmt.Sprintf("%.1f%%", percentage)), nil
}

func Stream(errch chan<- error) (<-chan models.Disk, error) {
	ch := make(chan models.Disk)
	go stream(ch, errch)
	return ch, nil
}

func stream(output chan<- models.Disk, err chan<- error) {
	for {
		d, e := Get()
		if e != nil {
			if err != nil {
				err <- e
			}
			time.Sleep(time.Second * 10)
		} else {
			output <- d
			time.Sleep(time.Minute)
		}
	}
}
