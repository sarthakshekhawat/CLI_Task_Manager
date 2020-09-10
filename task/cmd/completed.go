package cmd

import (
	"fmt"

	"github.com/sarthakshekhawat/CLI_Task_Manager/task/db"
	"github.com/spf13/cobra"
)

// completedCmd represents the completed command
var completedCmd = &cobra.Command{
	Use:   "completed",
	Short: "Shows all the completed tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.ListAllCompletedTasks()
		if err != nil {
			fmt.Println("Something went wrong. Error: ", err)
			return
		}
		if len(tasks) == 0 {
			fmt.Println("You have not completed any task in the past 24hrs!")
		} else {
			fmt.Println("You have completed following tasks in the past 24hrs:")
			for i, t := range tasks {
				fmt.Printf("%d. %s\n", i+1, t.Value)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(completedCmd)
}
