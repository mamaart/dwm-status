package netmodule

import (
	"errors"
	"log"
	"net"

	"github.com/vishvananda/netlink"
)

type NetModule struct {
	output <-chan []byte
}

func New() *NetModule {
	output := make(chan []byte)
	go func() {
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
					output <- []byte(IFace{
						Name: "unknown",
						Addr: x.LinkAddress.IP.String(),
					}.String())
				} else {
					output <- []byte(IFace{
						Name: iface.Name,
						Addr: x.LinkAddress.IP.String(),
					}.String())
				}
			}
		}
	}()
	return &NetModule{output: output}
}

func (n *NetModule) Reader() <-chan []byte {
	return n.output
}

func Get() (out IFace, err error) {
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
				return IFace{Name: e.Name, Addr: ip.String()}, nil
			}
		}

	}
	return out, errors.New("didn't find any global unicast")
}
