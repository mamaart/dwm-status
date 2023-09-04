package main

import (
	"strconv"

	"github.com/mamaart/statusbar/internal/models"
	"github.com/spf13/cobra"
)

func main() {
	Cmd().Execute()
}

func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	cmd.AddCommand(Add())
	cmd.AddCommand(Del())

	return cmd
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
			New().Delete(id)
		},
	}
}

func Add() *cobra.Command {
	return &cobra.Command{
		Use:     "add id description",
		Aliases: []string{"put", "ad", "addd"},
		Args:    cobra.MinimumNArgs(2),
		Example: `add 1 "go to work"`,
		Short:   "adds a task to your todolist",
		Run: func(cmd *cobra.Command, args []string) {
			id, err := strconv.Atoi(args[0])
			if err != nil {
				cmd.Help()
				return
			}
			New().Post(models.Task{
				Id:          id,
				Description: args[1],
			})
		},
	}
}
