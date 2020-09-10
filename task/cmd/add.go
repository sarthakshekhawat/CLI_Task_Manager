package cmd

import (
	"fmt"
	"strings"

	"github.com/sarthakshekhawat/CLI_Task_Manager/task/db"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task in your list",
	Run: func(cmd *cobra.Command, args []string) {
		str := strings.Join(args, " ")
		err := db.AddTask(str)
		if err != nil {
			fmt.Println("Somthing went wrong while adding the task. Error:", err)
		} else {
			fmt.Printf("Task \"%s\" has been added in your list\n", str)
		}
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
