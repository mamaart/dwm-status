package application

import (
	"errors"
	"fmt"
	"net"
	"time"

	"github.com/mamaart/statusbar/internal/models"
)

func (a *App) ifLoop() error {
	for {
		iface, err := getIface()
		if err != nil {
			fmt.Println(err)
			time.Sleep(time.Second * 10)
			continue
		}
		a.iface <- iface
		time.Sleep(time.Second * 2)
	}
}

func getIface() (out models.IFace, err error) {
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
