package main

import (
	"strconv"

	"github.com/mamaart/statusbar/internal/models"
	"github.com/mamaart/statusbar/internal/statusbarctl/client"
	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{
		Use: "statusbarctl",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	cmd.CompletionOptions.HiddenDefaultCmd = true

	cmd.AddCommand(Add())
	cmd.AddCommand(Del())
	cmd.ExecuteC()
}

func Del() *cobra.Command {
	return &cobra.Command{
		Use:     "del id",
		Aliases: []string{"delete", "d", "de"},
		Args:    cobra.MinimumNArgs(1),
		Example: "del 1",
		Short:   "deletes a task from the todo list",
		Run: func(cmd *cobra.Command, args []string) {
			id, err := strconv.Atoi(args[0])
			if err != nil {
				cmd.Help()
				return
			}
			client.New().Delete(id)
		},
	}
}

func Add() *cobra.Command {
	return &cobra.Command{
		Use:     "add description",
		Aliases: []string{"put"},
		Args:    cobra.MinimumNArgs(1),
		Example: `add "go to work"`,
		Short:   "adds a task to your todolist",
		Run: func(cmd *cobra.Command, args []string) {
			client.New().Post(models.Task{
				Description: args[0],
			})
		},
	}
}
