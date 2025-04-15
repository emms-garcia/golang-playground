package cmd

import (
	"fmt"
	"strconv"

	"github.com/emms-garcia/golang-playground/todo-cli/internal/db"
	"github.com/spf13/cobra"
)

var doneCmd = &cobra.Command{
	Use:   "done [task ID]",
	Short: "Mark a task as done",
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

		_, err = db.Exec("UPDATE todos SET done = 1 WHERE id = ?", id)
		if err != nil {
			return fmt.Errorf("failed to mark task as done: %v", err)
		}
		fmt.Printf("Task %d marked as done\n", id)
		return nil
	},
}
