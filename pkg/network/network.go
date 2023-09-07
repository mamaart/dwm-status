package network

import (
	"errors"
	"net"
	"time"

	"github.com/mamaart/statusbar/internal/models"
)

func Get() (out models.IFace, err error) {
	ifs, err := net.Interfaces()
	if err != nil {
		return out, err
	}

	for _, e := range ifs {
		addrs, err := e.Addrs()
		if err != nil {
			return out, err
		}

		for _, a := range addrs {
			ip, _, err := net.ParseCIDR(a.String())
			if err != nil {
				return out, err
			}
			if ip.IsGlobalUnicast() {
				return models.IFace{Name: e.Name, Addr: ip.String()}, nil
			}
		}

	}
	return out, errors.New("didn't find any global unicast")
}

func Stream(errch chan<- error) (chan models.IFace, error) {
	ch := make(chan models.IFace)
	go stream(ch, errch)
	return ch, nil
}

func stream(output chan<- models.IFace, err chan<- error) {
	for {
		w, e := Get()
		if e != nil {
			if err != nil {
				err <- e
			}
			time.Sleep(time.Minute)
		} else {
			output <- w
			time.Sleep(time.Minute * 10)
		}
	}
}
