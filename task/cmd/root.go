package cmd

import "github.com/spf13/cobra"

//RootCmd represents the root command
var RootCmd = &cobra.Command{
	Use:   "task",
	Short: "Task is a CLI Task Manager",
}
