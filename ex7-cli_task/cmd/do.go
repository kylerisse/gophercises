package cmd

import (
	"../task"
	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "mark item as complete",
	Run: func(cmd *cobra.Command, args []string) {
		tl := task.OpenTaskList()
		for _, i := range args {
			tl.Done(i)
		}
	},
}

func init() {
	rootCmd.AddCommand(doCmd)
}
