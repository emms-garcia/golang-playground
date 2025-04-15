package cmd

import (
	"fmt"

	"github.com/emms-garcia/golang-playground/todo-cli/internal/db"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [task]",
	Short: "Add a new task to the to-do list",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		db, err := db.GetDatabase()
		if err != nil {
			return err
		}
		defer db.Close()

		task := args[0]
		if _, err = db.Exec("INSERT INTO todos (title, done) VALUES (?, ?)", task, false); err != nil {
			return fmt.Errorf("failed to add task: %v", err)
		}
		fmt.Println("Task added:", task)
		return nil
	},
}

func init() {
	addCmd.Flags().BoolP("completed", "c", false, "Whether the task is completed")
}
