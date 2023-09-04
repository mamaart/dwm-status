package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/mamaart/statusbar/internal/models"
)

type Client struct {
	httpClient *http.Client
}

func New() *Client {
	return &Client{
		httpClient: unixClient("/tmp/statusbar.sock"),
	}
}

func (c *Client) Post(task models.Task) {
	data := must(json.Marshal(task))
	r := must(http.NewRequest(http.MethodPost, "http://unix/", bytes.NewReader(data)))
	resp, err := c.httpClient.Do(r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.Status)
}

func (c *Client) Delete(id int) {
	r := must(http.NewRequest(http.MethodDelete, fmt.Sprintf("http://unix/%d", id), nil))
	resp, err := c.httpClient.Do(r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.Status)
}

func unixClient(socketPath string) *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			Dial: func(network, addr string) (net.Conn, error) {
				return net.Dial("unix", socketPath)
			},
		},
	}
}

func must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}
