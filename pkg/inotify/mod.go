package inotify

import (
	"fmt"
	"path/filepath"
	"strings"
	"syscall"
	"time"
	"unsafe"
)

func Listen(path string) chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)

		dir := filepath.Dir(path)
		fd, err := syscall.InotifyInit1(syscall.IN_CLOEXEC | syscall.IN_NONBLOCK)
		if err != nil {
			ch <- err.Error()
			return
		}
		defer syscall.Close(fd)

		w, err := syscall.InotifyAddWatch(fd, dir, syscall.IN_MODIFY)
		if err != nil {
			ch <- err.Error()
			return
		}
		defer syscall.InotifyRmWatch(fd, uint32(w))

		fdset := &syscall.FdSet{}
		fdset.Bits[0] = 1 << uint(fd)

		buf := make([]byte, 1024*(syscall.SizeofInotifyEvent+16))
		for {
			var n int
			for {
				timeout := syscall.NsecToTimeval(time.Second.Nanoseconds())

				_, err = syscall.Select(fd+1, fdset, nil, nil, &timeout)

				if err != nil {
					if err == syscall.EINTR {
						continue
					}
					ch <- err.Error()
					return
				}
				n, err = syscall.Read(fd, buf)
				if err != nil {
					if err == syscall.EAGAIN {
						continue
					}
					ch <- err.Error()
					return
				}
				if n > 0 {
					break
				}
			}

			if n < syscall.SizeofInotifyEvent {
				ch <- "Short inotify read"
				return
			}

			var offset int

			for offset+syscall.SizeofInotifyEvent <= n {

				event := (*syscall.InotifyEvent)(unsafe.Pointer(&buf[offset]))
				namebuf := buf[offset+syscall.SizeofInotifyEvent : offset+syscall.SizeofInotifyEvent+int(event.Len)]

				offset += syscall.SizeofInotifyEvent + int(event.Len)

				name := strings.TrimRight(string(namebuf), "\x00")
				name = filepath.Join(dir, name)
				if name == path {
					// TODO loop forever and emit to channel
					ch <- fmt.Sprintln(
						"Wd:", uint32(event.Wd),
						", Name:", name,
						", Mask:", event.Mask,
						", Cookie:", event.Cookie,
					)
				}
			}
		}
	}()
	return ch
}
