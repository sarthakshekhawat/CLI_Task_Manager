package cmd

import (
	"fmt"
	"strconv"

	"github.com/sarthakshekhawat/CLI_Task_Manager/task/db"
	"github.com/spf13/cobra"
)

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Removes the task from the list",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				println("Can not parse the element with id = ", arg)
			} else {
				ids = append(ids, id)
			}
		}

		tasks, err := db.ListAllTasks()
		if err != nil {
			fmt.Println("Something went wrong. Error: ", err)
			return
		}
		for _, id := range ids {
			if id < 1 || id > len(tasks) {
				fmt.Printf("id=%d, is not valid\n", id)
			} else {
				err := db.DeleteTask(tasks[id-1].Key)
				if err != nil {
					fmt.Println("Something went wrong. Error: ", err)
				} else {
					fmt.Printf("Task \"%s\" has been deleted from the list\n", tasks[id-1].Value)
				}
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(rmCmd)
}
