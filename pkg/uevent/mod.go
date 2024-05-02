package uevent

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
	"syscall"
)

func Listen() chan bool {
	ch := make(chan bool)
	go func() {
		fd, err := syscall.Socket(
			syscall.AF_NETLINK,
			syscall.SOCK_RAW,
			syscall.NETLINK_KOBJECT_UEVENT,
		)
		if err != nil {
			log.Fatal(err)
		}
		defer syscall.Close(fd)

		nl := syscall.SockaddrNetlink{
			Family: syscall.AF_NETLINK,
			Pid:    uint32(os.Getpid()),
			Groups: 1,
		}

		err = syscall.Bind(fd, &nl)
		if err != nil {
			log.Fatal(err)
		}

		rd := bufio.NewReader(os.NewFile(uintptr(fd), "fd3"))
		for {
			s, err := rd.ReadBytes(0x00)
			if err != nil {
				log.Fatal(err)
			}

			xs := strings.Split(string(bytes.TrimSuffix(s, []byte{0x00})), "=")
			if len(xs) == 2 {
				fmt.Println("---", xs[0], "=", xs[1])
				if xs[0] == "POWER_SUPPLY_ONLINE" {
					ch <- xs[1] == "1"
				}
			}
		}
	}()
	return ch
}
