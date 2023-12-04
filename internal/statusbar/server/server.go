package server

import (
	"context"
	"crypto/tls"
	"log"
	"time"

	"github.com/mamaart/statusbar/pkg/p256"
	"github.com/quic-go/quic-go"
)

func Run(ch chan<- byte) {
	// p256.Generate("olla")

	listener, err := quic.ListenAddr(":4343", &tls.Config{
		Certificates: []tls.Certificate{p256.Get("/etc/statusbar/olla")},
		NextProtos:   []string{"dwm-status"},
	}, nil)
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept(context.Background())
		if err != nil {
			log.Println(err)
			time.Sleep(time.Second * 5)
			continue
		}
		stream, err := conn.AcceptStream(context.Background())
		if err != nil {
			log.Println(err)
			time.Sleep(time.Second * 5)
			continue
		}
		b := make([]byte, 256)
		for {
			n, err := stream.Read(b)
			if err != nil {
				log.Println(err)
				break
			}
			for _, b := range b[:n] {
				ch <- b
			}
		}
	}
}

func must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}
