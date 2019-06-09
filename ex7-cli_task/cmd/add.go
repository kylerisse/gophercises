package cmd

import (
	"strings"

	"../task"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add a task to the list",
	Run: func(cmd *cobra.Command, args []string) {
		tl := task.OpenTaskList()
		tl.Add(strings.Join(args, " "))
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
