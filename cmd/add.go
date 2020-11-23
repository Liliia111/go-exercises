package cmd

import (
	"awesomeProject4/task/db"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds tasks to tasks list.",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		_, err := db.CreateTask(task)
		if err != nil {
			fmt.Println("Something went wrong", err.Error())
			return
		}
		fmt.Printf("Added \"%s\" to your task list.\n", task)

	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
