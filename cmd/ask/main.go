package main

import (
	"log"

	"github.com/mamaart/statusbar/pkg/ai"
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
			if err := ai.Ask(args[0]); err != nil {
				log.Fatal(err)
			}
		},
	}
}
