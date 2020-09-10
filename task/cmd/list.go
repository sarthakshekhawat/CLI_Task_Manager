package cmd

import (
	"fmt"

	"github.com/sarthakshekhawat/CLI_Task_Manager/task/db"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Shows all the tasks in your list",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.ListAllTasks()
		if err != nil {
			fmt.Println("Something went wrong. Error: ", err)
			return
		}
		if len(tasks) == 0 {
			fmt.Println("You do not have any task in the list!")
		} else {
			fmt.Println("You have following tasks in the list:")
			for i, t := range tasks {
				fmt.Printf("%d. %s\n", i+1, t.Value)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
