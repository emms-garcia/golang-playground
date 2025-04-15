package cmd

import (
	"fmt"
	"strconv"

	"github.com/emms-garcia/golang-playground/todo-cli/internal/db"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [task ID]",
	Short: "Delete a task",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("invalid ID")
		}

		db, err := db.GetDatabase()
		if err != nil {
			return fmt.Errorf("error connecting to database: %v", err)
		}

		_, err = db.Exec("DELETE todos WHERE id = ?", id)
		if err != nil {
			return fmt.Errorf("failed to delete task: %v", err)
		}
		fmt.Printf("Task %d was deleted\n", id)
		return nil
	},
}
