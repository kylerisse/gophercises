package cmd

import (
	"fmt"

	"../task"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list list",
	Run: func(cmd *cobra.Command, args []string) {
		tl := task.OpenTaskList()
		fmt.Print(tl.List())
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
