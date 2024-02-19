package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/spf13/cobra"
)

func main() {
	cmd := Ask()
	cmd.CompletionOptions.HiddenDefaultCmd = true
	cmd.ExecuteC()
}

func Ask() *cobra.Command {
	return &cobra.Command{
		Use:     "ask [question]",
		Args:    cobra.MinimumNArgs(1),
		Example: `ask "what is 2 + 2?"`,
		Short:   "asks a question to the ai",
		Run: func(_ *cobra.Command, args []string) {
			v := url.Values{}
			v.Add("question", args[0])
			resp, err := http.PostForm("http://localhost:4343/ask", v)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(resp.Status)
			data, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(string(data))
		},
	}
}
