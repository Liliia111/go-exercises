package cmd

import (
	"awesomeProject4/task/db"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "do tasks.",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Failed to parse argument: ", arg)
			} else {
				ids = append(ids, id)
			}
		}
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("Something went wrong", err)
			return
		}
		for _, id := range ids {
			if id <= 0 || id > len(tasks) {
				fmt.Println("invalid task number:", id)
				continue
			}
			task := tasks[id-1]
			err := db.DeleteTask(task.Key)
			if err != nil {
				fmt.Printf("Failed to mark \"%d\" as completed. Error: %s\n", id, err)
			} else {
				fmt.Printf("Marked as \"%d\" completed", id)
			}
		}
		fmt.Println(ids)
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
