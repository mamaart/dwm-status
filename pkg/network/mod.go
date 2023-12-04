package network

import (
	"errors"
	"log"
	"net"

	"github.com/mamaart/statusbar/internal/models"
	"github.com/vishvananda/netlink"
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

func Stream(errch chan<- error) (<-chan models.IFace, error) {
	ch := make(chan models.IFace)
	go stream(ch, errch)
	return ch, nil
}

func stream(output chan<- models.IFace, err chan<- error) {
	ch := make(chan netlink.AddrUpdate)
	done := make(chan struct{})
	if err := netlink.AddrSubscribeWithOptions(ch, done, netlink.AddrSubscribeOptions{
		ListExisting: true,
	}); err != nil {
		log.Println(err)
		return
	}
	for x := range ch {
		if x.LinkAddress.IP.IsGlobalUnicast() {
			iface, err := net.InterfaceByIndex(x.LinkIndex)
			if err != nil {
				output <- models.IFace{
					Name: "unknown",
					Addr: x.LinkAddress.IP.String(),
				}
			} else {
				output <- models.IFace{
					Name: iface.Name,
					Addr: x.LinkAddress.IP.String(),
				}
			}
		}
	}
}
