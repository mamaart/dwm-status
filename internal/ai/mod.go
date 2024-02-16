package ai

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

func Ask(q string) error {
	v := url.Values{}
	v.Add("question", q)
	resp, err := http.PostForm("http://localhost:4343/ask", v)
	if err != nil {
		return err
	}
	fmt.Println(resp.Status)
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
	return nil
}
